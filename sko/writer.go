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
	"time"
)

func NewObjectWriter(key []byte) *ObjectWriter {
	r := &ObjectWriter{
		Meta: &ObjectMeta{
			Key: key,
		},
	}
	return r
}

func (it *ObjectWriter) ExpireSet(v int64) *ObjectWriter {
	if v > 0 {
		it.Meta.Expired = uint64((time.Now().UnixNano() / 1e6) + v)
	}
	return it
}

func (it *ObjectWriter) DelValid() error {

	if it.Meta == nil {
		return errors.New("Meta Not Found")
	}

	if !objectMetaKeyValid(it.Meta.Key) {
		return errors.New("Invalid Meta/Key")
	}

	return nil
}

func (it *ObjectWriter) PutValid() error {

	if it.Meta == nil {
		return errors.New("Meta Not Found")
	}

	if !objectMetaKeyValid(it.Meta.Key) {
		return errors.New("Invalid Meta/Key")
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

func (it *ObjectWriter) DataValueSet(
	value interface{}, codec DataValueCodec) *ObjectWriter {

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

	return it
}

func (it *ObjectWriter) PutEncode() ([]byte, []byte, error) {

	var (
		meta []byte
		data []byte
		err  error
	)

	meta, err = protobufEncode(it.Meta)
	if err == nil && len(meta) > 0 {
		data, err = protobufEncode(it.Data)
	}

	if err != nil {
		return nil, nil, err
	}

	if len(meta) < 1 || len(data) < 1 {
		return nil, nil, errors.New("invalid meta or data")
	}

	metav := append([]byte{objectRawBytesVersion1, uint8(len(meta))}, meta...)

	return metav, append(metav, data...), nil
}
