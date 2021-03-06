// Copyright 2019 Eryx <evorui аt gmаil dοt cοm>, All rights reserved.
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

package sko // import "github.com/lynkdb/iomix/sko"

import (
	"errors"
	"regexp"
	"time"
)

var (
	incrNamespaceReg = regexp.MustCompile("^[a-z]{1}[a-z0-9]{3,9}$")
)

func NewObjectWriter(key []byte, value interface{}) *ObjectWriter {
	r := &ObjectWriter{
		Meta: &ObjectMeta{
			Key: key,
		},
	}
	return r.DataValueSet(value, nil)
}

func (it *ObjectWriter) ModeCreateSet(v bool) *ObjectWriter {
	it.Mode = it.Mode | ObjectWriterModeCreate
	if !v {
		it.Mode = it.Mode - ObjectWriterModeCreate
	}
	return it
}

func (it *ObjectWriter) IncrNamespaceSet(ns string) *ObjectWriter {
	it.IncrNamespace = ns
	return it
}

func (it *ObjectWriter) ModeDeleteSet(v bool) *ObjectWriter {
	it.Mode = it.Mode | ObjectWriterModeDelete
	if !v {
		it.Mode = it.Mode - ObjectWriterModeDelete
	}
	return it
}

func (it *ObjectWriter) ExpireSet(v int64) *ObjectWriter {
	if v > 0 {
		it.Meta.Expired = uint64((time.Now().UnixNano() / 1e6) + v)
	}
	return it
}

func (it *ObjectWriter) DataValueSet(
	value interface{}, codec DataValueCodec) *ObjectWriter {

	if value != nil {

		if codec == nil {
			codec = dataValueCodecStd
		}

		bsValue, err := codec.Encode(value)
		if err == nil {
			it.Data = &ObjectData{
				Check: bytesCrc32Checksum(bsValue),
				Value: bsValue,
			}
		}
	}

	return it
}

func (it *ObjectWriter) PrevDataCheckSet(
	value interface{}) *ObjectWriter {

	if value != nil {

		bsValue, err := dataValueCodecStd.Encode(value)
		if err == nil {
			it.PrevDataCheck = bytesCrc32Checksum(bsValue)
		}
	}

	return it
}

func (it *ObjectWriter) CommitValid() error {

	if it.Meta == nil {
		return errors.New("Meta Not Found")
	}

	if !objectMetaKeyValid(it.Meta.Key) {
		return errors.New("Invalid Meta/Key")
	}

	if AttrAllow(it.Mode, ObjectWriterModeDelete) {
		it.IncrNamespace = ""
		it.Meta.Updated = uint64(time.Now().UnixNano() / 1e6)
		return nil
	}

	if (it.Meta.IncrId > 0 || it.IncrNamespace != "") &&
		!incrNamespaceReg.MatchString(it.IncrNamespace) {
		return errors.New("Invalid IncrNamespace")
	}

	if it.Data == nil {
		return errors.New("Data Not Found")
	}

	if err := it.Data.Valid(); err != nil {
		return err
	}

	it.Meta.Updated = uint64(time.Now().UnixNano() / 1e6)

	if it.Meta.Expired > 0 &&
		it.Meta.Expired <= it.Meta.Updated {
		return errors.New("Invalid Meta/Expired")
	}

	it.Meta.DataAttrs = it.Data.Attrs
	it.Meta.DataCheck = it.Data.Check

	return nil
}

func (it *ObjectWriter) MetaEncode() ([]byte, error) {

	meta, err := protobufEncode(it.Meta)
	if err == nil && len(meta) > 0 {
		return append([]byte{objectRawBytesVersion1, uint8(len(meta))}, meta...), nil
	}

	return nil, errors.New("invalid meta")
}

func (it *ObjectWriter) PutEncode() ([]byte, []byte, error) {

	meta, err := it.MetaEncode()
	if err != nil {
		return nil, nil, err
	}

	data, err := protobufEncode(it.Data)
	if err != nil || len(data) < 1 {
		return nil, nil, errors.New("invalid data")
	}

	return meta, append(meta, data...), nil
}
