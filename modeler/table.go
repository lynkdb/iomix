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

type Table struct {
	Name    string    `json:"name"`
	Engine  string    `json:"engine,omitempty"`
	Charset string    `json:"charset,omitempty"`
	Columns []*Column `json:"columns"`
	Indexes []*Index  `json:"indexes"`
	Comment string    `json:"comment,omitempty"`
}

func NewTable(name, engine, charset string) *Table {
	return &Table{
		Name:    name,
		Engine:  engine,
		Charset: charset,
		Columns: []*Column{},
		Indexes: []*Index{},
	}
}

func (table *Table) AddColumn(col *Column) {

	for k, v := range table.Columns {

		if v.Name != col.Name {
			continue
		}

		table.Columns[k] = col
		return
	}

	table.Columns = append(table.Columns, col)
}

func (table *Table) AddIndex(index *Index) {

	for k, v := range table.Indexes {

		if v.Name != index.Name {
			continue
		}

		table.Indexes[k] = index
		return
	}

	table.Indexes = append(table.Indexes, index)
}
