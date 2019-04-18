package buildsql

import (
	"fmt"
	"github.com/mujianhong/buildsql"
	"testing"
)

var echoRand = func(str string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", str)
}

func TestNewSelectBuilder(t *testing.T) {
	se := buildsql.NewSelectBuilder()
	if fmt.Sprintf("%T", se) != "*buildsql.SelectBuilder" {
		t.Error("Structural error")
	}
}

func TestToString(t *testing.T) {
	se := buildsql.NewSelectBuilder()
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
		t.Errorf("buildsqlString is %s sql is %s", echoRand(buildSQLString), echoRand(SQL))
	}

}
