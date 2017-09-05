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

package modeler

type Modeler interface {
	IndexAdd(db_name, table_name string, index *Index) error
	IndexDel(db_name, table_name string, index *Index) error
	IndexSet(db_name, table_name string, index *Index) error
	IndexQuery(db_name, table_name string) ([]*Index, error)

	ColumnAdd(db_name, table_name string, col *Column) error
	ColumnDel(db_name, table_name string, col *Column) error
	ColumnSet(db_name, table_name string, col *Column) error
	ColumnQuery(db_name, table_name string) ([]*Column, error)
	ColumnTypeSql(col *Column) string

	TableAdd(table *Table) error
	TableQuery(db_name string) ([]*Table, error)
	TableExist(db_name, table_name string) bool

	Sync(db_name string, ds DatabaseEntry) error
	QuoteStr(str string) string
	DatabaseEntry(db_name string) (DatabaseEntry, error)
}
