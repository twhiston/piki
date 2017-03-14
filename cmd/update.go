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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update the code of the dashboard via git",
	Long:  `does a git pull for the api (or defined) folder, does a hard reset on flag`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Parent().PersistentFlags().GetString("dir")
		hard, _ := cmd.PersistentFlags().GetBool("hard")
		fmt.Print(helpers.RunScriptInDirectory("/bin/sh", dir, "-c", "sudo chown -R pi .git"))
		if hard {
			fmt.Print(helpers.RunScriptInDirectory("/bin/sh", dir, "-c", "git fetch"))
			branch, _ := cmd.PersistentFlags().GetString("branch")
			fmt.Print(helpers.RunScriptInDirectory("/bin/sh", dir, "-c", "git reset --hard "+branch))
		}
		fmt.Print(helpers.RunScriptInDirectory("/bin/sh", dir, "-c", "git pull"))
	},
}

func init() {
	dashCmd.AddCommand(updateCmd)

	updateCmd.PersistentFlags().Bool("hard", false, "If hard is set this will reset the repo before pulling the newest version to ensure all changes are gone")
	updateCmd.PersistentFlags().String("branch", "origin/master", "The branch to hard reset to")

}
