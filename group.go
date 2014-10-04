// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Group
package habitrpg

type Group struct {
	V              float64 `json:"__v"`
	ID             string  `json:"_id"`
	ChallengeCount float64 `json:"challengeCount"`
	Challenges     []struct {
		ID   string `json:"_id"`
		Name string `json:"name"`
	} `json:"challenges"`
	Chat []struct {
		ID        string   `json:"id"`
		Likes     struct{} `json:"likes"`
		Text      string   `json:"text"`
		Timestamp float64  `json:"timestamp"`
		Uuid      string   `json:"uuid"`
	} `json:"chat"`
	Invites     []interface{} `json:"invites"`
	Leader      string        `json:"leader"`
	MemberCount float64       `json:"memberCount"`
	Members     []User        `json:"members"`
	Name        string        `json:"name"`
	Privacy     string        `json:"privacy"`
	Quest       struct {
		Active   bool `json:"active"`
		Progress struct {
			Collect struct{} `json:"collect"`
		} `json:"progress"`
	} `json:"quest"`
	Type string `json:"type"`
}
