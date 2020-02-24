package commabase

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Table holds information about a table.
// A table is a csv file inside a database.
type Table struct {
	Name     string
	Path     string
	Database *Database
	Columns  []string
	Rows     []Row
}

func (t *Table) String() string {
	return t.Path
}

// NewTable creates a new *Table object.
func NewTable(name, path string, database *Database) *Table {
	// Open and read a csv file.
	columns, err := readColumns(path)
	if err != nil {
		return &Table{}
	}
	return &Table{name, path, database, columns, make([]Row, 0)}
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

// Select selects the rows with the fields.
func (t *Table) Select(fields ...string) *Table {
	if len(fields) == 0 {
		return t
	}
	filtered := make([]Row, 0)
	for _, row := range t.Rows {
		filteredRow := make(Row, 0)
		for i := 0; i < len(fields); i++ {
			filteredRow[fields[i]] = row[fields[i]]
		}
		filtered = append(filtered, filteredRow)
	}
	t.Rows = filtered
	return t
}

// Where filters the rows of the database.
func (t *Table) Where(clause func(Row) bool) *Table {
	filtered := make([]Row, 0)
	for _, row := range t.Rows {
		if clause(row) {
			filtered = append(filtered, row)
		}
	}
	t.Rows = filtered
	return t
}

// INSERT INTO
// Adds a row to a Table.

// DELETE
// Removes a row from a table.

// UPDATE
