package buildsql

import (
	"fmt"
	"testing"
)

func TestNewInsertBuilder(t *testing.T) {

	insert := NewInsertBuilder()

	fmt.Println(insert.ToString())
}
