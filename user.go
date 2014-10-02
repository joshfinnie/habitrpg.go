// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's User
package habitrpg

type User struct {
	__v        int
	_id        string
	rewards    []Reward
	todos      []Todo
	dailys     []Daily
	habits     []Habit
	challenges []string
	tags       struct {
		name string
		id   string
	}
	stats   Stat
	profile struct {
		name string
	}
	preferences Preference
	party       struct {
		quest struct {
			progress struct {
				collect interface{}
				down    int
				up      int //float
			}
		}
		order string
	}
	newMessages struct {
		value bool
		name  string
	}
	lastCron    string
	items       Item
	inwitations []string
	history     struct {
		todos []struct {
			value int //float
			date  int
		}
		exp []struct {
			value int //float
			date  int
		}
	}
	flags struct {
		partyEnabled               bool
		levelDrops                 interface{}
		freeRebirth                bool
		rebirthEnabled             bool
		classSelected              bool
		rewrite                    bool
		newStuff                   bool
		itemsEnabled               bool
		dropsEnabled               bool
		showTour                   bool
		customizationsNotification bool
	}
	purchased struct {
		plan struct {
			mysteryItems []string
			gemsBought   int
		}
		txnCount   int
		background interface{}
		shirt      interface{}
		hair       interface{}
		skin       interface{}
		ads        bool
	}
	filters     interface{}
	balance     int
	contributor interface{}
	backer      interface{}
	auth        struct {
		timestamps struct {
			loggedin string
			created  string
		}
		local struct {
			username        string
			email           string
			salt            string
			hashed_password string
		}
	}
	achievements struct {
		beastmaster bool
		perfect     int
		challenges  []Challenge
	}
	_v int
	id string
}
