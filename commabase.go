package commabase

import (
	"os"
)

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	// Schodinger's file
	return false, err
}

func dirExists(path string) (bool, error) {
	f, err := os.Stat(path)
	if err == nil && f.IsDir() {
		return true, nil
	} else if os.IsNotExist(err) || !f.IsDir() {
		return false, nil
	}
	// Schodinger's dir
	return false, err
}
