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

	"github.com/lynkdb/iomix/connect"
	"github.com/lynkdb/iomix/rdb/modeler"
)

type Queryer interface {
	Select(s string) Queryer
	From(s string) Queryer
	Order(s string) Queryer
	Limit(num int64) Queryer
	Offset(num int64) Queryer
	Parse() (sql string, params []interface{})
	Where() Filter
	SetFilter(fr Filter)
}

type Filter interface {
	And(expr string, args ...interface{}) Filter
	Or(expr string, args ...interface{}) Filter
	Parse() (where string, params []interface{})
}

type Connector interface {
	//
	Insert(table_name string, item map[string]interface{}) (Result, error)
	Delete(table_name string, fr Filter) (Result, error)
	Update(table_name string, item map[string]interface{}, fr Filter) (Result, error)
	Count(table_name string, fr Filter) (num int64, err error)
	InsertIgnore(table_name string, item map[string]interface{}) (Result, error)
	Query(q Queryer) (rs []Entry, err error)
	QueryRaw(sql string, params ...interface{}) (rs []Entry, err error)
	Fetch(q Queryer) (Entry, error)
	ExecRaw(query string, args ...interface{}) (Result, error)

	//
	Options() *connect.ConnOptions
	DB() *sql.DB
	Modeler() (modeler.Modeler, error)

	//
	NewFilter() Filter
	NewQueryer() Queryer
}
