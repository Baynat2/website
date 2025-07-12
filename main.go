package main

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Create dist folder if not exists
	err := os.MkdirAll("dist", os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create dist directory: %v", err)
	}

	// Parse HTML template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	// Create output HTML file
	outFile, err := os.Create("dist/index.html")
	if err != nil {
		log.Fatalf("Failed to create index.html: %v", err)
	}
	defer outFile.Close()

	// Inject message into template
	err = tmpl.Execute(outFile, "Hello World :)")
	if err != nil {
		log.Fatalf("Failed to render template: %v", err)
	}

	// Copy static assets (CSS etc.)
	err = copyStaticFiles("static", "dist")
	if err != nil {
		log.Fatalf("Failed to copy static files: %v", err)
	}
}

func copyStaticFiles(srcDir, destDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			destPath := filepath.Join(destDir, info.Name())

			srcFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			destFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, srcFile)
			return err
		}
		return nil
	})
}
