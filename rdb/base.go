// Copyright 2014 Eryx <evorui аt gmаil dοt cοm>, All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rdb

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/lynkdb/iomix/connect"
	"github.com/lynkdb/iomix/rdb/modeler"
)

type Result sql.Result

type BaseStmtBindVarFunc func(sql string, vars []interface{}) (string, []interface{})
type BaseQuoteStrFunc func(name string) string

type Base struct {
	mu              sync.RWMutex
	db              *sql.DB
	opts            connect.ConnOptions
	stmts           map[string]string
	BindVar         BaseStmtBindVarFunc
	QuoteStr        BaseQuoteStrFunc
	TypeDatetimeFmt string
}

var baseStmts = map[string]string{
	"insert":       "INSERT INTO %s (%s) VALUES (%s)",
	"insertIgnore": "INSERT OR IGNORE INTO %s (%s) VALUES (%s)",
	"delete":       "DELETE FROM %s WHERE %s",
	"update":       "UPDATE %s SET %s WHERE %s",
	"count":        "SELECT COUNT(*) FROM %s %s %s",
}

func NewBase(opts connect.ConnOptions, db *sql.DB) (*Base, error) {
	b := &Base{
		opts:  opts,
		db:    db,
		stmts: map[string]string{},
		QuoteStr: func(name string) string {
			if name == "*" {
				return name
			}
			return "`" + name + "`"
		},
	}
	for k, v := range baseStmts {
		b.stmts[k] = v
	}
	return b, nil
}

func (dc *Base) Setup(opts connect.ConnOptions, conn *sql.DB) error {
	dc.opts = opts
	dc.db = conn
	return nil
}

func (dc *Base) StmtSet(name, value string) {

	dc.mu.Lock()
	defer dc.mu.Unlock()

	dc.stmts[name] = value
}

func (dc *Base) Options() *connect.ConnOptions {
	return &dc.opts
}

func (dc *Base) DB() *sql.DB {
	return dc.db
}

func (dc *Base) Insert(tableName string, item map[string]interface{}) (Result, error) {

	var res Result

	cols, vars, vals := []string{}, []string{}, []interface{}{}
	for key, val := range item {
		cols = append(cols, dc.QuoteStr(key))
		vars = append(vars, "?")
		vals = append(vals, val)
	}

	sqlstmt, ok := dc.stmts["insert"]
	if !ok {
		return res, errors.New("CurdStmt:insert missing")
	}

	sql := fmt.Sprintf(sqlstmt,
		dc.QuoteStr(tableName),
		strings.Join(cols, ","),
		strings.Join(vars, ","))

	if dc.BindVar != nil {
		sql, vals = dc.BindVar(sql, vals)
	}

	stmt, err := dc.db.Prepare(sql)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	res, err = stmt.Exec(vals...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (dc *Base) Delete(tableName string, fr Filter) (Result, error) {

	var res Result

	frsql, params := fr.Parse()
	if len(params) == 0 {
		return res, errors.New("Error in query syntax")
	}

	sqlstmt, ok := dc.stmts["delete"]
	if !ok {
		return res, errors.New("CurdStmt:delete missing")
	}

	sql := fmt.Sprintf(sqlstmt, tableName, frsql)

	if dc.BindVar != nil {
		sql, params = dc.BindVar(sql, params)
	}

	stmt, err := dc.db.Prepare(sql)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(params...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (dc *Base) Update(tableName string, item map[string]interface{}, fr Filter) (Result, error) {

	var res Result

	frsql, params := fr.Parse()
	if len(params) == 0 {
		return res, errors.New("Error in query syntax")
	}

	cols, vals := []string{}, []interface{}{}
	for key, val := range item {
		cols = append(cols, dc.QuoteStr(key)+" = ?")
		vals = append(vals, val)
	}

	vals = append(vals, params...)

	sqlstmt, ok := dc.stmts["update"]
	if !ok {
		return res, errors.New("CurdStmt:update missing")
	}

	sql := fmt.Sprintf(sqlstmt,
		tableName,
		strings.Join(cols, ","),
		frsql)

	if dc.BindVar != nil {
		sql, vals = dc.BindVar(sql, vals)
	}

	stmt, err := dc.db.Prepare(sql)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vals...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (dc *Base) Count(tableName string, fr Filter) (num int64, err error) {

	frsql, params := fr.Parse()
	has_where := "WHERE"
	if len(params) == 0 {
		has_where = ""
	}

	sqlstmt, ok := dc.stmts["count"]
	if !ok {
		return 0, errors.New("CurdStmt:update missing")
	}

	sql := fmt.Sprintf(sqlstmt, tableName, has_where, frsql)

	if dc.BindVar != nil {
		sql, params = dc.BindVar(sql, params)
	}

	stmt, err := dc.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(params...)
	err = row.Scan(&num)

	return num, err
}

func (dc *Base) InsertIgnore(tableName string, item map[string]interface{}) (Result, error) {

	var res Result

	sqlstmt, ok := dc.stmts["insertIgnore"]
	if !ok {
		return res, errors.New("CurdStmt:insertIgnore missing")
	}

	cols, vars, vals := []string{}, []string{}, []interface{}{}
	for key, val := range item {
		cols = append(cols, dc.QuoteStr(key))
		vars = append(vars, "?")
		vals = append(vals, val)
	}

	sql := fmt.Sprintf(sqlstmt, tableName, strings.Join(cols, ","), strings.Join(vars, ","))

	if dc.BindVar != nil {
		sql, vals = dc.BindVar(sql, vals)
	}

	stmt, err := dc.db.Prepare(sql)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	res, err = stmt.Exec(vals...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (dc *Base) BatchInsertIgnore(tableName string, cols []string, values ...interface{}) (Result, error) {

	var res Result

	if len(cols) == 0 || len(values) == 0 {
		return res, errors.New("no data request")
	}

	if (len(values) % len(cols)) != 0 {
		return res, errors.New("invalid data request")
	}

	sqlstmt, ok := dc.stmts["insertIgnore"]
	if !ok {
		return res, errors.New("CurdStmt:insertIgnore missing")
	}

	colx := ""
	for _, v := range cols {
		if colx != "" {
			colx += ","
		}
		colx += dc.QuoteStr(v)
	}

	varx := "?"
	if len(cols) > 1 {
		varx += strings.Repeat(",?", len(cols)-1)
	}
	if n := (len(values) / len(cols)); n > 1 {
		varx += strings.Repeat("),("+varx, n-1)
	}

	sql := fmt.Sprintf(sqlstmt, tableName, colx, varx)

	if dc.BindVar != nil {
		sql, values = dc.BindVar(sql, values)
	}

	stmt, err := dc.db.Prepare(sql)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	res, err = stmt.Exec(values...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (dc *Base) QueryRaw(sql string, params ...interface{}) (rs []Entry, err error) {

	if dc.BindVar != nil {
		sql, params = dc.BindVar(sql, params)
	}

	stmt, err := dc.db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(params...)
	if err != nil {
		return
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return
	}

	for rows.Next() {

		entry := Entry{
			Fields:       map[string]*Field{},
			datetime_fmt: dc.TypeDatetimeFmt,
		}

		var retvals []interface{}
		for i := 0; i < len(cols); i++ {
			var val interface{}
			retvals = append(retvals, &val)
		}

		if err := rows.Scan(retvals...); err != nil {
			continue
		}

		for ii, key := range cols {

			rawValue := reflect.Indirect(reflect.ValueOf(retvals[ii]))
			if rawValue.Interface() == nil {
				continue
			}

			entry.Fields[key] = &Field{
				valueType: reflect.TypeOf(rawValue.Interface()),
				value:     rawValue,
			}
		}

		rs = append(rs, entry)
	}

	return
}

func (dc *Base) Query(q Queryer) (rs []Entry, err error) {

	sql, params := q.Parse()
	if len(params) == 0 {
		return rs, errors.New("Error in query syntax")
	}

	if dc.BindVar != nil {
		sql, params = dc.BindVar(sql, params)
	}

	return dc.QueryRaw(sql, params...)
}

func (dc *Base) Fetch(q Queryer) (Entry, error) {

	q.Limit(1)

	entry := Entry{
		Fields:       map[string]*Field{},
		datetime_fmt: dc.TypeDatetimeFmt,
	}

	sql, params := q.Parse()
	if len(params) == 0 {
		return entry, errors.New("Error in query syntax")
	}

	if dc.BindVar != nil {
		sql, params = dc.BindVar(sql, params)
	}

	rs, err := dc.QueryRaw(sql, params...)
	if err != nil {
		return entry, err
	}

	if len(rs) == 0 {
		entry.status = status_not_found
		return entry, errors.New("Entry Not Found")
	}

	return rs[0], nil
}

func (dc *Base) ExecRaw(query string, args ...interface{}) (Result, error) {
	return dc.db.Exec(query, args...)
}

func (dc *Base) Modeler() (modeler.Modeler, error) {
	return nil, errors.New("No Modeler INIT")
}

func (dc *Base) Close() {
	if dc.db != nil {
		dc.db.Close()
	}
}
