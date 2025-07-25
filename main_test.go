package main

import (
	"html/template"
	"os"
	"path/filepath"
	"testing"
)

func TestTemplateRendering(t *testing.T) {
	err := os.MkdirAll("dist", os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create dist directory: %v", err)
	}

	os.MkdirAll("templates", os.ModePerm)
	err = os.WriteFile("templates/index.html", []byte("{{.}}"), 0644)
	if err != nil {
		t.Fatalf("Failed to write test template: %v", err)
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	outFile, err := os.Create("dist/index.html")
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	defer outFile.Close()

	err = tmpl.Execute(outFile, "Hello Test")
	if err != nil {
		t.Fatalf("Template execution failed: %v", err)
	}
}

func TestCopyStaticFiles(t *testing.T) {
	os.MkdirAll("static", os.ModePerm)
	os.MkdirAll("dist", os.ModePerm)

	testFile := filepath.Join("static", "test.css")
	err := os.WriteFile(testFile, []byte("body {color: pink;}"), 0644)
	if err != nil {
		t.Fatalf("Failed to create static file: %v", err)
	}

	err = copyStaticFiles("static", "dist")
	if err != nil {
		t.Fatalf("Failed to copy static files: %v", err)
	}

	// Check if file exists in dist
	if _, err := os.Stat("dist/test.css"); os.IsNotExist(err) {
		t.Fatalf("Static file was not copied")
	}
}
