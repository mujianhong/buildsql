package buildsql

import (
	"fmt"
	"strings"
)

type whereStruct struct {
	where []string
}

func (w *whereStruct) setWhere(where string) {
	w.where = append(w.where, fmt.Sprintf("(%s)", where))
}

func (w whereStruct) toString() string {

	if len(w.where) > 0 {
		return fmt.Sprintf("WHERE %s", strings.Join(w.where, " AND "))
	}

	return ""
}
