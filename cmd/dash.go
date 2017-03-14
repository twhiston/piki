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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/twhiston/piki/helpers"
)

// dashCmd represents the dash command
var dashCmd = &cobra.Command{
	Use:   "dash",
	Short: "commands to control the dashboard are below this namespace",
	Long:  ``,
}

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "refresh the browser screen",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//export DISPLAY=:0
		//WID=$(xdotool search --onlyvisible --class chromium|head -1)
		//xdotool windowactivate ${WID}
		//xdotool key ctrl+F5
		//xdotool key F11
		fmt.Println(helpers.RunScript("/bin/sh", "-c", "export DISPLAY=:0"))
		fmt.Println(helpers.RunScript("/bin/sh", "-c", "WID=$(xdotool search --onlyvisible --class chromium|head -1); xdotool windowactivate ${WID}"))
		fmt.Println(helpers.RunScript("/bin/sh", "-c", "xdotool key ctrl+F5"))
		fmt.Println(helpers.RunScript("/bin/sh", "-c", "xdotool key F11"))
	},
}

func init() {
	RootCmd.AddCommand(dashCmd)
	dashCmd.PersistentFlags().String("dir", "/var/www/html/PikiDashboard", "Directory of the dasboard")
	dashCmd.AddCommand(refreshCmd)
}
