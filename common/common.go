package common

import (
	"fmt"
	"os"
	"path/filepath"
)

// dataPathRoot is the local definition for where data is stored.
const dataPathRoot = "./data"

var (
	// UnsafeChars are the chars that are not safe within a url.
	UnsafeChars = []string{
		"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",",
		"/", ":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^",
		"`", "{", "|", "}",
	}

	// SafeChars are the safe chars, respectively.
	SafeChars = []string{
		"%21", "%22", "%23", "%24", "%25", "%26", "%27", "%28", "%29",
		"%2A", "%2B", "%2C", "%2F", "%3A", "%3B", "%3C", "%3D", "%3E",
		"%3F", "%40", "%5B", "%5C", "%5D", "%5E", "%5F", "%60",
	}

	// DataPath is the global definition for where data is stored.
	DataPath = GetDataPath()

	// DBPath is the global definition for where database data is stored.
	DBPath = filepath.FromSlash(fmt.Sprintf("%s/db", DataPath))

	// LogsDir is the global logs directory definition.
	LogsDir = filepath.FromSlash(fmt.Sprintf("%s/logs", DataPath))
)

// GetDataPath gets the working data path.
func GetDataPath() string {
	path, _ := filepath.Abs(filepath.FromSlash(fmt.Sprintf(dataPathRoot))) // Get data path

	return path // Return path
}

// CreateDirIfDoesNotExist creates a directory if it does not already exist.
func CreateDirIfDoesNotExist(dir string) error {
	dir = filepath.FromSlash(dir) // Replace dir separators depending on OS

	if _, err := os.Stat(dir); os.IsNotExist(err) { // Check dir does not exist
		err = os.MkdirAll(dir, 0755) // Recursively make directory
		if err != nil {              // Check for errors
			return err // Return found error
		}
	}

	return nil // No error occurred, return nil
}
