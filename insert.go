package buildsql

import (
	"fmt"
	"strings"
)

type insertBuilder struct {
	//
	SQLBuilder
	//
	sqlStruct
	//
	se selectBuilder

	isSetSelect bool

	vaules []string
}

// 获取sql
func (i insertBuilder) ToString() string {

	sql := "INSERT INTO %s %s"

	if len(i.columns) > 0 {
		sql = strings.TrimSpace(fmt.Sprintf(sql, i.table, fmt.Sprintf("(%s)", strings.Join(i.columns, ", "))))
	} else {
		sql = strings.TrimSpace(fmt.Sprintf(sql, i.table, ""))
	}

	if i.isSetSelect {
		return fmt.Sprintf("%s (%s)", sql, i.se.ToString())
	}

	return ""
}

// SetSelectBuilder 设置 select语句
func (i *insertBuilder) SetSelectBuilder(se selectBuilder) *insertBuilder {

	i.se = se
	i.isSetSelect = true
	return i
}

// SetTable 设置表名
func (i *insertBuilder) SetTable(table string) *insertBuilder {

	i.table = table
	return i
}

// SetColumns 设置字段
func (i *insertBuilder) SetColumns(column ...string) *insertBuilder {

	if len(column) > 0 {
		for _, v := range column {
			i.columns = append(i.columns, v)
		}
	}
	return i
}
