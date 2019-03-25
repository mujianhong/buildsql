package buildsql

import (
	"fmt"
	"strings"
	"testing"
)

func TestDeleteTable(t *testing.T) {
	delete := Delete()
	delete.Table("test_table")

	if delete.tableName != "test_table" {
		t.Error("table set up the ")
	}
}

func TestDeleteToString(t *testing.T) {
	delete := Delete()

	if delete.ToString() != "" {
		t.Error("Error building delete statement")
	}

	table := "test_stable"
	sql := fmt.Sprintf("DELETE FROM %s", table)
	delete.Table(table)
	buildSQL := strings.TrimSpace(delete.ToString())

	if sql != buildSQL {
		t.Errorf("Error building delete statement sql: %s, buildSQL: %s.", sql, buildSQL)
	}

	delete.Where("user_id = 1")

	sql = fmt.Sprintf("%s WHERE (%s)", sql, "user_id = 1")
	buildSQL = strings.TrimSpace(delete.ToString())

	if sql != buildSQL {
		t.Errorf("Error building delete statement sql: %s, buildSQL: %s.", sql, buildSQL)
	}
}
