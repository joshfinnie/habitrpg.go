// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's User
package habitrpg

type User struct {
	V            float64 `json:"__v"`
	ID           string  `json:"_id"`
	V            float64 `json:"_v"`
	Achievements struct {
		BeastMaster bool        `json:"beastMaster"`
		Challenges  []Challenge `json:"challenges"`
		Perfect     float64     `json:"perfect"`
		Streak      float64     `json:"streak"`
	} `json:"achievements"`
	Auth struct {
		Local struct {
			Email          string `json:"email"`
			HashedPassword string `json:"hashed_password"`
			Salt           string `json:"salt"`
			Username       string `json:"username"`
		} `json:"local"`
		Timestamps struct {
			Created  string `json:"created"`
			Loggedin string `json:"loggedin"`
		} `json:"timestamps"`
	} `json:"auth"`
	Backer      struct{}    `json:"backer"`
	Balance     float64     `json:"balance"`
	Challenges  []Challenge `json:"challenges"`
	Contributor struct{}    `json:"contributor"`
	Dailys      []Daily     `json:"dailys"`
	Filters     struct{}    `json:"filters"`
	Flags       struct {
		ClassSelected              bool     `json:"classSelected"`
		CustomizationsNotification bool     `json:"customizationsNotification"`
		DropsEnabled               bool     `json:"dropsEnabled"`
		FreeRebirth                bool     `json:"freeRebirth"`
		ItemsEnabled               bool     `json:"itemsEnabled"`
		LevelDrops                 struct{} `json:"levelDrops"`
		NewStuff                   bool     `json:"newStuff"`
		PartyEnabled               bool     `json:"partyEnabled"`
		RebirthEnabled             bool     `json:"rebirthEnabled"`
		Rewrite                    bool     `json:"rewrite"`
		ShowTour                   bool     `json:"showTour"`
	} `json:"flags"`
	Habits  []Habit `json:"habits"`
	History struct {
		Exp []struct {
			Date  string  `json:"date"`
			Value float64 `json:"value"`
		} `json:"exp"`
		Todos []struct {
			Date  string  `json:"date"`
			Value float64 `json:"value"`
		} `json:"todos"`
	} `json:"history"`
	ID          string `json:"id"`
	Invitations struct {
		Guilds []interface{} `json:"guilds"`
	} `json:"invitations"`
	Items       Item     `json:"items"`
	LastCron    string   `json:"lastCron"`
	NewMessages struct { // TODO: Figure this structure out...
		_2dae4d3_0b7b_4ca3_8378_32b3b697d642 struct {
			Name  string `json:"name"`
			Value bool   `json:"value"`
		} `json:"62dae4d3-0b7b-4ca3-8378-32b3b697d642"`
	} `json:"newMessages"`
	Party struct {
		Order string `json:"order"`
		Quest struct {
			Progress struct {
				Collect struct{} `json:"collect"`
				Down    float64  `json:"down"`
				Up      float64  `json:"up"`
			} `json:"progress"`
		} `json:"quest"`
	} `json:"party"`
	Preferences Preference `json:"preferences"`
	Profile     struct {
		Name string `json:"name"`
	} `json:"profile"`
	Purchased struct {
		Ads        bool     `json:"ads"`
		Background struct{} `json:"background"`
		Hair       struct{} `json:"hair"`
		Plan       struct {
			GemsBought   float64       `json:"gemsBought"`
			MysteryItems []interface{} `json:"mysteryItems"`
		} `json:"plan"`
		Shirt    struct{} `json:"shirt"`
		Skin     struct{} `json:"skin"`
		TxnCount float64  `json:"txnCount"`
	} `json:"purchased"`
	Rewards []Reward `json:"rewards"`
	Stats   Stat     `json:"stats"`
	Tags    []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`
	Todos []Todo `json:"todos"`
}
