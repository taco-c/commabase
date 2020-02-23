package commabase

import (
	"errors"
	"fmt"
	"os"
)

// Database holds information concerning a database.
// A database is just a folder in which the tables (csv files) live.
type Database struct {
	Name     string
	RootPath string
	Table    map[string]Table
}

// NewDatabase creates a new database on the filesystem,
// and returns a Database object.
func NewDatabase(name, root string) (*Database, error) {
	path := fmt.Sprintf("%s/%s", root, name)

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
	return &Database{name, root, make(map[string]Table)}, nil
}

// OpenDatabase opens an existing database on the filesystem,
// and returns a Database object.
func OpenDatabase(name, root string) (*Database, error) {
	path := fmt.Sprintf("%s/%s", root, name)

	exists, err := dirExists(path)
	if err != nil {
		return &Database{}, err
	}

	if !exists {
		e := fmt.Sprintf("Could not open database: %s: directory does not exist.", path)
		return &Database{}, errors.New(e)
	}

	// Loop through csv-files and add them as tables.
	return &Database{name, root, make(map[string]Table)}, nil
}

func (db *Database) AddTable(name string) {
	db.Table[name] = *NewTable(name, db.Name)
}

func (db *Database) String() string {
	return db.Name
}
