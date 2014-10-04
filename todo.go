// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Todo
package habitrpg

import (
	"time"
)

type Todo struct {
	Attribute         string        `json:"attribute"`
	Challenge         Challenge     `json:"challenge"`
	Checklist         []interface{} `json:"checklist"`
	CollapseChecklist bool          `json:"collapseChecklist"`
	Completed         bool          `json:"completed"`
	DateCreated       string        `json:"dateCreated"`
	ID                string        `json:"id"`
	Notes             string        `json:"notes"`
	Priority          float64       `json:"priority"`
	Tags              struct{}      `json:"tags"`
	Text              string        `json:"text"`
	Type              string        `json:"type"`
	Value             float64       `json:"value"`
}

// dateCreatedTime is a convenience wrapper that returns the dateCreated time,
// parsed as a time.Time struct
func (t Todo) dateCreatedTime() (time.Time, error) {
	return time.Parse(time.RubyDate, t.DateCreated)
}
