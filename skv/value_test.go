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
	"math/rand"
	"testing"
)

var (
	value_bench_uint_sets = [][]byte{}
	value_bench_nint_sets = [][]byte{}
)

func init() {

	for i := 0; i < 25; i++ {
		bs, _ := ValueEncode(rand.Uint64(), nil)
		value_bench_uint_sets = append(value_bench_uint_sets, bs)

		nbs, _ := ValueEncode(-rand.Int63(), nil)
		value_bench_nint_sets = append(value_bench_nint_sets, nbs)
	}

	for i := 0; i < 50; i++ {
		bs, _ := ValueEncode(uint64(rand.Uint32()), nil)
		value_bench_uint_sets = append(value_bench_uint_sets, bs)

		nbs, _ := ValueEncode(-int64(rand.Uint32()), nil)
		value_bench_nint_sets = append(value_bench_nint_sets, nbs)
	}

	for i := 0; i < 25; i++ {
		bs, _ := ValueEncode(uint64(rand.Int31n(16777216)), nil)
		value_bench_uint_sets = append(value_bench_uint_sets, bs)

		nbs, _ := ValueEncode(-int64(rand.Int31n(16777216)), nil)
		value_bench_nint_sets = append(value_bench_nint_sets, nbs)
	}
}

func TestValueEncode(t *testing.T) {

	us := []uint{0, 1, 1<<32 - 1}
	for _, u := range us {

		bs, err := ValueEncode(u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if ValueBytes(bs).Uint() != u {
			t.Fatal("Failed on ValueBytes/Decode/Uint")
		}
	}

	u8s := []uint8{0, 1, 1<<8 - 1}
	for _, u := range u8s {

		bs, err := ValueEncode(u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if ValueBytes(bs).Uint8() != u {
			t.Fatal("Failed on ValueBytes/Decode/Uint")
		}
	}

	u16s := []uint16{0, 1, 255, 256, 257, 1<<16 - 1}
	for _, u := range u16s {

		bs, err := ValueEncode(u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if ValueBytes(bs).Uint16() != u {
			t.Fatal("Failed on ValueBytes/Decode/Uint")
		}
	}

	u32s := []uint32{0, 1, 255, 256, 257, 65535, 65536, 1<<32 - 1}
	for _, u := range u32s {

		bs, err := ValueEncode(u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if ValueBytes(bs).Uint32() != u {
			t.Fatal("Failed on ValueBytes/Decode/Uint")
		}
	}

	u64s := []uint64{0, 1, 255, 256, 257, 65535, 65536, 1<<64 - 1}
	for _, u := range u64s {

		bs, err := ValueEncode(u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if ValueBytes(bs).Uint64() != u {
			t.Fatal("Failed on ValueBytes/Decode/Uint")
		}
	}

	for _, u := range u64s {

		bs, _ := ValueEncode(u, nil)

		n := ValueBytes(bs).Uint64()

		if value_bindecode_uint64(bs) != value_rawdecode_uint64(bs) ||
			value_bindecode_uint64(bs) != n {
			t.Fatal("Failed on ValueBytes/Decode/Uint")
		}
	}
}

func TestValueEncodeNegative(t *testing.T) {

	us := []int{-1 << 31, -1, 0, 1, 1<<31 - 1}
	for _, u := range us {

		bs, err := ValueEncode(u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if ValueBytes(bs).Int() != u {
			t.Fatal("Failed on ValueBytes/Decode/Int")
		}
	}

	u8s := []int8{-1 << 7, -1, 0, 1, 1<<7 - 1}
	for _, u := range u8s {

		bs, err := ValueEncode(u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if ValueBytes(bs).Int8() != u {
			t.Fatal("Failed on ValueBytes/Decode/Int")
		}
	}

	u16s := []int16{-1 << 15, -1, 0, 1, 1<<15 - 1}
	for _, u := range u16s {

		bs, err := ValueEncode(u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if ValueBytes(bs).Int16() != u {
			t.Fatal("Failed on ValueBytes/Decode/Int")
		}
	}

	u32s := []int32{-1 << 31, -1, 0, 1, 1<<31 - 1}
	for _, u := range u32s {

		bs, err := ValueEncode(u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if ValueBytes(bs).Int32() != u {
			t.Fatal("Failed on ValueBytes/Decode/Int")
		}
	}

	u64s := []int64{-1 << 63, -1, 0, 1, 1<<63 - 1}
	for _, u := range u64s {

		bs, err := ValueEncode(u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if ValueBytes(bs).Int64() != u {
			t.Fatal("Failed on ValueBytes/Decode/Int")
		}
	}
}

func Benchmark_Value_BinDecode(b *testing.B) {

	if len(value_bench_uint_sets) < 1 {
		b.Fatal("No Samples")
	}

	for i := 0; i < b.N; i++ {
		for _, b := range value_bench_uint_sets {
			value_bindecode_uint64(b)
		}
	}
}

func Benchmark_Value_RawDecode(b *testing.B) {

	if len(value_bench_uint_sets) < 1 {
		b.Fatal("No Samples")
	}

	for i := 0; i < b.N; i++ {
		for _, b := range value_bench_uint_sets {
			value_rawdecode_uint64(b)
		}
	}
}

func value_rawdecode_uint64(v []byte) uint64 {

	if len(v) <= 1 || v[0] != value_ns_uint {
		return 0
	}

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

	return 0
}

func value_bindecode_uint64(v []byte) uint64 {

	if len(v) <= 1 || v[0] != value_ns_uint {
		return 0
	}

	ubs := make([]byte, 8)
	off := 9 - len(v)
	for i, iv := range v[1:] {
		ubs[off+i] = iv
	}

	return binary.BigEndian.Uint64(ubs)
}

func Benchmark_ValueBytes_Decode(b *testing.B) {

	if len(value_bench_uint_sets) < 1 {
		b.Fatal("No Samples")
	}

	for i := 0; i < b.N; i++ {
		for _, b := range value_bench_uint_sets {
			ValueBytes(b).Uint64()
		}
	}
}

func Benchmark_ValueBytes_Decode_Negative(b *testing.B) {

	if len(value_bench_nint_sets) < 1 {
		b.Fatal("No Samples")
	}

	for i := 0; i < b.N; i++ {
		for _, b := range value_bench_nint_sets {
			ValueBytes(b).Int64()
		}
	}
}
