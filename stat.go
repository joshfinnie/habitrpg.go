// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Stat
package habitrpg

type Stat struct {
	Buffs struct {
		Con       float64 `json:"con"`
		Int       float64 `json:"int"`
		Per       float64 `json:"per"`
		Snowball  bool    `json:"snowball"`
		SpookDust bool    `json:"spookDust"`
		Stealth   float64 `json:"stealth"`
		Str       float64 `json:"str"`
		Streaks   bool    `json:"streaks"`
	} `json:"buffs"`
	Class       string  `json:"class"`
	Con         float64 `json:"con"`
	Exp         float64 `json:"exp"`
	Gp          float64 `json:"gp"`
	Hp          float64 `json:"hp"`
	Int         float64 `json:"int"`
	Lvl         float64 `json:"lvl"`
	MaxHealth   float64 `json:"maxHealth"`
	MaxMP       float64 `json:"maxMP"`
	Mp          float64 `json:"mp"`
	Per         float64 `json:"per"`
	Points      float64 `json:"points"`
	Str         float64 `json:"str"`
	ToNextLevel float64 `json:"toNextLevel"`
	Training    struct {
		Con float64 `json:"con"`
		Int float64 `json:"int"`
		Per float64 `json:"per"`
		Str float64 `json:"str"`
	} `json:"training"`
}
