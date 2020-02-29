package commabase

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	os.MkdirAll(path, 0755)
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

// From gets all the rows from the file.
func (db *Database) From(table string) *Rows {
	path := fmt.Sprintf("%s/%s.csv", db.Path, table)
	f, _ := os.Open(path)
	defer f.Close()

	sc := bufio.NewReader(f)
	var columns []string
	var row Row
	allRows := make(Rows, 0)

	// rows := make(Rows, 0)
	i := 0
	for line, err := sc.ReadString('\n'); err == nil; line, err = sc.ReadString('\n') {
		line = strings.TrimSuffix(line, "\n")
		if i == 0 {
			columns = strings.Split(line, ",")
			i++
			continue
		}
		items := strings.Split(line, ",")
		row = make(Row, 0)
		for i, column := range columns {
			row[column] = items[i]
		}
		allRows = append(allRows, row)
	}

	return &allRows
}

func (db *Database) String() string {
	return db.Path
}
