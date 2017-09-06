// Copyright 2014 Eryx <evorui аt gmаil dοt cοm>, All rights reserved.
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

package modeler

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type DatabaseEntry struct {
	DbName  string   `json:"dbname"`
	Engine  string   `json:"engine,omitempty"`
	Charset string   `json:"charset,omitempty"`
	Version int      `json:"version,omitempty"`
	Tables  []*Table `json:"tables"`
}

func LoadDatabaseEntryFromFile(file string) (DatabaseEntry, error) {

	var ds DatabaseEntry
	var err error

	if _, err := os.Stat(file); err != nil && os.IsNotExist(err) {
		return ds, err
	}

	fp, err := os.Open(file)
	if err != nil {
		return ds, err
	}
	defer fp.Close()

	cfg, err := ioutil.ReadAll(fp)
	if err != nil {
		return ds, err
	}

	return LoadDatabaseEntryFromString(string(cfg))
}

func LoadDatabaseEntryFromString(js string) (DatabaseEntry, error) {

	var ds DatabaseEntry

	err := json.Unmarshal([]byte(js), &ds)

	return ds, err
}
