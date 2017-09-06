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

var columnTypes = map[string]string{
	"bool":            "bool",
	"string":          "varchar(%v)",
	"string-text":     "longtext",
	"date":            "date",
	"datetime":        "datetime",
	"int8":            "tinyint",
	"int16":           "smallint",
	"int32":           "integer",
	"int64":           "bigint",
	"uint8":           "tinyint unsigned",
	"uint16":          "smallint unsigned",
	"uint32":          "integer unsigned",
	"uint64":          "bigint unsigned",
	"float64":         "double precision",
	"float64-decimal": "numeric(%v, %v)",
}

// database column
type Column struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Length   string   `json:"length,omitempty"`
	NullAble bool     `json:"null_able,omitempty"`
	IncrAble bool     `json:"incr_able,omitempty"`
	Default  string   `json:"default,omitempty"`
	Comment  string   `json:"comment,omitempty"`
	Extra    []string `json:"extra,omitempty"`
}

func NewColumn(colName, colType, len string, null bool, def string) *Column {
	return &Column{
		Name:     colName,
		Type:     colType,
		Length:   len,
		NullAble: null,
		IncrAble: false,
		Default:  def,
		Comment:  "",
		Extra:    []string{},
	}
}
