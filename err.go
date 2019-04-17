package buildsql

import (
	"errors"
)

var (
	// ErrTableNameNotFound 表名没有设置
	ErrTableNameNotFound = errors.New("Table name does not exist")
	// ErrDatabaseTypeNotSpecified 没有指定数据库类型
	ErrDatabaseTypeNotSpecified = errors.New("No database type is specified")
	// ErrDatabaseTypeSpecified 重复指定数据库类型
	ErrDatabaseTypeSpecified = errors.New("The database type has been specified")
)
