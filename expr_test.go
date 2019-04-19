package buildsql

import (
	"fmt"
	"testing"
)

func TestExpr_newExprBuilder(t *testing.T) {

	ex := newExprBuilder()
	if fmt.Sprintf("%T", ex) != "buildsql.expr" {
		t.Skipf("Structural error Struc type is %T", ex)
	}
}

func TestExpr_Avg(t *testing.T) {

	ex := newExprBuilder()

	str := ex.Avg("gold", "g")

	if str != "AVG(gold) AS g" {
		t.Errorf("Avg() is %s hope is AVG(gold) AS g", str)
	}

	str = ex.Avg("gold", "")

	if str != "AVG(gold)" {
		t.Errorf("Avg() is %s hope is AVG(gold)", str)
	}
}

func TestExpr_Count(t *testing.T) {

	ex := newExprBuilder()

	str := ex.Count("gold", "g")

	if str != "COUNT(gold) AS g" {
		t.Errorf("Count() is %s hope is COUNT(gold) AS g", str)
	}

	str = ex.Count("gold", "")

	if str != "COUNT(gold)" {
		t.Errorf("Count() is %s hope is COUNT(gold)", str)
	}
}

func TestExpr_Max(t *testing.T) {

	ex := newExprBuilder()

	str := ex.Max("gold", "g")

	if str != "MAX(gold) AS g" {
		t.Errorf("Max() is %s hope is MAX(gold) AS g", str)
	}

	str = ex.Max("gold", "")

	if str != "MAX(gold)" {
		t.Errorf("Max() is %s hope is MAX(gold)", str)
	}
}

func TestExpr_Min(t *testing.T) {

	ex := newExprBuilder()

	str := ex.Min("gold", "g")

	if str != "Min(gold) AS g" {
		t.Errorf("Min() is %s hope is Min(gold) AS g", str)
	}

	str = ex.Min("gold", "")

	if str != "Min(gold)" {
		t.Errorf("Min() is %s hope is Min(gold)", str)
	}
}

func TestExpr_Sum(t *testing.T) {

	ex := newExprBuilder()

	str := ex.Sum("gold", "g")

	if str != "SUM(gold) AS g" {
		t.Errorf("Sum() is %s hope is SUM(gold) AS g", str)
	}

	str = ex.Sum("gold", "")

	if str != "SUM(gold)" {
		t.Errorf("Sum() is %s hope is SUM(gold)", str)
	}
}
