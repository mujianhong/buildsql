package buildsql

// SQLBuilder Builder接口
type SQLBuilder interface {
	ToString() string
}

// NewSelectBuilder 实例化类
func NewSelectBuilder() *selectBuilder {

	s := &selectBuilder{}
	s.Ex = newExprBuilder()
	s.Wex = newWhereExprBuilder()
	s.table = ""
	s.orderBy = map[string]string{}
	s.isSetLimit = false
	return s
}
