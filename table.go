package commabase

import "fmt"

// Table holds information about a table.
// A table is a csv file inside a database.
type Table struct {
	ShortName    string
	LongName     string
	Path         string
	DatabaseName string
}

func NewTable(tableName, databaseName string) *Table {
	longName := fmt.Sprintf("%s.%s", databaseName, tableName)
	path := fmt.Sprintf("%s/%s.csv", databaseName, tableName)
	return &Table{tableName, longName, path, databaseName}
}

func (t *Table) String() string {
	return t.LongName
}

// SELECT
//func (t *Table) Select()

// INSERT INTO

// DELETE

// UPDATE
