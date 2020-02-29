package commabase

// A Row rows.
type Row map[string]string

// Rows row.
type Rows []Row

// func (r *Rows) Iter() <-chan Row {
// 	ch := make(chan Row)
// 	for i := 0; i < len(*r); i++ {
// 		ch <- *r.Get(i)
// 	}
// 	return ch
// }

// Select filters on columns.
func (r *Rows) Select(columns ...string) *Rows {
	if len(columns) == 0 {
		return r
	}
	filtered := make(Rows, 0)
	for _, row := range *r {
		filteredRow := make(Row, 0)
		for i := 0; i < len(columns); i++ {
			filteredRow[columns[i]] = row[columns[i]]
		}
		filtered = append(filtered, filteredRow)
	}
	return &filtered
}

// Where filters the rows of the database.
func (r *Rows) Where(clause func(Row) bool) *Rows {
	filtered := make(Rows, 0)
	for _, row := range *r {
		if clause(row) {
			filtered = append(filtered, row)
		}
	}
	return &filtered
}

// Limit limits the rows to the desired lenght.
func (r *Rows) Limit(count int) *Rows {
	filtered := make(Rows, 0)
	for i, row := range *r {
		if i >= count {
			break
		}
		filtered = append(filtered, row)
	}

	return &filtered
}
