package buildsql

import (
	"fmt"
	"strings"
)

// SelectBuilder select build 类
type SelectBuilder struct {
	// SQLBuilder sql绑定接口
	SQLBuilder
	// Ex 聚合函数
	Ex expr
	// WhereExpr
	Wex whereExpr
	// table 表名
	table string
	// columns 要获取的字段名称
	columns []string
	// where 条件
	wheres []string
	// orderBy 排序
	orderBy map[string]string
	// groupBy 分组
	groupBy []string
	// having
	having []string
	// limit
	limit int
	// offset
	offset int
	// isSetLimit
	isSetLimit bool
}

// NewSelectBuilder 实例化类
func NewSelectBuilder() *SelectBuilder {

	s := &SelectBuilder{}
	s.Ex = newExprBuilder()
	s.Wex = newWhereExprBuilder()
	s.table = ""
	s.orderBy = map[string]string{}
	s.isSetLimit = false
	return s
}

// ToString 获取sql语句
func (s SelectBuilder) ToString() string {

	sql := "SELECT %s FROM %s %s %s %s %s"

	columns := "*"

	if len(s.columns) > 0 {
		columns = strings.Join(s.columns, ", ")
	}

	where := ""

	if len(s.wheres) > 0 {
		where = fmt.Sprintf("WHERE %s", strings.Join(s.wheres, " AND "))
	}

	groupBy := ""

	if len(s.groupBy) > 0 {
		having := ""

		if len(s.having) > 0 {
			having = fmt.Sprintf("HAVING %s", strings.Join(s.having, " AND "))
		}

		groupBy = strings.TrimSpace(fmt.Sprintf("GROUP BY %s %s", strings.Join(s.groupBy, ", "), having))
	}

	orderBy := ""

	if len(s.orderBy) > 0 {
		orderByArr := []string{}
		for column, order := range s.orderBy {
			orderByArr = append(orderByArr, fmt.Sprintf("%s %s", column, order))
		}

		orderBy = fmt.Sprintf("ORDER BY %s", strings.Join(orderByArr, ", "))
	}

	limit := ""

	if s.isSetLimit {
		limit = fmt.Sprintf("LIMIT %d", s.limit)

		if s.offset > 0 {
			limit = strings.TrimSpace(fmt.Sprintf("%s OFFSET %d", limit, s.offset))
		}
	}

	return strings.TrimSpace(fmt.Sprintf(sql, columns, s.table, where, groupBy, orderBy, limit))
}

// Limit 设置分页开始
func (s *SelectBuilder) Limit(limit int) *SelectBuilder {

	s.limit = limit
	s.isSetLimit = true
	return s
}

// Offset 设置分页结束 同 Limit 组合使用
func (s *SelectBuilder) Offset(offset int) *SelectBuilder {

	s.offset = offset
	return s
}

// SetPage 设置分页
func (s *SelectBuilder) SetPage(limit, offset int) *SelectBuilder {

	return s.Limit(limit).Offset(offset)
}

// GroupBy 分组
func (s *SelectBuilder) GroupBy(column ...string) *SelectBuilder {

	if len(column) > 0 {
		for _, v := range column {
			s.groupBy = append(s.groupBy, v)
		}
	}
	return s
}

// Having Having添加 同group by组合使用
func (s *SelectBuilder) Having(column ...string) *SelectBuilder {

	if len(column) > 0 {
		for _, v := range column {
			s.having = append(s.having, fmt.Sprintf("(%s)", v))
		}
	}
	return s
}

// OrderBy 排序 正叙
func (s *SelectBuilder) OrderBy(column ...string) *SelectBuilder {

	if len(column) > 0 {
		for _, v := range column {
			s.orderBy[v] = "ASC"
		}
	}
	return s
}

// OrderByDesc 排序 倒叙
func (s *SelectBuilder) OrderByDesc(column ...string) *SelectBuilder {

	if len(column) > 0 {
		for _, v := range column {
			s.orderBy[v] = "DESC"
		}
	}
	return s
}

// SetTable 设置表
func (s *SelectBuilder) SetTable(table string) *SelectBuilder {

	s.table = table
	return s
}

// SetColumn 设置select的字段
func (s *SelectBuilder) SetColumn(column ...string) *SelectBuilder {

	if len(column) > 0 {
		for _, v := range column {
			s.columns = append(s.columns, v)
		}
	}
	return s
}

// As 设置字段别名
func (s *SelectBuilder) As(column, alias string) *SelectBuilder {

	s.columns = append(s.columns, fmt.Sprintf("%s AS %s", column, alias))
	return s
}

// Where 条件设置
func (s *SelectBuilder) Where(expr ...string) *SelectBuilder {

	if len(expr) > 0 {
		for _, v := range expr {
			if v != "" {
				s.wheres = append(s.wheres, fmt.Sprintf("(%s)", v))
			}
		}
	}
	return s
}
