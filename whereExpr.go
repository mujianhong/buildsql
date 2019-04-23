package buildsql

import (
	"fmt"
)

const (
	equal        = "="
	notEqual     = "!="
	greater      = ">"
	greaterEqual = ">="
	less         = "<"
	lessEqual    = "<="
	between      = "BETWEEN"
	and          = "AND"
	like         = "LIKE"
	not          = "NOT"
	in           = "IN"
	or           = "OR"
	null         = "NULL"
	is           = "IS"
)

type whereExpr struct {
}

func newWhereExprBuilder() whereExpr {
	ex := whereExpr{}
	return ex
}

// Equal 等于
func (ex whereExpr) Equal(column string, value interface{}) string {

	return ex.toString(column, equal, value)
}

// NotEqual 不等于
func (ex whereExpr) NotEqual(column string, value interface{}) string {

	return ex.toString(column, notEqual, value)
}

// GreaterThan 大于 column > value
func (ex whereExpr) GreaterThan(column string, value interface{}) string {

	return ex.toString(column, greater, value)
}

// GreaterEqualThan 大于等于 column >= value
func (ex whereExpr) GreaterEqualThan(column string, value interface{}) string {

	return ex.toString(column, greaterEqual, value)
}

// LessThanThan 小于 column < value
func (ex whereExpr) LessThanThan(column string, value interface{}) string {

	return ex.toString(column, less, value)
}

// LessThanEqualThan 小于等于 column <= value
func (ex whereExpr) LessThanEqualThan(column string, value interface{}) string {

	return ex.toString(column, lessEqual, value)
}

// Between 在某个范围内 [column] Between [greater] AND [less]
func (ex whereExpr) Between(column string, greater, less interface{}) string {

	return fmt.Sprintf("%s %s %s %s %s", column, between, conversion(greater), and, conversion(less))
}

// NotBetween 不在某个范围内 [column] Between [greater] AND [less]
func (ex whereExpr) NotBetween(column string, greater, less interface{}) string {

	return fmt.Sprintf("%s %s %s %s %s %s", column, not, between, conversion(greater), and, conversion(less))
}

// LIKE 搜索某种模式 column like "value"
func (ex whereExpr) Like(column, value string) string {

	return fmt.Sprintf("%s %s '%s'", column, like, value)
}

// NotLike return string 'column not like "value"'
func (ex whereExpr) NotLike(column, value string) string {

	return fmt.Sprintf("%s %s %s '%s'", column, not, like, value)
}

// In in查询
func (ex whereExpr) In(column string, value interface{}) string {

	return fmt.Sprintf("%s %s (%s)", column, in, conversion(value))
}

// NotIn in查询
func (ex whereExpr) NotIn(column string, value interface{}) string {

	return fmt.Sprintf("%s %s %s (%s)", column, not, in, conversion(value))
}

// OR in查询
// whereExpr.Or(whereExpr.Equal("xxx" , 1), whereExpr.Equal("xxx" , 2))
// @return xxx = 1 OR xxx = 2
func (ex whereExpr) Or(ex1, ex2 string) string {

	return fmt.Sprintf("%s %s %s", ex1, or, ex2)
}

// IsNull return 'column IS NULL'
func (ex whereExpr) IsNull(column string) string {

	return fmt.Sprintf("%s %s %s", column, is, null)
}

// NotIsNull return 'column NOT IS NULL'
func (ex whereExpr) NotIsNull(column string) string {

	return fmt.Sprintf("%s %s %s %s", column, not, is, null)

}

func (ex whereExpr) toString(column, operator string, value interface{}) string {

	return fmt.Sprintf("%s %s %s", column, operator, conversion(value))
}

