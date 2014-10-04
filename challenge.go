// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Challenge
package habitrpg

import (
	"time"
)

type Challenge struct {
	V           float64       `json:"__v"`
	ID          string        `json:"_id"`
	Dailys      []interface{} `json:"dailys"`
	Description string        `json:"description"`
	Group       string        `json:"group"`
	Habits      []Habib       `json:"habits"`
	Leader      string        `json:"leader"`
	MemberCount float64       `json:"memberCount"`
	Members     []User        `json:"members"`
	Name        string        `json:"name"`
	Official    bool          `json:"official"`
	Prize       float64       `json:"prize"`
	Rewards     []interface{} `json:"rewards"`
	ShortName   string        `json:"shortName"`
	Timestamp   string        `json:"timestamp"`
	Todos       []Todo        `json:"todos"`
}

// dateCreatedTime is a convenience wrapper that returns the dateCreated time,
// parsed as a time.Time struct
func (c Challenge) timestampTime() (time.Time, error) {
	return time.Parse(time.RubyDate, c.Timestamp)
}
