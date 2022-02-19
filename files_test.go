package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

const EXAMPLE_FILE = "resources/example.txt"

var homeDirectory string

func init() {
	if homeDirectory = os.Getenv("HOME"); len(homeDirectory) == 0 {
		log.Fatal("set HOME environment variable")
	}

	if err := os.Mkdir("build", 0755); err != nil {
		log.Fatal(fmt.Sprintf("failed to create build directory: %s", err))
	}
}

func tearDown() {
	files := []string{"build/example.txt", "build"}
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			log.Fatal(fmt.Sprintf("failed to remove %s directory", file))
		}
	}
}

func TestFileCopy(t *testing.T) {
	defer tearDown()

	srcFile := fmt.Sprintf("%s/../../etc/passwd", homeDirectory)
	destFile := fmt.Sprintf("%s/passwd", homeDirectory)
	if err := copyFile(srcFile, destFile); err != nil {
		t.Errorf("expected no error from bogus command, but got %v", err)
	}

	if err := copyFileSafe(srcFile, destFile); err == nil {
		t.Errorf("expected error but got none")
	}

	if err := copyFileSafe(EXAMPLE_FILE, "build/example.txt"); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}
