package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const COPY_SIZE = 1024

// Copies a file provided in the source parameter to the target.
// Leverages the default io.Copy function that can be exploited as
// the copy terminates when either an EOF is seen or an error
// occurs reading the file.
// Also this method does not clean the fileName passed into the
// method. This can be exploited to read arbitrary files from the
// file system.
// See copyFileSafe for an alternative.
func copyFile(source string, target string) error {
	// nolint:gosec // provided for demonstration purposes
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(srcFile)

	// nolint:gosec // provided for demonstration purposes
	targetFile, err := os.Create(target)
	if err != nil {
		return err
	}

	_, err = io.Copy(targetFile, reader)
	return err
}

// Copies a file safely.
// Using filepath.Clean the filename is parsed to its actual value,
// and the code checks that is the require prefix in the path.
// Otherwise an error is returned.
// The code is also using io.CopyN() to copy the file in COPY_SIZE
// chunks and verifies the return value accordingly.
// Also note that errors should be compared with the errors.Is()
// function call and not using regular operators such as ==.
func copyFileSafe(source string, target string) error {
	cleanSource := filepath.Clean(source)
	if !strings.HasPrefix(cleanSource, "resources") {
		return fmt.Errorf("invalid source path")
	}

	cleanTarget := filepath.Clean(target)
	if !strings.HasPrefix(cleanTarget, "build") {
		return fmt.Errorf("invalid target path")
	}

	srcFile, err := os.Open(cleanSource)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(srcFile)

	targetFile, err := os.Create(cleanTarget)
	if err != nil {
		return err
	}

	for {
		_, err := io.CopyN(targetFile, reader, COPY_SIZE)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
	}
	return nil
}
