package sqlt

import "reflect"

type tableMeta struct {
	name string
}

// Shared Resource, It's not safe for multi-thread
var tablesMeta map[string]tableMeta

// RegisterTable .
func RegisterTable(o interface{}, tableName string) error {
	return nil
}

func registerCURDSql(o interface{}) {

}

//SQLAutoMapper .
type SQLAutoMapper struct {
	tpl SQLTemplate
}

//Save .
func (am *SQLAutoMapper) Save(o interface{}) error {
	return nil
}

//FindAll .
func (am *SQLAutoMapper) Find(t reflect.Type, subSQL string, args ...interface{}) ([]interface{}, error) {
	return nil, nil
}
