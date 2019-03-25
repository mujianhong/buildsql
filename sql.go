package buildsql

type publicAttribute struct {
	tableName string
	where     whereStruct
}

// Select xx
func Select() SelectStruct {
	s := SelectStruct{}
	s.where = whereStruct{}
	return s
}

// Update xx
func Update() UpdateStruct {
	u := UpdateStruct{}
	u.where = whereStruct{}
	return u
}

// Delete XX
func Delete() DeleteStruct {
	d := DeleteStruct{}
	d.where = whereStruct{}
	return d
}

// Insert XX
func Insert() InsertStruct {
	i := InsertStruct{}
	return i
}
