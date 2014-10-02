// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Gear
package habitrpg

type Gear struct {
	costume struct {
		shield string
		head   string
		armor  string
		weapon string
	}
	equipped struct {
		shield string
		head   string
		armor  string
		weapon string
	}
	owned struct {
		armor_warrior_1  bool
		armor_wizard_1   bool
		head_warrior_1   bool
		head_wizard_1    bool
		shield_warrior_1 bool
		shield_warrior_2 bool
		weapon_warrior_1 bool
		weapon_warrior_2 bool
		weapon_warrior_3 bool
		weapon_warrior_4 bool
		weapon_wizard_0  bool
		weapon_wizard_1  bool
		weapon_warrior_0 bool
	}
}
