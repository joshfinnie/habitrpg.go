// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Task
package habitrpg

import (
	"time"
)

type Task struct {
	id          string
	dateCreated string
	text        string
	notes       string
	tags        interface{}
	value       int
	priority    int
	attribute   string
	challenge   *Challenge
	down        bool
	up          bool
	history     []struct {
		date  int
		value int
	}
	taskType string `json:"type"`
}

// dateCreatedTime is a convenience wrapper that returns the dateCreated time,
// parsed as a time.Time struct
func (t Task) dateCreatedTime() (time.Time, error) {
	return time.Parse(time.RubyDate, t.dateCreated)
}
