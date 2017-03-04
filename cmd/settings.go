// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (

	"github.com/spf13/cobra"
	"fmt"
)

// settingsCmd represents the settings command
var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Interact with the dashboard settings via the API",
	Long: `Uses the API to interact with the dashboard app`,

}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the settings",
	Long: `Get the api settings, or optionally a specific setting`,
	Run: func(cmd *cobra.Command, args []string) {

		base := validateBasePath(cmd)
		RenderApiGetCall(base,"api/settings")
	},
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a settings value",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		base := validateBasePath(cmd)

		setting, err := cmd.Parent().PersistentFlags().GetString("name")
		if(setting == ""){
			fmt.Println("--setting cannot be null")
		}
		if(err != nil){
			panic(err)
		}

		value, err := cmd.PersistentFlags().GetString("value")

		//Yeah that sucks to build some json this way but that's whats happening for now
		json := "{\""+setting+"\": \""+value+"\"}"
		fmt.Println(json)
		RenderApiPostCall(base,"api/settings",json)
	},
}

func init() {
	apiCmd.AddCommand(settingsCmd)
	settingsCmd.AddCommand(setCmd)
	settingsCmd.AddCommand(getCmd)

	settingsCmd.PersistentFlags().String("name","","the name of the setting")
	setCmd.PersistentFlags().String("value","","the value to set")
}
