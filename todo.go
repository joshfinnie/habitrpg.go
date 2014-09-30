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
	text              string
	attribute         string
	priority          int
	value             int
	notes             string
	dateCreated       string
	id                string
	checklist         []string
	collapseChecklist bool
	completed         bool
	todoType          string `json:"type"`
}

// dateCreatedTime is a convenience wrapper that returns the dateCreated time,
// parsed as a time.Time struct
func (t Todo) dateCreatedTime() (time.Time, error) {
	return time.Parse(time.RubyDate, t.dateCreated)
}
