package commabase

import "fmt"

// Table holds information about a table.
// A table is a csv file inside a database.
type Table struct {
	ShortName    string
	LongName     string
	Path         string
	DatabaseName string
	Rows         []Row
}

func (t *Table) String() string {
	return t.LongName
}

// NewTable create a new blank Table.
func NewTable(tableName, databaseName string) *Table {
	longName := fmt.Sprintf("%s.%s", databaseName, tableName)
	path := fmt.Sprintf("%s/%s.csv", databaseName, tableName)
	return &Table{tableName, longName, path, databaseName, make([]Row, 0)}
}

// OpenTable will open an existing csv file and return a *Table.
func OpenTable(tableName, databaseName string) *Table {
	longName := fmt.Sprintf("%s.%s", databaseName, tableName)
	path := fmt.Sprintf("%s/%s.csv", databaseName, tableName)
	// Open and read a csv file.
	return &Table{tableName, longName, path, databaseName, make([]Row, 0)}
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

// DELETE

// UPDATE
