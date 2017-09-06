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

type Base struct {
	mu    sync.RWMutex
	db    *sql.DB
	opts  connect.ConnOptions
	stmts map[string]string
}

var baseStmts = map[string]string{
	"insertIgnore": "INSERT OR IGNORE INTO `%s` (`%s`) VALUES (%s)",
}

func NewBase(opts connect.ConnOptions, db *sql.DB) (*Base, error) {
	return &Base{
		opts:  opts,
		db:    db,
		stmts: baseStmts,
	}, nil
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

func (dc *Base) Insert(table_name string, item map[string]interface{}) (Result, error) {

	var res Result

	cols, vars, vals := []string{}, []string{}, []interface{}{}
	for key, val := range item {
		cols = append(cols, key)
		vars = append(vars, "?")
		vals = append(vals, val)
	}

	sql := fmt.Sprintf("INSERT INTO `%s` (`%s`) VALUES (%s)",
		table_name,
		strings.Join(cols, "`,`"),
		strings.Join(vars, ","))

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

func (dc *Base) Delete(table_name string, fr Filter) (Result, error) {

	var res Result

	frsql, params := fr.Parse()
	if len(params) == 0 {
		return res, errors.New("Error in query syntax")
	}

	sql := fmt.Sprintf("DELETE FROM `%s` WHERE %s", table_name, frsql)

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

func (dc *Base) Update(table_name string, item map[string]interface{}, fr Filter) (Result, error) {

	var res Result

	frsql, params := fr.Parse()
	if len(params) == 0 {
		return res, errors.New("Error in query syntax")
	}

	cols, vals := []string{}, []interface{}{}
	for key, val := range item {
		cols = append(cols, "`"+key+"` = ?")
		vals = append(vals, val)
	}

	vals = append(vals, params...)

	sql := fmt.Sprintf("UPDATE `%s` SET %s WHERE %s",
		table_name,
		strings.Join(cols, ","),
		frsql)

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

func (dc *Base) Count(table_name string, fr Filter) (num int64, err error) {

	frsql, params := fr.Parse()
	has_where := "WHERE"
	if len(params) == 0 {
		has_where = ""
	}

	sql := fmt.Sprintf("SELECT COUNT(*) FROM `%s` %s %s", table_name, has_where, frsql)

	stmt, err := dc.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(params...)
	err = row.Scan(&num)

	return num, err
}

func (dc *Base) InsertIgnore(table_name string, item map[string]interface{}) (Result, error) {

	var res Result

	sqlstmt, ok := dc.stmts["insertIgnore"]
	if !ok {
		return res, errors.New("CurdStmt:insertIgnore missing")
	}

	cols, vars, vals := []string{}, []string{}, []interface{}{}
	for key, val := range item {
		cols = append(cols, key)
		vars = append(vars, "?")
		vals = append(vals, val)
	}

	sql := fmt.Sprintf(sqlstmt, table_name, strings.Join(cols, "`,`"), strings.Join(vars, ","))
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

func (dc *Base) QueryRaw(sql string, params ...interface{}) (rs []Entry, err error) {

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

		entry := Entry{Fields: map[string]*Field{}}

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

func (dc *Base) Query(q *QuerySet) (rs []Entry, err error) {

	sql, params := q.Parse()
	if len(params) == 0 {
		return rs, errors.New("Error in query syntax")
	}

	return dc.QueryRaw(sql, params...)
}

func (dc *Base) Fetch(q *QuerySet) (Entry, error) {

	q.Limit(1)

	entry := Entry{Fields: map[string]*Field{}}

	sql, params := q.Parse()
	if len(params) == 0 {
		return entry, errors.New("Error in query syntax")
	}

	rs, err := dc.QueryRaw(sql, params...)
	if err != nil {
		return entry, err
	}

	if len(rs) > 0 {
		return rs[0], nil
	}

	return entry, errors.New("Entry Not Found")
}

func (dc *Base) ExecRaw(query string, args ...interface{}) (Result, error) {
	return dc.db.Exec(query, args...)
}

func (dc *Base) Modeler() (modeler.Modeler, error) {
	return nil, errors.New("No Modeler INIT")
}
