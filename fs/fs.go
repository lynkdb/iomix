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

package fs // import "code.hooto.com/lynkdb/iomix/fs"

import (
	"os"
	"time"
)

type Connector interface {
	Stat(path string) (FsObjectMeta, error)
	Open(path string) (FsObject, error)
	OpenFile(path string, flag int, perm os.FileMode) (FsObject, error)
	List(path string, limit int) ([]FsObjectMeta, error)
	MkdirAll(path string, perm os.FileMode) error
	Close() error
}

type FsObject interface {
	Stat() (FsObjectMeta, error)
	Seek(offset int64, whence int) (ret int64, err error)
	Truncate(size int64) error
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	WriteAt(b []byte, off int64) (n int, err error)
	Close() error
}

type FsObjectMeta interface {
	Name() string
	Size() int64
	IsDir() bool
	ModTime() time.Time
}
