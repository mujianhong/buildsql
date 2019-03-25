package buildsql

import (
	"fmt"
)

// DeleteStruct xx
type DeleteStruct struct {
	publicAttribute
}

// Table xx
func (u *DeleteStruct) Table(table string) {
	u.tableName = table
}

// Where XX
func (u *DeleteStruct) Where(where string) {
	u.where.setWhere(where)
}

// ToString xx
func (u DeleteStruct) ToString() string {

	if u.tableName == "" {
		return ""
	}

	sql := fmt.Sprintf("DELETE FROM %s", u.tableName)

	return fmt.Sprintf("%s %s", sql, u.where.toString())
}
