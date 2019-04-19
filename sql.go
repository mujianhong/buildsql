package buildsql

// SQLBuilder Builder接口
type SQLBuilder interface {
	ToString() string
}

func newExprBuilder() expr {
	ex := expr{}
	return ex
}

// NewSelectBuilder 实例化类
func NewSelectBuilder() *selectBuilder {

	s := &selectBuilder{}
	s.Ex = newExprBuilder()
	s.Wex = newWhereExprBuilder()
	s.table = ""
	s.isSetLimit = false
	return s
}
