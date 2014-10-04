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
	Attribute   string   `json:"attribute"`
	Challenge   struct{} `json:"challenge"`
	DateCreated string   `json:"dateCreated"`
	Down        bool     `json:"down"`
	History     []struct {
		Date  float64 `json:"date"`
		Value float64 `json:"value"`
	} `json:"history"`
	ID       string   `json:"id"`
	Notes    string   `json:"notes"`
	Priority float64  `json:"priority"`
	Tags     struct{} `json:"tags"`
	Text     string   `json:"text"`
	Type     string   `json:"type"`
	Up       bool     `json:"up"`
	Value    float64  `json:"value"`
}

// dateCreatedTime is a convenience wrapper that returns the dateCreated time,
// parsed as a time.Time struct
func (h Habit) dateCreatedTime() (time.Time, error) {
	return time.Parse(time.RubyDate, h.DateCreated)
}
