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
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Field struct {
	value        reflect.Value
	valueType    reflect.Type
	datetime_fmt string
}

type Entry struct {
	Fields       map[string]*Field
	datetime_fmt string
}

func (e *Entry) Field(field_name string) *Field {
	if field, ok := e.Fields[field_name]; ok {
		if e.datetime_fmt != "" {
			field.datetime_fmt = e.datetime_fmt
		}
		return field
	}
	return nil
}

func (f *Field) Bytes() []byte {

	if f != nil && f.value.Interface() != nil {

		vv := reflect.ValueOf(f.value.Interface())

		switch f.valueType.Kind() {

		case reflect.Slice:
			if f.valueType.Elem().Kind() == reflect.Uint8 {
				return f.value.Interface().([]byte)
			}

		case reflect.String:
			return []byte(vv.String())
		}
	}

	return []byte{}
}

func (f *Field) String() string {

	if f != nil && f.value.Interface() != nil {

		vv := reflect.ValueOf(f.value.Interface())

		switch f.valueType.Kind() {

		case reflect.Slice:
			if f.valueType.Elem().Kind() == reflect.Uint8 {
				return string(f.value.Interface().([]byte))
			}

		case reflect.String:
			return vv.String()

		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return fmt.Sprintf("%d", vv.Int())

		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return fmt.Sprintf("%d", vv.Uint())

		case reflect.Struct:
			if t, ok := f.value.Interface().(time.Time); ok {
				return t.String()
			}
		}
	}

	return ""
}

// JsonDecode returns the map that marshals from the reply bytes as json in response .
func (f *Field) JsonDecode(v interface{}) error {
	return json.Unmarshal(f.Bytes(), &v)
}

func (f *Field) Int8() int8 {
	return int8(f.Int64())
}

func (f *Field) Int16() int16 {
	return int16(f.Int64())
}

func (f *Field) Int32() int32 {
	return int32(f.Int64())
}

func (f *Field) Int64() int64 {

	if f != nil && f.value.Interface() != nil {

		vv := reflect.ValueOf(f.value.Interface())

		switch f.valueType.Kind() {

		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return vv.Int()

		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return int64(vv.Uint())

		case reflect.Struct:
			if t, ok := f.value.Interface().(time.Time); ok {
				return t.UnixNano()
			}
		}
	}

	return 0
}

func (f *Field) Int() int {
	return int(f.Int64())
}

func (f *Field) Uint8() uint8 {
	return uint8(f.Uint64())
}

func (f *Field) Uint16() uint16 {
	return uint16(f.Uint64())
}

func (f *Field) Uint32() uint32 {
	return uint32(f.Uint64())
}

func (f *Field) Uint64() uint64 {

	if f != nil && f.value.Interface() != nil {

		vv := reflect.ValueOf(f.value.Interface())

		switch f.valueType.Kind() {

		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return uint64(vv.Int())

		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return vv.Uint()

		case reflect.Struct:
			if t, ok := f.value.Interface().(time.Time); ok {
				return uint64(t.UnixNano())
			}
		}
	}

	return 0
}

func (f *Field) Uint() uint {
	return uint(f.Uint64())
}

func (f *Field) Float() float64 {

	if f != nil && f.value.Interface() != nil {

		vv := reflect.ValueOf(f.value.Interface())

		switch f.valueType.Kind() {
		case reflect.Float32, reflect.Float64:
			return vv.Float()
		}
	}

	return 0
}

func (f *Field) Bool() bool {

	if b, err := strconv.ParseBool(f.String()); err == nil {
		return b
	}

	return false
}

func (f *Field) TimeParse(format string) time.Time {

	if f == nil || f.value.Interface() == nil {
		return time.Now().In(TimeZone)
	}

	vv := reflect.ValueOf(f.value.Interface())

	timeString := ""
	switch f.valueType.Kind() {
	case reflect.Slice:
		if f.valueType.Elem().Kind() == reflect.Uint8 {
			timeString = string(f.value.Interface().([]byte))
		}
	case reflect.String:
		timeString = vv.String()

	case reflect.Struct:
		if t, ok := f.value.Interface().(time.Time); ok {
			return t
		}
	}

	if format == "datetime" && f.datetime_fmt != "" {
		format = f.datetime_fmt
	}

	return TimeParse(timeString, format)
}

func (f *Field) TimeFormat(format, formatTo string) string {
	return f.TimeParse(format).Format(TimeFormat(formatTo))
}
