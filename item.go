// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Item
package habitrpg

type Item struct {
	CurrentMount    string         `json:"currentMount"`
	CurrentPet      string         `json:"currentPet"`
	Eggs            Egg            `json:"eggs"`
	Food            Food           `json:"food"`
	Gear            Gear           `json:"gear"`
	HatchingPotions HatchingPotion `json:"hatchingPotions"`
	LastDrop        struct {
		Count float64 `json:"count"`
		Date  string  `json:"date"`
	} `json:"lastDrop"`
	Mounts  Mount    `json:"mounts"`
	Pets    Pet      `json:"pets"`
	Quests  struct{} `json:"quests"`
	Special struct {
		Snowball          float64       `json:"snowball"`
		SpookDust         float64       `json:"spookDust"`
		ValentineReceived []interface{} `json:"valentineReceived"`
	} `json:"special"`
}
