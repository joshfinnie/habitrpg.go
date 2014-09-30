// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Habit
package habitrpg

import (
	"time"
)

type Habit struct {
	text        string
	attribute   string
	priority    int
	value       int // really should be a float
	note        string
	dateCreated string
	id          string
	down        bool
	up          bool
	history     []struct {
		value int // should be a float
		date  int
	}
	habitType string `json:"type"`
}

// dateCreatedTime is a convenience wrapper that returns the dateCreated time,
// parsed as a time.Time struct
func (h Habit) dateCreatedTime() (time.Time, error) {
	return time.Parse(time.RubyDate, h.dateCreated)
}
