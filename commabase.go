package commabase

import (
	"bufio"
	"log"
	"os"
	"strings"
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

func readColumns(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s, err := bufio.NewReader(f).ReadString('\n')
	if err != nil {
		return make([]string, 0), err
	}
	return strings.Split(s, ","), nil
}
