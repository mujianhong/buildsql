package buildsql

import (
	"fmt"
	"strconv"
	"strings"
)

// SQLBuilder Builder接口
type SQLBuilder interface {
	ToString() string
}

type sqlStruct struct {
	table string
	// columns
	columns []string
	// where
	wheres []string
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

func NewInsertBuilder() *insertBuilder {
	i := &insertBuilder{}
	i.isSetSelect = false
	return i
}

func conversion(value interface{}) string {

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
		return fmt.Sprintf("%s", strings.Join(arr, ", "))
	case []uint, []uint8, []uint16, []uint32, []uint64, []int, []int8, []int16, []int32, []int64:
		arr := []string{}
		for _, v := range value.([]int) {
			arr = append(arr, strconv.Itoa(v))
		}
		return fmt.Sprintf("%s", strings.Join(arr, ", "))
	case *selectBuilder:
		s := value.(*selectBuilder)
		return fmt.Sprintf("%s", s.ToString())
	}

	return ""
}
