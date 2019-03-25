package buildsql

import (
	"testing"
)

func TestInsertTable(t *testing.T) {
	insert := Insert()
	insert.Table("test_table")

	if insert.tableName != "test_table" {
		t.Error("table set up the ")
	}
}

func TestInsertSet(t *testing.T) {
	insert := Insert()
	insert.Table("test_table")
	value := map[string]interface{}{"id": 123, "nickname": "mm"}

	insert.Set(value)

	if len(insert.columns) != len(insert.values) {
		t.Error("values set up the ")
	}

	if insert.columns[0] != "id" {
		t.Error("column set up the ")
	}

	if insert.columns[1] != "nickname" {
		t.Error("column set up the ")
	}

	if insert.values[0] != "123" {
		t.Error("value set up the ")
	}

	if insert.values[1] != "'mm'" {
		t.Error("value set up the ")
	}
}

func TestInsertToString(t *testing.T) {

	insert := Insert()

	if insert.ToString() != "" {
		t.Error("Error building insert statement")
	}

	insert.Table("test_table")

	if insert.ToString() != "" {
		t.Error("Error building insert statement")
	}

	value := map[string]interface{}{"id": 123, "nickname": "mm", "Boolean": true}

	insert.Set(value)

	sql := "INSERT INTO test_table (id, nickname, Boolean) VALUES (123, 'mm', true)"
	buildSQL := insert.ToString()
	if buildSQL != sql {
		t.Errorf("Error building delete statement sql: %s, buildSQL: %s.", sql, buildSQL)
	}
}
