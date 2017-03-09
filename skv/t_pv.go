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
	"path/filepath"
	"strings"
)

//
type PvPath struct {
	Fold string
	Name string
}

func (p *PvPath) EntryIndex() []byte {
	return append([]byte{ns_pv, uint8(len(p.Fold))}, append([]byte(p.Fold), []byte(p.Name)...)...)
}

func pvPathClean(path string) string {
	return strings.Trim(strings.Trim(filepath.Clean(path), "/"), ".")
}

func PvPathParse(path string) *PvPath {

	p := &PvPath{}

	is_fold := false
	if len(path) > 0 && path[len(path)-1] == '/' {
		is_fold = true
	}

	path = pvPathClean(path)

	if is_fold {
		p.Fold, p.Name = path, ""
	} else {
		if i := strings.LastIndex(path, "/"); i > 0 {
			p.Fold, p.Name = path[:i], path[i+1:]
		} else {
			p.Fold, p.Name = "", path
		}
	}

	return p
}

func PvPathFoldIndex(fold string) []byte {
	fold = pvPathClean(fold)
	return append([]byte{ns_pv, uint8(len(fold))}, []byte(fold)...)
}
