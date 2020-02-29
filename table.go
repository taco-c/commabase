package commabase

// Table holds information about a table.
// A table is a csv file inside a database.
type Table struct {
	Name     string
	Path     string
	Database *Database
	Columns  []string
	Rows     Rows
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

func (t *Table) String() string {
	return t.Path
}
