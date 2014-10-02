// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Item
package habitrpg

type Item struct {
	currentPet string
	lastDrop   struct {
		count int
		date  string
	}
	quests          interface{}
	mounts          interface{}
	food            Food
	hatchingPotions HatchingPotion
	eggs            Egg
	pets            Pet
	special         struct {
		valentineReceived []string
		snowball          int
	}
	gear Gear
}
