// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Gear
package habitrpg

type Gear struct {
	Costume struct {
		Armor  string `json:"armor"`
		Head   string `json:"head"`
		Shield string `json:"shield"`
		Weapon string `json:"weapon"`
	} `json:"costume"`
	Equipped struct {
		Armor  string `json:"armor"`
		Head   string `json:"head"`
		Shield string `json:"shield"`
		Weapon string `json:"weapon"`
	} `json:"equipped"`
	Owned struct {
		ArmorWarrior1  bool `json:"armor_warrior_1"`
		ArmorWizard1   bool `json:"armor_wizard_1"`
		HeadWarrior1   bool `json:"head_warrior_1"`
		HeadWizard1    bool `json:"head_wizard_1"`
		HeadWizard2    bool `json:"head_wizard_2"`
		HeadWizard3    bool `json:"head_wizard_3"`
		ShieldWarrior1 bool `json:"shield_warrior_1"`
		ShieldWarrior2 bool `json:"shield_warrior_2"`
		WeaponWarrior0 bool `json:"weapon_warrior_0"`
		WeaponWarrior1 bool `json:"weapon_warrior_1"`
		WeaponWarrior2 bool `json:"weapon_warrior_2"`
		WeaponWarrior3 bool `json:"weapon_warrior_3"`
		WeaponWarrior4 bool `json:"weapon_warrior_4"`
		WeaponWizard0  bool `json:"weapon_wizard_0"`
		WeaponWizard1  bool `json:"weapon_wizard_1"`
		WeaponWizard2  bool `json:"weapon_wizard_2"`
		// There should be more...
	} `json:"owned"`
}
