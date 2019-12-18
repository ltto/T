package sqlTpl

import (
	"database/sql"

	"github.com/ltto/T/Tsql"
	"github.com/ltto/T/gobox/ref"
)

func rows2maps(rows *sql.Rows) (resultsSlice Tsql.QueryResult, err error) {
	for i := 0; rows.Next(); i++ {
		result, err := row2map(rows)
		if err != nil {
			return Tsql.QueryResult{}, err
		}
		resultsSlice.Append(result)
	}
	return resultsSlice, nil
}

func row2map(rows *sql.Rows) (resultsMap map[string][]ref.Val, err error) {
	var fields []*sql.ColumnType
	if fields, err = rows.ColumnTypes(); err != nil {
		return nil, err
	}
	var scanResultContainers = make([]interface{}, len(fields))
	for i := 0; i < len(fields); i++ {
		var scanResultContainer interface{}
		scanResultContainers[i] = &scanResultContainer
	}
	if err := rows.Scan(scanResultContainers...); err != nil {
		return nil, err
	}
	var result = make(map[string][]ref.Val)
	for ii, key := range fields {
		result[key.Name()] = []ref.Val{ref.NewVal(scanResultContainers[ii])}
	}
	return result, nil
}
