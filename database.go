package commabase

import (
	"errors"
	"fmt"
	"os"
)

type Database struct {
	Name string
	RootPath string
	Table map[string]Table
}

func NewDatabase(name, root string) (*Database, error) {
	path := fmt.Sprintf("%s/%s", root, name)
	
	_, err := os.Stat(path)
	if err == nil {
		// File/dir exists
		e := fmt.Sprintf("Could not create new database: %s: file already exists", path)
		return &Database{}, errors.New(e)
	} else if os.IsNotExist(err) {
		// File/dir does not exist
		// create
		fmt.Println("Creates dir")
		os.MkdirAll(path, 0666)
		return &Database{name, root, make(map[string]Table)}, nil
	}
	// Schodinger's file/dir
	return &Database{}, err
}

func (db *Database) AddTable(name string) {
	db.Table[name] = *NewTable(name, db.Name)
}

func (db *Database) String() string {
	return db.Name
}

