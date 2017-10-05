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

package skv // import "github.com/lynkdb/iomix/skv"

import (
	"encoding/binary"
	"errors"
	"hash/crc32"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	prog_ttl_zero int64 = 1500000000000000000
)

const (
	ProgKeyEntryUnknown uint32 = 0
	ProgKeyEntryBytes   uint32 = 1
	ProgKeyEntryUint    uint32 = 4
	ProgKeyEntryIncr    uint32 = 16
)

type ProgKeyEntry struct {
	Type uint32
	Data []byte
}

type ProgKey struct {
	enc       []byte
	fold_meta []byte
	Items     []*ProgKeyEntry
}

func NewProgKey(values ...interface{}) ProgKey {
	k := ProgKey{}
	for _, value := range values {
		k.Append(value)
	}
	return k
}

func newProgKeyEntry(value interface{}) (*ProgKeyEntry, error) {

	set := &ProgKeyEntry{}

	switch value.(type) {

	case []byte:
		set.Type = ProgKeyEntryBytes
		if bs := value.([]byte); len(bs) > 0 {
			set.Data = bs
		}

	case string:
		set.Type = ProgKeyEntryBytes
		if bs := []byte(value.(string)); len(bs) > 0 {
			set.Data = bs
		}

	case uint8:
		set.Type, set.Data = ProgKeyEntryUint, []byte{value.(uint8)}

	case uint16:
		set.Type, set.Data = ProgKeyEntryUint, make([]byte, 2)
		binary.BigEndian.PutUint16(set.Data, value.(uint16))

	case uint32:
		set.Type, set.Data = ProgKeyEntryUint, make([]byte, 4)
		binary.BigEndian.PutUint32(set.Data, value.(uint32))

	case uint64:
		set.Type, set.Data = ProgKeyEntryUint, make([]byte, 8)
		binary.BigEndian.PutUint64(set.Data, value.(uint64))

	default:
		return nil, errors.New("Invalid Data Type")
	}

	return set, nil
}

func (k *ProgKey) Append(value interface{}) error {
	if len(k.Items) > 20 {
		return errors.New("too many Items")
	}

	set, err := newProgKeyEntry(value)
	if err != nil {
		return err
	}
	k.Items = append(k.Items, set)

	if len(k.enc) > 0 {
		k.enc = []byte{}
	}
	if len(k.fold_meta) > 0 {
		k.fold_meta = []byte{}
	}

	return nil
}

func (k *ProgKey) AppendTypeValue(t uint32, value interface{}) error {
	if len(k.Items) > 20 {
		return errors.New("too many Items")
	}

	switch t {
	case ProgKeyEntryIncr:
		k.Items = append(k.Items, &ProgKeyEntry{Type: t})

	case ProgKeyEntryBytes, ProgKeyEntryUint:
		return k.Append(value)

	default:
		return errors.New("Invalid Data Type")
	}

	return nil
}

func (k *ProgKey) Set(idx int, value interface{}) error {

	if idx+1 > len(k.Items) {
		return errors.New("Invalid index")
	}

	set, err := newProgKeyEntry(value)
	if err != nil {
		return err
	}

	k.Items[idx] = set

	return nil
}

func (k *ProgKey) LastEntry() (int, *ProgKeyEntry) {
	if i := (len(k.Items) - 1); i >= 0 {
		return i, k.Items[i]
	}
	return -1, nil
}

func (k *ProgKey) Value(i int) []byte {
	if i > 0 && i <= len(k.Items) {
		return k.Items[i-1].Data
	}
	return []byte{}
}

func (k *ProgKey) Size() int {
	return len(k.Items)
}

func (k *ProgKey) Valid() bool {
	return len(k.Items) > 0
}

func (k *ProgKey) Encode(ns uint8) []byte {
	if len(k.enc) > 0 {
		return k.enc
	}
	if len(k.Items) == 0 {
		return []byte{}
	}

	k.enc = []byte{ns, uint8(len(k.Items))}
	for i, v := range k.Items {
		if (i + 1) < len(k.Items) {
			if len(v.Data) > 0 {
				k.enc = append(k.enc, uint8(len(v.Data)))
				k.enc = append(k.enc, v.Data...)
			} else {
				k.enc = append(k.enc, uint8(0))
			}
		} else if len(v.Data) > 0 {
			k.enc = append(k.enc, v.Data...)
		}
	}

	return k.enc
}

func (k *ProgKey) EncodeMeta(ns uint8) []byte {
	if len(k.Items) == 0 {
		return []byte{}
	}
	if len(k.fold_meta) > 0 {
		return k.fold_meta
	}

	k.fold_meta = []byte{ns, uint8(len(k.Items))}
	for i := 0; i < (len(k.Items) - 1); i++ {
		k.fold_meta = append(k.fold_meta, uint8(len(k.Items[i].Data)))
		k.fold_meta = append(k.fold_meta, k.Items[i].Data...)
	}

	return k.fold_meta
}

func (k *ProgKey) EncodeIndex(ns uint8, idx int) []byte {
	if len(k.Items) == 0 {
		return []byte{}
	}
	if idx < 0 || (idx+1) > len(k.Items) {
		return []byte{}
	}

	enc := []byte{ns, uint8(len(k.Items))}
	for i := 0; i <= idx; i++ {
		enc = append(enc, uint8(len(k.Items[i].Data)))
		enc = append(enc, k.Items[i].Data...)
	}

	return enc
}

func ProgKeyDecode(bs []byte) *ProgKey {
	if len(bs) > 2 {
		var (
			k   = &ProgKey{}
			off = 2
		)
		for i := 0; i < int(bs[1])-1; i++ {
			nlen := int(bs[off])
			if (off + nlen + 1) <= len(bs) {
				k.Items = append(k.Items, &ProgKeyEntry{
					Data: bs[(off + 1):(off + nlen + 1)],
				})
				off += (nlen + 1)
			} else {
				return nil
			}
		}
		if off < len(bs) {
			k.Items = append(k.Items, &ProgKeyEntry{Data: bs[off:]})
		}
		return k
	}
	return nil
}

type ProgValue struct {
	enc   []byte
	meta  *ValueMeta
	value []byte
}

//
func NewProgValue(value interface{}) ProgValue {
	obj := ProgValue{}
	obj.Set(value)
	return obj
}

func (o *ProgValue) Valid() bool {
	if o.meta == nil && len(o.value) < 1 {
		return false
	}
	return true
}

func (o *ProgValue) Encode() []byte {
	if len(o.enc) > 1 {
		return o.enc
	}
	o.enc = []byte{value_ns_prog, 0}
	if o.meta != nil {

		if len(o.value) > 1 {
			if o.meta.Sum > 0 {
				o.meta.Sum = crc32.ChecksumIEEE(o.value[1:])
			}
			if o.meta.Size > 0 {
				o.meta.Size = uint64(len(o.value) - 1)
			}
		}

		if bs, err := proto.Marshal(o.meta); err == nil {
			if len(bs) > 0 && len(bs) < 200 {
				o.enc[1] = uint8(len(bs))
				o.enc = append(o.enc, bs...)
			}
		}
	}
	if len(o.value) > 1 {
		o.enc = append(o.enc, o.value...)
	}
	return o.enc
}

func (o *ProgValue) ValueSize() int64 {
	return int64(len(o.value) - 1)
}

func (o *ProgValue) Set(value interface{}) error {
	var err error
	if o.value, err = ValueEncode(value, nil); err == nil {
		if len(o.enc) > 0 {
			o.enc = []byte{}
		}
	}
	return err
}

func (o *ProgValue) ValueBytes() ValueBytes {
	return ValueBytes(o.value)
}

func (o *ProgValue) Meta() *ValueMeta {
	if o.meta == nil {
		o.meta = &ValueMeta{}
	}
	return o.meta
}

type ProgKeyValue struct {
	Key ProgKey
	Val ProgValue
}

const (
	ProgOpMetaSum   uint64 = 1 << 1
	ProgOpMetaSize  uint64 = 1 << 2
	ProgOpCreate    uint64 = 1 << 13
	ProgOpForce     uint64 = 1 << 14
	ProgOpFoldMeta  uint64 = 1 << 15
	ProgOpLogEnable uint64 = 1 << 16
)

type ProgWriteOptions struct {
	Expired     time.Time
	PrevSum     uint32
	PrevVersion uint64
	Version     uint64
	Actions     uint64
}

func (o *ProgWriteOptions) OpSet(v uint64) *ProgWriteOptions {
	o.Actions = (o.Actions | v)
	return o
}

func (o *ProgWriteOptions) OpAllow(v uint64) bool {
	return (v & o.Actions) == v
}

func (o *ProgWriteOptions) ExpiredUnixNano() uint64 {
	if ns := o.Expired.UTC().UnixNano(); ns > prog_ttl_zero {
		return uint64(ns)
	}
	return 0
}

func (m *ValueMeta) Encode() []byte {
	if bs, err := proto.Marshal(m); err == nil {
		return append([]byte{value_ns_prog, uint8(len(bs))}, bs...)
	}
	return []byte{}
}

func (m *ValueMeta) Timeout() bool {
	if m.Expired > 0 && m.Expired <= uint64(time.Now().UTC().UnixNano()) {
		return true
	}
	return false
}

//
// Programmable Key/Value
type ProgConnector interface {
	ProgPut(k ProgKey, v ProgValue, opts *ProgWriteOptions) *Result
	ProgGet(k ProgKey) *Result
	ProgDel(k ProgKey, opts *ProgWriteOptions) *Result
	ProgScan(offset, cutset ProgKey, limit int) *Result
	ProgRevScan(offset, cutset ProgKey, limit int) *Result
}
