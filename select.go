package buildsql

import "fmt"

// SelectBuilder select build 类
type SelectBuilder struct {
	// SqlBuilder sql绑定接口
	SqlBuilder
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
}
// NewSelectBuilder 实例化类
func NewSelectBuilder() *SelectBuilder {

	s := &SelectBuilder{}
	s.Ex = newExprBuilder()
	s.Wex = newWhereExprBuilder()
	return s
}
// ToString 获取sql语句
func (s SelectBuilder) ToString() string {
	return "select"
}
// SetTable 设置表
func (s *SelectBuilder) SetTable(table string) *SelectBuilder {

	s.table = table
	return s
}

// SetColumn 设置select的字段
func (s *SelectBuilder) SetColumn(column ...string) *SelectBuilder{

	if len(column) > 0 {
		for _, v := range column{
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
		for _, v := range expr{
			if v != "" {
				s.wheres = append(s.wheres, fmt.Sprintf("(%s)", v))
			}
		}
	}
	return s
}
