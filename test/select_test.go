package buildsql

import (
	"testing"
	"github.com/mujianhong/buildsql"
	"fmt"
)

func TestNewSelectBuilder(t *testing.T) {
	se := buildsql.NewSelectBuilder()
	se.Where(se.Wex.In("xx", []int{1,2,3,4,5}))
	se.Where(se.Wex.NotIn("xx", []string{"dd","cc","fff"}))
	se.Where(se.Wex.In("xx", se))
	se.Where(se.Wex.Or(se.Wex.In("xx", se), se.Wex.NotIn("xx", []string{"dd","cc","fff"})))
	se.Where(se.Wex.IsNull("dddd"))
	se.Where(se.Wex.NotIsNull("dddd"))
	fmt.Println(se)
}