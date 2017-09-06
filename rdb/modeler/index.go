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

const (
	IndexTypeNull       int = 0
	IndexTypeIndex      int = 1
	IndexTypeUnique     int = 2
	IndexTypePrimaryKey int = 3
)

// database index
type Index struct {
	Name string   `json:"name"`
	Type int      `json:"type"`
	Cols []string `json:"cols"`
}

// add columns which will be composite index
func (index *Index) AddColumn(cols ...string) *Index {

	for _, col := range cols {

		exist := false
		for _, v := range index.Cols {
			if v == col {
				exist = true
			}
		}

		if !exist {
			index.Cols = append(index.Cols, col)
		}
	}

	return index
}

// new an index
func NewIndex(name string, indexType int) *Index {
	return &Index{name, indexType, []string{}}
}
