package buildsql

type insertBuilder struct {
	//
	SQLBuilder
	//
	table string
	//
	se selectBuilder
	//
	vaules []string
	//
	columns []string
}

func NewInsertBuilder() *insertBuilder {
	i := &insertBuilder{}
	return i
}

// 获取sql
func (i insertBuilder) ToString() string {

	return ""
}

// SetSelectBuilder 设置 select语句
func (i *insertBuilder) SetSelectBuilder(se selectBuilder) *insertBuilder {

	i.se = se
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
