package buildsql

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	equal         = "="
	not_equal     = "!="
	greater       = ">"
	greater_equal = ">="
	less          = "<"
	less_equal    = "<="
	between       = "BETWEEN"
	and           = "AND"
	like          = "LIKE"
	not           = "NOT"
	in            = "IN"
	or = "OR"
	null = "NULL"
	is = "IS"
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

	return ex.toString(column, not_equal, value)
}

// GreaterThan 大于 column > value
func (ex whereExpr) GreaterThan(column string, value interface{}) string {

	return ex.toString(column, greater, value)
}

// GreaterEqualThan 大于等于 column >= value
func (ex whereExpr) GreaterEqualThan(column string, value interface{}) string {

	return ex.toString(column, greater_equal, value)
}

// LessThanThan 小于 column < value
func (ex whereExpr) LessThanThan(column string, value interface{}) string {

	return ex.toString(column, less, value)
}

// LessThanEqualThan 小于等于 column >= value
func (ex whereExpr) LessThanEqualThan(column string, value interface{}) string {

	return ex.toString(column, less_equal, value)
}

// Between 在某个范围内 [column] Between [greater] AND [less]
func (ex whereExpr) Between(column string, greater, less interface{}) string {

	return fmt.Sprintf("%s %s %s %s %s", column, between, ex.conversion(greater), and, ex.conversion(less))
}

// NotBetween 不在某个范围内 [column] Between [greater] AND [less]
func (ex whereExpr) NotBetween(column string, greater, less interface{}) string {

	return fmt.Sprintf("%s %s %s %s %s %s", column, not, between, ex.conversion(greater), and, ex.conversion(less))
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

	return fmt.Sprintf("%s %s %s", column, in, ex.conversion(value))
}

// NotIn in查询
func (ex whereExpr) NotIn(column string, value interface{}) string {

	return fmt.Sprintf("%s %s %s %s", column, not, in, ex.conversion(value))
}

// OR in查询
// whereExpr.Or(whereExpr.Equal("xxx" , 1), whereExpr.Equal("xxx" , 2))
// @return xxx = 1 OR xxx = 2
func (ex whereExpr) Or(ex1 ,ex2 string) string {

	return fmt.Sprintf("%s %s %s", ex1, or, ex2)
}
// IsNull return 'column IS NULL'
func (ex whereExpr) IsNull(column string) string {

	return fmt.Sprintf("%s %s %s", column , is, null)
}

// NotIsNull return 'column NOT IS NULL'
func (ex whereExpr) NotIsNull(column string) string {

	return fmt.Sprintf("%s %s %s %s", column, not, is, null)

}
func (ex whereExpr) toString(column, operator string, value interface{}) string {

	return fmt.Sprintf("%s %s %s", column, operator, ex.conversion(value))
}

func (ex whereExpr) conversion(value interface{}) string {

	switch value.(type) {
	case bool:
		if value == true {
			return "true"
		}
		return "false"
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", value)
	case float32, float64:
		return fmt.Sprintf("%f", value)
	case string:
		return fmt.Sprintf("'%s'", value)
	case []string:
		arr := []string{}
		for _, v := range value.([]string) {
			arr = append(arr, fmt.Sprintf("'%s'", v))
		}
		return fmt.Sprintf("(%s)", strings.Join(arr, ", "))
	case []uint, []uint8, []uint16, []uint32, []uint64, []int, []int8, []int16, []int32, []int64:
		arr := []string{}
		for _, v := range value.([]int) {
			arr = append(arr, strconv.Itoa(v))
		}
		return fmt.Sprintf("(%s)", strings.Join(arr, ", "))
	case *SelectBuilder:
		s := value.(*SelectBuilder)
		return fmt.Sprintf("(%s)", s.ToString())
	}

	return ""
}
