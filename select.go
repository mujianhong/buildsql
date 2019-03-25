package buildsql

import (
	"fmt"
	"strings"
)

// SelectStruct xx
type SelectStruct struct {
	publicAttribute
	columns []string
	groupBy []string
	having  []string
	orderBy []string
	limit   int
	offSet  int
}

// Where XX
func (s *SelectStruct) Where(where string) {
	s.where.setWhere(where)
}

// SetColumn xx
func (s *SelectStruct) SetColumn(column ...string) {
	for _, v := range column {
		s.columns = append(s.columns, fmt.Sprintf("%s", strings.TrimSpace(v)))
	}

}

// GroupBy xx
func (s *SelectStruct) GroupBy(column ...string) {
	for _, v := range column {
		s.groupBy = append(s.groupBy, v)
	}
}

// Having xxx
func (s *SelectStruct) Having(having ...string) {

	for _, v := range having {
		s.having = append(s.having, fmt.Sprintf("(%s)", v))
	}

}

// OrderBy xx
func (s *SelectStruct) OrderBy(column ...string) {

	for _, v := range column {
		s.orderBy = append(s.orderBy, v)
	}
}

// Limit xx
func (s *SelectStruct) Limit(limit int) {
	s.limit = limit
}

// OffSet xx
func (s *SelectStruct) OffSet(offSet int) {
	s.offSet = offSet
}

// ToString xx
func (s SelectStruct) ToString() string {

	columns := "*"

	if len(s.columns) > 0 {
		columns = strings.Join(s.columns, ", ")
	}

	sql := fmt.Sprintf("SELECT %s", columns)

	if s.tableName == "" {
		// 表名字没有设置 后面where 之类的拼接是没有意义的
		return sql
	}

	sql = fmt.Sprintf("%s FROM %s %s", sql, s.tableName, s.where.toString())

	if len(s.groupBy) > 0 {
		sql = fmt.Sprintf("%s GROUP BY %s", sql, strings.Join(s.groupBy, ", "))
		// HAVING 配合 GROUP BY使用
		if len(s.having) > 0 {
			sql = fmt.Sprintf("%s HAVING %s", sql, strings.Join(s.having, " AND "))
		}
	}

	if len(s.orderBy) > 0 {
		sql = fmt.Sprintf("%s ORDER BY %s", sql, strings.Join(s.orderBy, ", "))
	}

	if s.limit > 0 {
		sql = fmt.Sprintf("%s LIMIT %d", sql, s.limit)
		// limit 设置了才会拼接 offset 否则 值设置 offset 没有意义
		if s.offSet > 0 {
			sql = fmt.Sprintf("%s OFFSET %d", sql, s.offSet)
		}
	}

	return sql
}

// Table xx
func (s *SelectStruct) Table(table string) {
	s.tableName = table
}
