// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Challenges
package habitrpg

type Challenge struct {
	_id         string
	description string
	group       *Group
	leader      *User
	name        string
	shortName   string
	official    bool
	habits      []Task
	dailys      []Task
	todos       []Task
	rewards     []Task
	members     []User
	memberCount int
	prize       int
	_isMember   bool
}
