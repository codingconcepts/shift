package repo

import (
	"database/sql"
	"fmt"
	"shift/internal/pkg/model"
)

// scan a row collection for a given table into a multi-dimensional array.
func scan(list *sql.Rows, t model.Table) ([][]any, error) {
	fields, err := list.Columns()
	if err != nil {
		return nil, fmt.Errorf("listing columns: %w", err)
	}

	var rows []map[string]any
	for list.Next() {
		scans := make([]any, len(fields))
		row := make(map[string]any)

		for i := range scans {
			scans[i] = &scans[i]
		}

		if err = list.Scan(scans...); err != nil {
			return nil, fmt.Errorf("scaning values: %w", err)
		}

		for i, v := range scans {
			if v != nil {
				row[fields[i]] = v
			}
		}
		rows = append(rows, row)
	}

	return mapToNArray(rows, t), nil
}

func mapToNArray(m []map[string]any, t model.Table) [][]any {
	array := [][]any{}

	for _, row := range m {
		columns := make([]any, len(t.Columns))
		for i, col := range t.Columns {
			columns[i] = row[col.Name]
		}
		array = append(array, columns)
	}

	return array
}
