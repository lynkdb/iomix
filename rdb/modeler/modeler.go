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
	SchemaDump() (*Schema, error)
	SchemaSync(ds *Schema) error
	SchemaSyncByJson(js string) error
	SchemaSyncByJsonFile(path string) error

	TableDump() ([]*Table, error)
	TableSync(table *Table) error
	TableExist(tableName string) bool

	ColumnDump(tableName string) ([]*Column, error)
	ColumnSync(tableName string, col *Column) error
	ColumnDel(tableName string, col *Column) error
	ColumnSet(tableName string, col *Column) error
	ColumnTypeSql(tableName string, col *Column) string

	IndexDump(tableName string) ([]*Index, error)
	IndexSync(tableName string, index *Index) error
	IndexDel(tableName string, index *Index) error
	IndexSet(tableName string, index *Index) error

	QuoteStr(str string) string
}
