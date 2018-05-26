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

package rdb

import (
	"time"
)

const (
	timeFormatDate     = "2006-01-02"
	timeFormatDateTime = "2006-01-02 15:04:05"
	timeFormatAtom     = time.RFC3339
)

var (
	TimeZone = time.UTC
)

func TimeFormat(format string) string {

	switch format {
	case "datetime":
		return timeFormatDateTime
	case "date":
		return timeFormatDate
	case "atom":
		return timeFormatAtom
	}

	return timeFormatDateTime
}

func TimeNow(format string) string {
	return time.Now().In(TimeZone).Format(TimeFormat(format))
}

func TimeNowAdd(format, add string) string {
	taf, err := time.ParseDuration(add)
	if err != nil {
		taf = 0
	}
	return time.Now().Add(taf).In(TimeZone).Format(TimeFormat(format))
}

func TimeZoneFormat(t time.Time, tz, format string) string {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return t.Format(TimeFormat(format))
	}
	return t.In(loc).Format(TimeFormat(format))
}

func TimeParse(timeString, format string) time.Time {

	tp, err := time.ParseInLocation(TimeFormat(format), timeString, TimeZone)
	if err != nil {
		return time.Now().In(TimeZone)
	}

	return tp
}
