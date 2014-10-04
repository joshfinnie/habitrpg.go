// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Food
package habitrpg

type Food struct {
	CandyBase            float64 `json:"Candy_Base"`
	CandyCottonCandyBlue float64 `json:"Candy_CottonCandyBlue"`
	CandyCottonCandyPink float64 `json:"Candy_CottonCandyPink"`
	CandyDesert          float64 `json:"Candy_Desert"`
	CandyGolden          float64 `json:"Candy_Golden"`
	CandyRed             float64 `json:"Candy_Red"`
	CandyShade           float64 `json:"Candy_Shade"`
	CandySkeleton        float64 `json:"Candy_Skeleton"`
	CandyWhite           float64 `json:"Candy_White"`
	CandyZombie          float64 `json:"Candy_Zombie"`
	Chocolate            float64 `json:"Chocolate"`
	CottonCandyBlue      float64 `json:"CottonCandyBlue"`
	CottonCandyPink      float64 `json:"CottonCandyPink"`
	Fish                 float64 `json:"Fish"`
	Meat                 float64 `json:"Meat"`
	RottenMeat           float64 `json:"RottenMeat"`
	Strawberry           float64 `json:"Strawberry"`
	// There should be more?
}
