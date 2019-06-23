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

import (
	"fmt"
	"strconv"
	"strings"
)

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
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Length      string   `json:"length,omitempty"`
	NotNullAble bool     `json:"not_null_able,omitempty"`
	IncrAble    bool     `json:"incr_able,omitempty"`
	Default     string   `json:"default,omitempty"`
	Comment     string   `json:"comment,omitempty"`
	Extra       []string `json:"extra,omitempty"`
}

func (it *Column) IsInt() bool {
	if strings.HasPrefix(it.Type, "int") ||
		strings.HasPrefix(it.Type, "uint") {
		return true
	}
	return false
}

func (it *Column) IsFloat() bool {
	if strings.HasPrefix(it.Type, "float") {
		return true
	}
	return false
}

func (it *Column) IsNumber() bool {
	if it.IsInt() || it.IsFloat() {
		return true
	}
	return false
}

func (it *Column) IsChar() bool {
	if strings.HasPrefix(it.Type, "string") {
		return true
	}
	return false
}

func (it *Column) Fix() {
	if it.IsNumber() {
		it.Default = "0"
	}
	if !it.IsInt() {
		it.IncrAble = false
	}
	if it.IncrAble {
		it.Default = ""
	}
	if it.Type == "string" {
		if it.Length == "" {
			it.Length = "20"
		}
	}
	if it.Type == "float64-decimal" {
		if it.Length == "" {
			it.Length = "10,2"
		} else {
			lens := strings.Split(it.Length, ",")
			if lens[0] == "" {
				lens[0] = "10"
			}
			if len(lens) < 2 {
				lens = append(lens, "2")
			}

			numeric_p, _ := strconv.Atoi(lens[0])
			numeric_s, _ := strconv.Atoi(lens[1])

			if numeric_s < 2 {
				numeric_s = 2
			} else if numeric_s > 10 {
				numeric_s = 10
			}

			if numeric_p < 4 {
				numeric_p = 4
			}

			if n := numeric_s - numeric_p; n > 0 {
				numeric_p += n
			}

			it.Length = fmt.Sprintf("%d,%d", numeric_p, numeric_s)
		}
	}
}

func NewColumn(colName, colType, len string, def string) *Column {

	c := &Column{
		Name:        colName,
		Type:        colType,
		Length:      len,
		NotNullAble: false,
		IncrAble:    false,
		Default:     def,
		Comment:     "",
		Extra:       []string{},
	}
	c.Fix()
	return c
}
