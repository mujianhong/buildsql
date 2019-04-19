package buildsql

import (
	"fmt"
	"strings"
)

// selectBuilder select build 类
type selectBuilder struct {
	// SQLBuilder sql绑定接口
	SQLBuilder
	// SQLStruct sql 公共字段
	sqlStruct
	// Ex 聚合函数
	Ex expr
	// WhereExpr
	Wex whereExpr
	// orderBy 排序
	orderBy []string
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

// ToString 获取sql语句
func (s selectBuilder) ToString() string {

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

		orderBy = fmt.Sprintf("ORDER BY %s", strings.Join(s.orderBy, ", "))
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
func (s *selectBuilder) Limit(limit int) *selectBuilder {

	s.limit = limit
	s.isSetLimit = true
	return s
}

// Offset 设置分页结束 同 Limit 组合使用
func (s *selectBuilder) Offset(offset int) *selectBuilder {

	s.offset = offset
	return s
}

// SetPage 设置分页
func (s *selectBuilder) SetPage(limit, offset int) *selectBuilder {

	return s.Limit(limit).Offset(offset)
}

// GroupBy 分组
func (s *selectBuilder) GroupBy(column ...string) *selectBuilder {

	if len(column) > 0 {
		for _, v := range column {
			s.groupBy = append(s.groupBy, v)
		}
	}
	return s
}

// Having Having添加 同group by组合使用
func (s *selectBuilder) Having(column ...string) *selectBuilder {

	if len(column) > 0 {
		for _, v := range column {
			s.having = append(s.having, fmt.Sprintf("(%s)", v))
		}
	}
	return s
}

// OrderBy 排序 正叙
func (s *selectBuilder) OrderBy(column ...string) *selectBuilder {

	if len(column) > 0 {
		for _, v := range column {
			s.orderBy = append(s.orderBy, fmt.Sprintf("%s ASC", v))
		}
	}
	return s
}

// OrderByDesc 排序 倒叙
func (s *selectBuilder) OrderByDesc(column ...string) *selectBuilder {

	if len(column) > 0 {
		for _, v := range column {
			s.orderBy = append(s.orderBy, fmt.Sprintf("%s DESC", v))
		}
	}
	return s
}

// SetTable 设置表
func (s *selectBuilder) SetTable(table string) *selectBuilder {

	s.table = table
	return s
}

// SetColumn 设置select的字段
func (s *selectBuilder) SetColumn(column ...string) *selectBuilder {

	if len(column) > 0 {
		for _, v := range column {
			s.columns = append(s.columns, v)
		}
	}
	return s
}

// As 设置字段别名
func (s *selectBuilder) As(column, alias string) *selectBuilder {

	s.SetColumn(fmt.Sprintf("%s AS %s", column, alias))
	return s
}

// Where 条件设置
func (s *selectBuilder) Where(expr ...string) *selectBuilder {

	if len(expr) > 0 {
		for _, v := range expr {
			if v != "" {
				s.wheres = append(s.wheres, fmt.Sprintf("(%s)", v))
			}
		}
	}
	return s
}
