// Copyright 2015 Eryx <evorui аt gmаil dοt cοm>, All rights reserved.
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

package skv

import (
	"hash/crc32"

	"github.com/golang/protobuf/proto"
	"github.com/lessos/lessgo/types"
)

const (
	ResultOK               uint8 = 0
	ResultNotFound         uint8 = 2
	ResultBadArgument      uint8 = 3
	ResultServerError      uint8 = 4
	ResultNetworkException uint8 = 5
	ResultTimeout          uint8 = 6
)

type Result struct {
	Status uint8
	Data   [][]byte
}

var (
	result_ok               = &Result{Status: ResultOK}
	result_not_found        = &Result{Status: ResultNotFound}
	result_bad_argument     = &Result{Status: ResultBadArgument}
	result_network_err      = &Result{Status: ResultNetworkException}
	result_server_exception = &Result{Status: ResultServerError}
)

func NewResultNotFound() *Result {
	return result_not_found
}

func NewResultBadArgument() *Result {
	return result_bad_argument
}

func NewResultNetworkError() *Result {
	return result_network_err
}

func NewResultError(status uint8, err string) *Result {

	return &Result{
		Status: status,
		Data: [][]byte{
			[]byte(err),
		},
	}
}

func NewResult(status uint8) *Result {

	if status == 0 {
		status = ResultOK
	}

	return &Result{
		Status: status,
	}
}

func (rs *Result) OK() bool {
	return rs.Status == ResultOK
}

func (rs *Result) NotFound() bool {
	return rs.Status == ResultNotFound
}

func (rs *Result) ValueSize() int64 {
	if bs := rs.Bytes(); len(bs) > 1 {
		return int64(len(bs) - 1)
	}
	return 0
}

func (rs *Result) Bytex() types.Bytex {

	if bs := rs.Bytes(); len(bs) > 1 {
		if bs[0] == value_ns_bytes {
			return types.Bytex(bs[1:])
		}
	}

	return types.Bytex{}
}

func (rs *Result) Bytes() []byte {

	if len(rs.Data) > 0 && len(rs.Data[0]) > 0 {

		if rs.Data[0][0] == value_ns_prog {
			meta_len := int(rs.Data[0][1])
			if len(rs.Data[0]) > (meta_len + 2) {
				return rs.Data[0][(meta_len + 2):]
			}
		}

		return rs.Data[0]
	}

	return []byte{}
}

func (rs *Result) Crc32() uint32 {
	if bs := rs.Bytes(); len(bs) > 1 {
		return crc32.ChecksumIEEE(bs[1:])
	}
	return 0
}

func (rs *Result) Int64() int64 {
	return ValueBytes(rs.Bytes()).Int64()
}

func (rs *Result) Uint8() uint8 {
	return ValueBytes(rs.Bytes()).Uint8()
}

func (rs *Result) Uint16() uint16 {
	return ValueBytes(rs.Bytes()).Uint16()
}

func (rs *Result) Uint32() uint32 {
	return ValueBytes(rs.Bytes()).Uint32()
}

func (rs *Result) Uint64() uint64 {
	return ValueBytes(rs.Bytes()).Uint64()
}

// func (rs *Result) Float32() float32 {
// 	return rs.Bytex().Float32()
// }

// func (rs *Result) Float64() float64 {
// 	return rs.Bytex().Float64()
// }

// func (rs *Result) Bool() bool {
// 	return rs.Bytex().Bool()
// }

// func (rs *Result) List() [][]byte {
// 	return rs.Data
// }

func (rs *Result) KvList() []*ResultEntry {

	ls := []*ResultEntry{}

	for i := 0; i < (len(rs.Data) - 1); i += 2 {
		ls = append(ls, &ResultEntry{Key: rs.Data[i], Value: rs.Data[i+1]})
	}

	return ls
}

func (rs *Result) KvLen() int {
	return len(rs.Data) / 2
}

func (rs *Result) KvEach(fn func(entry *ResultEntry) int) int {

	rl := 0

	if len(rs.Data) < 2 {
		return 0
	}

	for i := 0; i < (len(rs.Data) - 1); i += 2 {

		if fn(&ResultEntry{Key: rs.Data[i], Value: rs.Data[i+1]}) < 0 {
			break
		}

		rl++
	}

	return rl
}

func (rs *Result) Decode(obj interface{}) error {
	return ValueDecode(rs.Bytes(), obj)
}

func (rs *Result) Meta() *ValueMeta {
	if len(rs.Data) > 0 && len(rs.Data[0]) > 1 && rs.Data[0][0] == value_ns_prog {
		meta_len := int(rs.Data[0][1])
		if (meta_len + 2) <= len(rs.Data[0]) {
			var meta ValueMeta
			if err := proto.Unmarshal(rs.Data[0][2:(2+meta_len)], &meta); err == nil {
				return &meta
			}
		}
	}
	return nil
}

//
type ResultEntry struct {
	Key   []byte
	Value []byte
}

func (re *ResultEntry) ValueSize() int64 {
	if bs := re.Bytes(); len(bs) > 1 {
		return int64(len(bs) - 1)
	}
	return 0
}

func (re *ResultEntry) Crc32() uint32 {
	if bs := re.Bytes(); len(bs) > 1 {
		return crc32.ChecksumIEEE(bs[1:])
	}
	return 0
}

func (re *ResultEntry) Bytes() []byte {

	if len(re.Value) > 1 && re.Value[0] == value_ns_prog {
		meta_len := int(re.Value[1])
		if len(re.Value) > (meta_len + 2) {
			return re.Value[(meta_len + 2):]
		}
	}

	return re.Value
}

func (re *ResultEntry) Bytex() types.Bytex {

	if bs := re.Bytes(); len(bs) > 1 && bs[0] == value_ns_bytes {
		return types.Bytex(bs[1:])
	}

	return types.Bytex{}
}

func (re *ResultEntry) Decode(obj interface{}) error {
	return ValueDecode(re.Bytes(), obj)
}

func (re *ResultEntry) Int64() int64 {
	return ValueBytes(re.Bytes()).Int64()
}

func (re *ResultEntry) Uint64() uint64 {
	return ValueBytes(re.Bytes()).Uint64()
}

func (re *ResultEntry) Meta() *ValueMeta {

	if len(re.Value) > 1 && re.Value[0] == value_ns_prog {
		meta_len := int(re.Value[1])
		if (meta_len + 2) <= len(re.Value) {
			var meta ValueMeta
			if err := proto.Unmarshal(re.Value[2:(2+meta_len)], &meta); err == nil {
				return &meta
			}
		}
	}

	return nil
}
