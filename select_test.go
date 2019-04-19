package buildsql

import (
	"fmt"
	"testing"
)

func TestNewSelectBuilder(t *testing.T) {

	se := NewSelectBuilder()
	if fmt.Sprintf("%T", se) != "*buildsql.selectBuilder" {
		t.Skipf("Structural error Struc type is %T", se)
	}
}

func TestSelectBuilder_As(t *testing.T) {

	se := NewSelectBuilder()

	se.As("xxx", "x")

	if len(se.columns) != 1 {
		t.Error()
		return
	}

	if se.columns[0] != "xxx AS x" {
		t.Error("SelectBuilder As function assembly error")
		return
	}
}

func TestSelectBuilder_GroupBy(t *testing.T) {

	se := NewSelectBuilder()

	se.GroupBy("user_id", "create_time")

	if len(se.groupBy) != 2 {
		t.Error()
		return
	}

	gr := []string{"user_id", "create_time"}

	for k, v := range se.groupBy {
		if gr[k] != v {
			t.Error()
			return
		}
	}
}

func TestSelectBuilder_Having(t *testing.T) {

	se := NewSelectBuilder()
	having := []string{"user_id = 1", "user_id = 2"}

	for _, v := range having {
		se.Having(v)
	}

	if len(se.having) != len(having) {
		t.Error()
		return
	}

	for k, v := range se.having {
		if fmt.Sprintf("(%s)", having[k]) != v {
			t.Errorf("having[%d] is %s selectBuilder.having[%d] is %s", k, having[k], k, v)
		}
	}
}

func TestSelectBuilder_Limit(t *testing.T) {

	se := NewSelectBuilder()
	limit := 1

	if se.limit != 0 {
		t.Error()
		return
	}
	se.Limit(limit)

	if se.limit != limit {
		t.Errorf(`limit{"type": %T, "value": %d} selectBuilder.limit{"type": %T, "value": %d}`, limit, limit, se.limit, se.limit)
		return
	}
}

func TestSelectBuilder_Offset(t *testing.T) {

	se := NewSelectBuilder()
	offset := 1

	if se.offset != 0 {
		t.Error()
		return
	}
	se.Offset(offset)

	if se.offset != offset {
		t.Errorf(`offset{"type": %T, "value": %d} selectBuilder.offset{"type": %T, "value": %d}`, offset, offset, se.offset, se.offset)
		return
	}
}

func TestSelectBuilder_SetPage(t *testing.T) {
	se := NewSelectBuilder()
	limit := 20
	offset := 40

	if se.offset != 0 || se.limit != 0 {
		t.Error()
		return
	}

	se.SetPage(limit, offset)

	if se.offset != offset {
		t.Errorf(`offset{"type": %T, "value": %d} selectBuilder.offset{"type": %T, "value": %d}`, offset, offset, se.offset, se.offset)
		return
	}

	if se.limit != limit {
		t.Errorf(`limit{"type": %T, "value": %d} selectBuilder.limit{"type": %T, "value": %d}`, limit, limit, se.limit, se.limit)
		return
	}
}

func TestSelectBuilder_OrderBy(t *testing.T) {

	se := NewSelectBuilder()
	order := []string{"user_id", "create_time"}

	if len(se.orderBy) > 0 {
		t.Error("se.orderBy value is", se.orderBy)
		return
	}

	for _, v := range order {
		se.OrderBy(v)
	}

	if len(se.orderBy) != len(order) {
		t.Errorf("order size is %d se.orderBy size is %d", len(order), len(se.orderBy))
	}

	for k, v := range order {
		if se.orderBy[k] != fmt.Sprintf("%s %s", v, "ASC") {
			t.Errorf("orderBy[k] is %s hope is %s", se.orderBy[k], fmt.Sprintf("%s%s", v, "ASC"))
		}
	}
}

func TestSelectBuilder_OrderByDesc(t *testing.T) {

	se := NewSelectBuilder()
	order := []string{"user_id", "create_time"}

	if len(se.orderBy) > 0 {
		t.Error("se.orderBy value is", se.orderBy)
		return
	}

	for _, v := range order {
		se.OrderByDesc(v)
	}

	if len(se.orderBy) != len(order) {
		t.Errorf("order size is %d se.orderBy size is %d", len(order), len(se.orderBy))
	}

	for k, v := range order {
		if se.orderBy[k] != fmt.Sprintf("%s %s", v, "DESC") {
			t.Errorf("orderBy[k] is %s hope is %s", se.orderBy[k], fmt.Sprintf("%s%s", v, "DESC"))
		}
	}
}

func TestSelectBuilder_SetColumn(t *testing.T) {

	se := NewSelectBuilder()

	if len(se.columns) > 0 {
		t.Error("se.columns value is", se.columns)
		return
	}
	// 直接获取测试默认值 *
	sql := se.ToString()

	if sql != "SELECT * FROM" {
		t.Errorf("se.ToString is %s hope is SELECT * FROM", sql)
	}

	column := []string{"user_id", "create_time"}

	for _, v := range column {
		se.SetColumn(v)
	}

	se.SetColumn("portrait", "nickname")
	column = append(column, "portrait")
	column = append(column, "nickname")

	if len(se.columns) != 4 {
		t.Errorf("se.columns size is %d", len(se.columns))
		return
	}

	for k, v := range se.columns {
		if column[k] != v {
			t.Errorf("se.columns[%d] is %s hope is %s", k, v, column[k])
		}
	}
}

func TestSelectBuilder_SetTable(t *testing.T) {
	se := NewSelectBuilder()

	if se.table != "" {
		t.Errorf("se.table is %s hope is ''", se.table)
		return
	}

	table := "users"

	se.SetTable(table)

	if se.table != table {
		t.Errorf("se.table is %s hope is %s", se.table, table)
	}
}

func TestSelectBuilder_Where(t *testing.T) {

	se := NewSelectBuilder()

	if len(se.wheres) != 0 {
		t.Errorf("se.wheres size %d hope is 0", len(se.wheres))
		return
	}

	wheres := []string{
		"user_id = 1",
		"create_time = 2",
	}

	for _, v := range wheres {
		se.Where(v)
	}
	se.Where("nickname = 'xx'", "portrait = 'https://github.com/mujianhong/buildsql'")
	wheres = append(wheres, "nickname = 'xx'")
	wheres = append(wheres, "portrait = 'https://github.com/mujianhong/buildsql'")

	if len(se.wheres) != len(wheres) {
		t.Errorf("se.wheres size is %d hope is %d", len(se.wheres), len(wheres))
		return
	}

	for k, v := range se.wheres {
		if fmt.Sprintf("(%s)", wheres[k]) != v {
			t.Errorf("se.wheres[%d] is %s hope is %s", k, v, fmt.Sprintf("(%s)", wheres[k]))
		}
	}

}

func TestSelectBuilder_ToString(t *testing.T) {

	se := NewSelectBuilder()
	if s := se.ToString(); s != "SELECT * FROM" {
		t.Errorf("ToString() is %s hope 'SELECT * FROM'", s)
	}

	se = NewSelectBuilder()
	table := "test"
	SQL := fmt.Sprintf("SELECT xxx AS x, user_id, create_time, nickname FROM %s WHERE (user_id = 1) AND (nickname = 'xxx') AND (create_time >= 1) AND (create_time <= 2) GROUP BY user_id, create_time ORDER BY user_id DESC, create_time ASC LIMIT 0 OFFSET 100", table)

	se.SetTable(table).
		Where(se.Wex.Equal("user_id", 1)).
		Where("nickname = 'xxx'").
		Where(se.Wex.GreaterEqualThan("create_time", 1), se.Wex.LessThanEqualThan("create_time", 2)).
		GroupBy("user_id", "create_time")
	se.OrderByDesc("user_id").OrderBy("create_time").Limit(0).Offset(100)
	se.As("xxx", "x")
	se.SetColumn("user_id")
	se.SetColumn("create_time", "nickname")

	if buildSQLString := se.ToString(); buildSQLString != SQL {
		t.Errorf("buildsqlString is %s hope is '%s'", buildSQLString, SQL)
	}
}
