// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Daily
package habitrpg

import (
	"time"
)

type Daily struct {
	Attribute         string        `json:"attribute"`
	Challenge         struct{}      `json:"challenge"`
	Checklist         []interface{} `json:"checklist"`
	CollapseChecklist bool          `json:"collapseChecklist"`
	Completed         bool          `json:"completed"`
	DateCreated       string        `json:"dateCreated"`
	History           []struct {
		Date  string  `json:"date"`
		Value float64 `json:"value"`
	} `json:"history"`
	ID       string  `json:"id"`
	Notes    string  `json:"notes"`
	Priority float64 `json:"priority"`
	Repeat   struct {
		F  bool `json:"f"`
		M  bool `json:"m"`
		S  bool `json:"s"`
		Su bool `json:"su"`
		T  bool `json:"t"`
		Th bool `json:"th"`
		W  bool `json:"w"`
	} `json:"repeat"`
	Streak float64  `json:"streak"`
	Tags   struct{} `json:"tags"`
	Text   string   `json:"text"`
	Type   string   `json:"type"`
	Value  float64  `json:"value"`
}

// dateCreatedTime is a convenience wrapper that returns the dateCreated time,
// parsed as a time.Time struct
func (d Daily) dateCreatedTime() (time.Time, error) {
	return time.Parse(time.RubyDate, d.DateCreated)
}
