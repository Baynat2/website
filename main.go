package main

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

func main() {

	os.MkdirAll("dist", os.ModePerm)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create("dist/index.html")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = tmpl.Execute(outFile, "Hello World:)")
	if err != nil {
		panic(err)
	}

	copyStaticFiles("static", "dist")
}

func copyStaticFiles(srcDir, destDir string) {
	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
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
