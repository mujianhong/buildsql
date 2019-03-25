package buildsql

import (
	"fmt"
	"strings"
)

// UpdateStruct xx
type UpdateStruct struct {
	publicAttribute
	set []string
}

// Table xx
func (u *UpdateStruct) Table(table string) {
	u.tableName = table
}

// Where XX
func (u *UpdateStruct) Where(where string) {
	u.where.setWhere(where)
}

// Set xx
func (u *UpdateStruct) Set(column string, value interface{}) {

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
		return
	}

	u.set = append(u.set, fmt.Sprintf("%s = %s", column, setVal))
}

// ToString xx
func (u UpdateStruct) ToString() string {

	if u.tableName == "" {
		return ""
	}

	sql := fmt.Sprintf("UPDATE %s", u.tableName)

	if len(u.set) == 0 {
		return ""
	}

	sql = fmt.Sprintf("%s SET %s", sql, strings.Join(u.set, ", "))

	return fmt.Sprintf("%s %s", sql, u.where.toString())
}
