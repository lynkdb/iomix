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
	"bytes"
)

func NewObjectReader() *ObjectReader {
	return &ObjectReader{}
}

func (it *ObjectReader) KeySet(key []byte) *ObjectReader {
	it.Mode = ObjectReaderModeKey
	for _, v := range it.Keys {
		if bytes.Compare(v, key) == 0 {
			return it
		}
	}
	it.Keys = append(it.Keys, key)
	return it
}

func (it *ObjectReader) KeyRangeSet(keyOffset, keyCutset []byte) *ObjectReader {
	it.Mode = ObjectReaderModeKeyRange
	it.KeyOffset = keyOffset
	it.KeyCutset = keyCutset
	if it.LimitNum == 0 {
		it.LimitNum = 10
	}
	return it
}

func (it *ObjectReader) LimitNumSet(num int64) *ObjectReader {
	it.LimitNum = num
	return it
}

func (it *ObjectReader) RevRangeSet(v bool) *ObjectReader {
	if v {
		it.Mode = AttrAppend(it.Mode, ObjectReaderModeRevRange)
	} else {
		it.Mode = AttrRemove(it.Mode, ObjectReaderModeRevRange)
	}
	return it
}
