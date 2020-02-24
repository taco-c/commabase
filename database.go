package commabase

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Database holds information concerning a database.
// A database is just a folder in which the tables (csv files) live.
type Database struct {
	Path  string
	Table map[string]*Table
}

// Create creates a new database on the filesystem,
// and returns a Database object.
func Create(path string) (*Database, error) {
	exists, err := fileExists(path)
	if err != nil {
		return &Database{}, err
	}

	if exists {
		e := fmt.Sprintf("Could not create new database: %s: file already exists.", path)
		return &Database{}, errors.New(e)
	}

	fmt.Println("Creates dir")
	os.MkdirAll(path, 0666)
	return &Database{path, make(map[string]*Table)}, nil
}

// Open opens an existing database on the filesystem,
// and returns a Database object.
func Open(path string) (*Database, error) {
	exists, err := dirExists(path)
	if err != nil {
		return &Database{}, err
	}

	if !exists {
		e := fmt.Sprintf("Could not open database: %s: directory does not exist.", path)
		return &Database{}, errors.New(e)
	}

	// Loop through csv-files and add them as tables.
	// tableFiles := make([]os.FileInfo, 0)

	db := &Database{}
	tables := make(map[string]*Table)
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		tables[info.Name()] = NewTable(info.Name(), path, db)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return &Database{path, tables}, nil
}

// CreateTable creates a csv file in the database directory.
func (db *Database) CreateTable(name string) {
	db.Table[name] = &Table{}
}

func (db *Database) String() string {
	return db.Path
}
