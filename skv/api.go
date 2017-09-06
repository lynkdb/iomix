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

package skv // import "github.com/lynkdb/iomix/skv"

//go:generate protoc --go_out=plugins=grpc:. pbtypes.proto

import (
	"time"
)

const (
	ScanLimitMax = 10000
)

type Connector interface {
	RawInterface
	KvInterface
	PvInterface
	PoInterface
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
	Ttl       int64     // time to live in milliseconds
	ExpireAt  time.Time // UTC time
	LogEnable bool
	Encoder   ValueEncoder
}

// Key-Value APIs
type KvInterface interface {
	KvNew(key []byte, value interface{}, opts *KvWriteOptions) *Result
	KvDel(key ...[]byte) *Result
	KvPut(key []byte, value interface{}, opts *KvWriteOptions) *Result
	KvGet(key []byte) *Result
	KvScan(offset, cutset []byte, limit int) *Result
	KvRevScan(offset, cutset []byte, limit int) *Result
	KvIncrby(key []byte, increment int64) *Result
}

// Path-Value types
type PathWriteOptions struct {
	Ttl         int64     // time to live in milliseconds
	ExpireAt    time.Time // absolute time to live at
	PrevVersion uint64
	PrevValue   interface{}
	Version     uint64
	LogEnable   bool
	Force       bool
	Encoder     interface{}
}

const (
	PathEventCreated uint8 = 1
	PathEventUpdated uint8 = 2
	PathEventDeleted uint8 = 3
)

type PathEventInterface interface {
	Path() string
	Action() uint8
	Version() uint64
}

type PathEventHandler func(ev PathEventInterface)

// Path-Value APIs
type PvInterface interface {
	//
	PvNew(path string, value interface{}, opts *PathWriteOptions) *Result
	PvDel(path string, opts *PathWriteOptions) *Result
	PvPut(path string, value interface{}, opts *PathWriteOptions) *Result
	PvGet(path string) *Result
	PvScan(fold, offset, cutset string, limit int) *Result
	PvRevScan(fold, offset, cutset string, limit int) *Result
	// PvIncrby(path string, increment int) *Result

	//
	// PvMetaGet(path string) *Result
	// PvMetaScan(fold, offset, cutset string, limit int) *Result
	// PvMetaVersionIncr(path string, group_number uint32, step int64) *Result

	//
	// PvLogScan(offset, cutset uint64, limit int) *Result

	//
	// PathEventRegister(handler PathEventHandler)

	// Status
}

// Path-Object
// path : {path}/{object key}
type PoInterface interface {
	//
	// PoSchemaSync(path string, schema PoSchema) *Result

	//
	PoNew(path string, key, value interface{}, opts *PathWriteOptions) *Result
	PoDel(path string, key interface{}, opts *PathWriteOptions) *Result
	PoPut(path string, key, value interface{}, opts *PathWriteOptions) *Result
	PoGet(path string, key interface{}) *Result
	PoScan(path string, offset, cutset interface{}, limit int) *Result
	PoRevScan(fold string, offset, cutset interface{}, limit int) *Result
	// PoQuery(path string, qry *PoQuerySet) *Result
}
