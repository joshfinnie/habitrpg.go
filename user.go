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
	stats struct {
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
	profile struct {
		name string
	}
	preferences struct {
		automaticAllocation bool
		costume             bool
		language            string
		timezoneOffset      int
		toolbarCollapsed    bool
		advancedCollapsed   bool
		tagsCollapsed       bool
		newTaskEdit         bool
		disabledClasses     bool
		stickyHeader        bool
		sleep               bool
		allocationMode      string
		sound               string
		shirt               string
		skin                string
		hideHeader          bool
		hair                struct {
			flower   int
			mustache int
			beard    int
			bangs    int
			color    string
		}
		size     string
		dayStart int
	}
	party struct {
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
	lastCron string
	items    struct {
		currentPet string
		lastDrop   struct {
			count int
			date  string
		}
		quests interface{}
		mounts interface{}
		food   struct {
			Candy_Base            int
			Candy_CottonCandyBlue int
			Candy_CottonCandyPink int
			Candy_Desert          int
			Candy_Golden          int
			Candy_Red             int
			Candy_Shade           int
			Candy_Skeleton        int
			Candy_White           int
			Candy_Zombie          int
			Chocolate             int
			CottonCandyBlue       int
			CottonCandyPink       int
			Fish                  int
			Meat                  int
			RottenMeat            int
			Strawberry            int
		}
		hpatchingPotions struct {
			Base            int
			CottonCandyPink int
			Desert          int
			Red             int
			Shade           int
			Skeleton        int
			White           int
			Zombie          int
		}
		eggs struct {
			BearCub   int
			Cactus    int
			Dragon    int
			FlyingPig int
			Fox       int
			LionCub   int
			TigerCub  int
			Wolf      int
		}
		pets struct {
			BearCub_Skeleton     int
			Cactus_Desert        int
			Cactus_Red           int
			Dragon_Base          int
			Dragon_Red           int
			Dragon_Zombie        int
			FlyingPig_Desert     int
			FlyingPig_Shade      int
			Fox_Desert           int
			TigerCub_White       int
			Wolf_Base            int
			Wolf_CottonCandyPink int
		}
		special struct {
			valentineReceived []string
			snowball          int
		}
		gear struct {
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
	}
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
