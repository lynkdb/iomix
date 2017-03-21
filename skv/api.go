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

package skv // import "code.hooto.com/lynkdb/iomix/skv"

import (
	"time"
)

const (
	ns_kv uint8 = 10
	ns_pv uint8 = 11

	ScanLimitMax = 10000
)

type Connector interface {
	RawInterface
	KvInterface
	PvInterface
	Close() error
}

//
type RawInterface interface {
	RawNew(key, value []byte, ttl int64) *Result
	RawDel(key ...[]byte) *Result
	RawPut(key, value []byte, ttl int64) *Result
	RawGet(key []byte) *Result
	RawScan(offset, cutset []byte, limit int) *Result
	RawRevScan(offset, cutset []byte, limit int) *Result
	// RawIncrby(key []byte, increment int) *Result
}

// Key-Value types
type KvWriteOptions struct {
	TimeToLive int64     // in milliseconds
	Expired    time.Time // UTC time
	LogEnable  bool
	Encoder    ValueEncoder
}

// Key-Value APIs
type KvInterface interface {
	KvNew(key []byte, value interface{}, opts *KvWriteOptions) *Result
	KvDel(key ...[]byte) *Result
	KvPut(key []byte, value interface{}, opts *KvWriteOptions) *Result
	KvGet(key []byte) *Result
	KvScan(offset, cutset []byte, limit int) *Result
	// KvRevScan(offset, cutset string, limit int) *Result
	KvIncrby(key []byte, increment int64) *Result
	// KvTtl(key string) *Result
}

// Key-Value types
type PvWriteOptions struct {
	Expire    int       // time to live in seconds
	ExpireAt  time.Time // absolute time to live at
	Version   uint64
	LogEnable bool
	Encoder   interface{}
}

// Path-Value APIs
type PvInterface interface {
	//
	PvNew(path string, value interface{}, opts *PvWriteOptions) *Result
	PvDel(path string) *Result
	PvPut(path string, value interface{}, opts *PvWriteOptions) *Result
	PvGet(path string) *Result
	PvScan(fold, offset, cutset string, limit int) *Result
	// PvRevScan(fold, offset, cutset string, limit int) *Result
	// PvIncrby(path string, increment int) *Result

	//
	// PvMetaGet(path string) *Result
	// PvMetaScan(fold, offset, cutset string, limit int) *Result
	// PvMetaVersionIncr(path string, group_number uint32, step int64) *Result

	//
	// PvLogScan(offset, cutset uint64, limit int) *Result

	//
	// PvEventRegister(handler PvEventHandler)

	// Status
}

const (
	PvEventCreated uint8 = 1
	PvEventUpdated uint8 = 2
	PvEventDeleted uint8 = 3
)

type PvEventInterface interface {
	Path() string
	Action() uint8
	Version() uint64
}

type PvEventHandler func(ev PvEventInterface)

// Indexed Path-Value
// path : {dir}/{primary key}
type PviInterface interface {
	//
	// PviSchemaSync(path_dir string, schema PviSchema) *Result

	//
	PviNew(path_dir string, key []byte, value interface{}, opts *PvWriteOptions) *Result
	PviDel(path_dir string, key []byte) *Result
	PviPut(path_dir string, key []byte, value interface{}, opts *PvWriteOptions) *Result
	PviGet(path_dir string, key []byte) *Result
	// PviQuery(path_dir string, qry *PviQuerySet) *Result
}
