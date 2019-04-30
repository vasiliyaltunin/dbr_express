//Package dbre is gocraft/dbr wrapper that allow you do simple sql query - select, insert, update and delete just in one string
package dbre

import (
	"reflect"

	"github.com/vasiliyaltunin/dbr"
)

//DbrExpressStmt - holds session for dbr
type DbrExpressStmt struct {
	Session *dbr.Session
}

var sess dbr.Session

//DbrExpress - inits session var for dbr
func DbrExpress(session *dbr.Session) *DbrExpressStmt {
	return &DbrExpressStmt{Session: session}
}

//Select - selects data from database
func (s *DbrExpressStmt) Select(table, fields string, data interface{}, whereStr string, whereVal ...interface{}) interface{} {

	var ptr reflect.Value

	ptr = reflect.New(reflect.TypeOf(data)) // create new pointer

	obj := ptr.Interface()

	sess := s.Session

	query := sess.Select(fields).
		From(table)

	if len(whereStr) > 0 {
		query.Where(whereStr, whereVal...)
	}

	_, err := query.Load(obj)

	if err != nil {
		panic("[DBRE] SQL query Select error: " + err.Error())
	}

	return obj

}

//SelectDbr - setup select and return *dbr.SelectStmt
func (s *DbrExpressStmt) SelectDbr(table, fields string, whereStr string, whereVal ...interface{}) *dbr.SelectStmt {

	sess := s.Session

	query := sess.Select(fields).
		From(table)

	if len(whereStr) > 0 {
		query.Where(whereStr, whereVal...)
	}

	return query

}

//Load - exec dbr statement
func (s *DbrExpressStmt) Load(query *dbr.SelectStmt, data interface{}) interface{} {
	var ptr reflect.Value

	ptr = reflect.New(reflect.TypeOf(data)) // create new pointer

	obj := ptr.Interface()

	_, err := query.Load(obj)

	if err != nil {
		panic("[DBRE] SQL query Load error: " + err.Error())
	}

	return obj
}

//Insert - inserts data into database
func (s *DbrExpressStmt) Insert(table string, columns []string, data interface{}) interface{} {

	obj := reflect.ValueOf(data).Interface()

	sess := s.Session

	_, err := sess.InsertInto(table).
		Columns(columns...).
		Record(obj).
		Exec()

	if err != nil {
		panic("[DBRE] SQL query Insert error: " + err.Error())
	}

	// id is set automatically
	return obj

}

//Update - updates data into database
func (s *DbrExpressStmt) Update(table string, columns []string, data interface{}, whereStr string, whereValue ...interface{}) {

	obj := reflect.ValueOf(data).Interface()

	sess := s.Session

	// m := structs.Map(data)

	_, err := sess.Update(table).
		Columns(columns...).
		Record(obj).
		Where(whereStr, whereValue...).
		Exec()

	if err != nil {
		panic("[DBRE] SQL query Update error: " + err.Error())
	}

}

//Delete - updates data into database
func (s *DbrExpressStmt) Delete(table string, whereStr, whereValue string) {

	sess := s.Session

	_, err := sess.DeleteFrom(table).
		Where(whereStr, whereValue).
		Exec()

	if err != nil {
		panic("[DBRE] SQL query Delete error: " + err.Error())
	}

}
