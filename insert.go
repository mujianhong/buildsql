package buildsql

import (
	"fmt"
	"strings"
)

// InsertStruct xx
type InsertStruct struct {
	tableName string
	columns   []string
	values    []string
}

// Table xx
func (i *InsertStruct) Table(table string) {
	i.tableName = table
}

// Set xx
func (i *InsertStruct) Set(columns map[string]interface{}) {

	for column, value := range columns {

		i.columns = append(i.columns, column)

		setVal := ""
		switch value.(type) {
		case bool:
			setVal = fmt.Sprintf("%t", value)
		case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64:
			setVal = fmt.Sprintf("%d", value)
		case float32, float64:
			setVal = fmt.Sprintf("'%f'", value)
		case string:
			setVal = fmt.Sprintf("'%s'", value)
		default:
			continue
		}
		i.values = append(i.values, setVal)
	}

}

// ToString xx
func (i InsertStruct) ToString() string {

	if i.tableName == "" {
		return ""
	}

	if len(i.columns) == 0 || (len(i.columns) != len(i.values)) {
		return ""
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", i.tableName, strings.Join(i.columns, ", "), strings.Join(i.values, ", "))

	return sql
}
