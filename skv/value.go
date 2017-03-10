// Copyright 2015 lynkdb Authors, All rights reserved.
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
	"encoding/binary"
	"encoding/json"
	"errors"

	"github.com/golang/protobuf/proto"
)

const (
	value_ns_bytes    uint8 = 0
	value_ns_uint     uint8 = 1
	value_ns_nint     uint8 = 2 // negative int
	value_ns_json     uint8 = 20
	value_ns_protobuf uint8 = 21
)

type ValueEncoder interface {
	Encode(value interface{}) error
}

type ValueDecoder interface {
	Decode(value []byte, object interface{}) error
}

type ValueBytes []byte

func (v ValueBytes) Int() int {
	return int(v.Int64())
}

func (v ValueBytes) Int8() int8 {
	return int8(v.Int64())
}

func (v ValueBytes) Int16() int16 {
	return int16(v.Int64())
}

func (v ValueBytes) Int32() int32 {
	return int32(v.Int64())
}

func (v ValueBytes) Int64() int64 {

	if len(v) > 0 && (v[0] == value_ns_nint || v[0] == value_ns_uint) {

		var i int64

		switch len(v) {
		case 2:
			i = int64(v[1])

		case 3:
			i = int64(v[1])<<8 | int64(v[2])

		case 4:
			i = int64(v[1])<<16 | int64(v[2])<<8 | int64(v[3])

		case 5:
			i = int64(v[1])<<24 | int64(v[2])<<16 | int64(v[3])<<8 |
				int64(v[4])

		case 6:
			i = int64(v[1])<<32 | int64(v[2])<<24 | int64(v[3])<<16 |
				int64(v[4])<<8 | int64(v[5])

		case 7:
			i = int64(v[1])<<40 | int64(v[2])<<32 | int64(v[3])<<24 |
				int64(v[4])<<16 | int64(v[5])<<8 | int64(v[6])

		case 8:
			i = int64(v[1])<<48 | int64(v[2])<<40 | int64(v[3])<<32 |
				int64(v[4])<<24 | int64(v[5])<<16 | int64(v[6])<<8 |
				int64(v[7])

		case 9:
			i = int64(v[1])<<56 | int64(v[2])<<48 | int64(v[3])<<40 |
				int64(v[4])<<32 | int64(v[5])<<24 | int64(v[6])<<16 |
				int64(v[7])<<8 | int64(v[8])
		}

		if v[0] == value_ns_nint {
			return -i
		}

		return i
	}

	return 0
}

func (v ValueBytes) Uint() uint {
	return uint(v.Uint64())
}

func (v ValueBytes) Uint8() uint8 {
	return uint8(v.Uint64())
}

func (v ValueBytes) Uint16() uint16 {
	return uint16(v.Uint64())
}

func (v ValueBytes) Uint32() uint32 {
	return uint32(v.Uint64())
}

func (v ValueBytes) Uint64() uint64 {

	if len(v) > 1 && v[0] == value_ns_uint {

		switch len(v) {
		case 2:
			return uint64(v[1])

		case 3:
			return uint64(v[1])<<8 | uint64(v[2])

		case 4:
			return uint64(v[1])<<16 | uint64(v[2])<<8 | uint64(v[3])

		case 5:
			return uint64(v[1])<<24 | uint64(v[2])<<16 | uint64(v[3])<<8 |
				uint64(v[4])

		case 6:
			return uint64(v[1])<<32 | uint64(v[2])<<24 | uint64(v[3])<<16 |
				uint64(v[4])<<8 | uint64(v[5])

		case 7:
			return uint64(v[1])<<40 | uint64(v[2])<<32 | uint64(v[3])<<24 |
				uint64(v[4])<<16 | uint64(v[5])<<8 | uint64(v[6])

		case 8:
			return uint64(v[1])<<48 | uint64(v[2])<<40 | uint64(v[3])<<32 |
				uint64(v[4])<<24 | uint64(v[5])<<16 | uint64(v[6])<<8 |
				uint64(v[7])

		case 9:
			return uint64(v[1])<<56 | uint64(v[2])<<48 | uint64(v[3])<<40 |
				uint64(v[4])<<32 | uint64(v[5])<<24 | uint64(v[6])<<16 |
				uint64(v[7])<<8 | uint64(v[8])
		}
	}

	return 0
}

func ValueDecode(value []byte, object interface{}) error {

	if len(value) < 2 {
		return errors.New("Invalid Data len<2")
	}

	switch value[0] {

	case value_ns_protobuf:
		if obj, ok := object.(proto.Message); ok {
			if err := proto.Unmarshal(value[1:], obj); err != nil {
				return errors.New("Invalid ProtoBuf " + err.Error())
			}
			return nil
		}

	case value_ns_json:
		return json.Unmarshal(value[1:], object)
	}

	return errors.New("Invalid Data")
}

func ValueEncode(value interface{}, encode ValueEncoder) ([]byte, error) {

	var enc_value []byte

	switch value.(type) {

	case []byte:
		enc_value = append([]byte{value_ns_bytes}, value.([]byte)...)

	case string:
		enc_value = append([]byte{value_ns_bytes}, []byte(value.(string))...)

		//
	case uint:
		enc_value = value_encode_uint(uint64(value.(uint)))

	case uint8:
		enc_value = value_encode_uint(uint64(value.(uint8)))

	case uint16:
		enc_value = value_encode_uint(uint64(value.(uint16)))

	case uint32:
		enc_value = value_encode_uint(uint64(value.(uint32)))

	case uint64:
		enc_value = value_encode_uint(value.(uint64))

		//
	case int:
		enc_value = value_encode_int(int64(value.(int)))

	case int8:
		enc_value = value_encode_int(int64(value.(int8)))

	case int16:
		enc_value = value_encode_int(int64(value.(int16)))

	case int32:
		enc_value = value_encode_int(int64(value.(int32)))

	case int64:
		enc_value = value_encode_int(value.(int64))

		//
	case proto.Message:
		if bs, err := proto.Marshal(value.(proto.Message)); err != nil {
			return nil, errors.New("BadArgument ProtoBuf " + err.Error())
		} else {
			enc_value = append([]byte{value_ns_protobuf}, bs...)
		}

		//
	case map[string]interface{}, struct{}, interface{}:
		if bs_json, err := json.Marshal(value); err != nil {
			return nil, errors.New("BadArgument JSON")
		} else {
			enc_value = append([]byte{value_ns_json}, bs_json...)
		}

	default:
		return nil, errors.New("BadArgument")
	}

	return enc_value, nil
}

func value_encode_uint(num uint64) []byte {

	enc_value := []byte{value_ns_uint}

	if num > 0 {

		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, num)

		for i := 0; i <= 7; i++ {
			if buf[i] > 0 {
				enc_value = append(enc_value, buf[i:]...)
				break
			}
		}
	}

	return enc_value
}

func value_encode_int(num int64) []byte {

	enc_value := []byte{value_ns_uint}

	if num < 0 {
		enc_value[0] = value_ns_nint
		num = (-num)
	}

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(num))

	for i := 0; i <= 7; i++ {
		if buf[i] > 0 {
			enc_value = append(enc_value, buf[i:]...)
			break
		}
	}

	return enc_value
}
