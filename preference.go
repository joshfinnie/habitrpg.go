// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Preference
package habitrpg

type Preference struct {
	AdvancedCollapsed   bool    `json:"advancedCollapsed"`
	AllocationMode      string  `json:"allocationMode"`
	AutomaticAllocation bool    `json:"automaticAllocation"`
	Costume             bool    `json:"costume"`
	DayStart            float64 `json:"dayStart"`
	DisableClasses      bool    `json:"disableClasses"`
	Hair                struct {
		Bangs    float64 `json:"bangs"`
		Base     float64 `json:"base"`
		Beard    float64 `json:"beard"`
		Color    string  `json:"color"`
		Flower   float64 `json:"flower"`
		Mustache float64 `json:"mustache"`
	} `json:"hair"`
	HideHeader       bool    `json:"hideHeader"`
	Language         string  `json:"language"`
	NewTaskEdit      bool    `json:"newTaskEdit"`
	Shirt            string  `json:"shirt"`
	Size             string  `json:"size"`
	Skin             string  `json:"skin"`
	Sleep            bool    `json:"sleep"`
	Sound            string  `json:"sound"`
	StickyHeader     bool    `json:"stickyHeader"`
	TagsCollapsed    bool    `json:"tagsCollapsed"`
	TimezoneOffset   float64 `json:"timezoneOffset"`
	ToolbarCollapsed bool    `json:"toolbarCollapsed"`
}
