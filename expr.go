package buildsql

import (
	"fmt"
)

type expr struct {
}

func newExprBuilder() expr {
	ex := expr{}
	return ex
}

// Sum 聚合函数
func (ex expr) Sum(column, alias string) string {

	if alias != "" {
		return fmt.Sprintf("SUM(%s) AS %s", column, alias)
	}
	return fmt.Sprintf("SUM(%s)", column)
}

// Count 聚合函数
func (ex expr) Count(column, alias string) string {

	if alias != "" {
		return fmt.Sprintf("COUNT(%s)", column)
	}
	return fmt.Sprintf("COUNT(%s) AS %s", column, alias)
}

// Avg 聚合函数
func (ex expr) Avg(column, alias string) string {

	if alias != "" {
		return fmt.Sprintf("AVG(%s)", column)
	}
	return fmt.Sprintf("AVG(%s) AS %s", column, alias)
}

// Max 聚合函数
func (ex expr) Max(column, alias string) string {

	if alias != "" {
		return fmt.Sprintf("Max(%s)", column)
	}
	return fmt.Sprintf("Max(%s) AS %s", column, alias)
}

// Min 聚合函数
func (ex expr) Min(column, alias string) string {

	if alias != "" {
		return fmt.Sprintf("Min(%s)", column)
	}
	return fmt.Sprintf("Min(%s) AS %s", column, alias)
}
