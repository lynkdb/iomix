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
	"io"
	"path/filepath"
	"strings"
)

const (
	ObjStorEntryAttrVersion1   uint64 = 1 << 1
	ObjStorEntryAttrIsDir      uint64 = 1 << 2
	ObjStorEntryAttrBlockSize4 uint64 = 1 << 4
	ObjStorEntryAttrCommiting  uint64 = 1 << 22
	ObjStorBlockSize4          uint64 = 4 * 1024 * 1024
)

type ObjStorConnector interface {
	OsMpInit(sets ObjStorEntryMpInit) Result
	OsMpPut(sets ObjStorEntryBlock) Result
	OsMpGet(sets ObjStorEntryBlock) Result
	OsGet(key string) Result
	OsScan(offset, cutset string, limit int) Result
	OsRevScan(offset, cutset string, limit int) Result
	OsFilePut(src_path string, dst_path string) Result
	OsFileOpen(path string) (io.ReadSeeker, error)
}

func (it *ObjStorEntryMeta) AttrAllow(v uint64) bool {
	return ((v & it.Attrs) == v)
}

func NewObjStorEntryMpInit(path string, size uint64) ObjStorEntryMpInit {
	return ObjStorEntryMpInit{
		Path:  ObjStorPathEncode(path),
		Size:  size,
		Attrs: ObjStorEntryAttrVersion1 | ObjStorEntryAttrBlockSize4,
	}
}

func (it *ObjStorEntryMpInit) Valid() bool {
	if len(it.Path) < 1 || it.Size < 1 {
		return false
	}
	return true
}

func ObjStorPathEncode(path string) string {
	s := strings.Trim(filepath.Clean(path), ".")
	if len(s) == 0 {
		return "/"
	}

	if len(s) > 1 && path[len(path)-1] == '/' {
		s += "/"
	}

	if s[0] != '/' {
		return ("/" + s)
	}

	return s
}

func NewObjStorEntryBlock(path string, size uint64, block_num uint32, data []byte, commit_key string) ObjStorEntryBlock {
	return ObjStorEntryBlock{
		Path:      ObjStorPathEncode(path),
		Size:      size,
		Attrs:     ObjStorEntryAttrVersion1 | ObjStorEntryAttrBlockSize4,
		Num:       block_num,
		Data:      data,
		CommitKey: commit_key,
	}
}

func (it *ObjStorEntryBlock) Valid() bool {
	if len(it.Path) < 1 || it.Size < 1 || len(it.Data) < 1 {
		return false
	}
	num_max := uint32(it.Size / ObjStorBlockSize4)
	if (it.Size % ObjStorBlockSize4) == 0 {
		num_max -= 1
	}
	if it.Num > num_max {
		return false
	}
	if it.Num < num_max {
		if uint64(len(it.Data)) != ObjStorBlockSize4 {
			return false
		}
	} else {
		if uint64(len(it.Data)) != (it.Size % ObjStorBlockSize4) {
			return false
		}
	}
	return true
}
