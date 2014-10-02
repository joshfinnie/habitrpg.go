// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Stat
package habitrpg

type Stat struct {
	training struct {
		con int
		str int
		per int
		int int
	}
	buffs struct {
		snowball bool
		streaks  bool
		stealth  int
		con      int
		per      int
		int      int
		str      int
	}
	per         int
	int         int
	con         int
	str         int
	points      int
	class       string
	lvl         int
	gp          int //float
	exp         int //float
	mp          int
	hp          int
	toNextLevel int
	maxHealth   int
	maxMP       int
}
