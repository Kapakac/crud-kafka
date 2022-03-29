package helpers

import (
	"database/sql"
)

func MakeArray(rows *sql.Rows) (allMaps []map[string]interface{}) {
	cols, _ := rows.Columns()

	for rows.Next() {
		values := make([]interface{}, len(cols))
		pointers := make([]interface{}, len(cols))
		for i, _ := range values {
			pointers[i] = &values[i]
		}
		rows.Scan(pointers...)
		resultMap := make(map[string]interface{})
		for i, val := range values {
			resultMap[cols[i]] = val
		}
		allMaps = append(allMaps, resultMap)
	}

	return allMaps
}
