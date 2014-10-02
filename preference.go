// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Preference
package habitrpg

type Preference struct {
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
