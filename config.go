// Copyright 2014 Josh Finnie and contributors. All rights reserved.
// Contributers can be found listed in the CONTRIBUTORS.txt document
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This file contains the struct for HabitRPG's Task
package habitrpg

type Config struct {
	UserId   string `json:"userId"`
	ApiToken string `json:"apiToken"`
}

// Builds our config struct
func getConfig() Config {

	file, _ := os.Open("config.json")
	contents, _ := ioutil.ReadAll(file)

	var config Config
	json.Unmarshal(contents, &config)

	return config
}
