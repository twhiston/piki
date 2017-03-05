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
	"io/ioutil"
	"regexp"
)

// bootCmd represents the boot command
var bootCmd = &cobra.Command{
	Use:   "boot",
	Short: "A brief description of your command",
	Long:  `Set if you want to boot to the app or rescue mode easily with this tool. Just pass in the type of boot you want`,
	Run: func(cmd *cobra.Command, args []string) {
		// read the whole file at once
		cmdLineFile, err := cmd.PersistentFlags().GetString("file")
		cmdLineFileString := helpers.GetFileAsString(cmdLineFile)
		bootType, _ := cmd.PersistentFlags().GetString("type")
		if bootType == "" {
			fmt.Println("--type cannot be blank")
			return
		}

		if !helpers.StringExists("rootwait", cmdLineFileString) {
			//if the rootwait string does not exist then we need to add it
			return
		}

		if bootType == "app" {
			//Regex must match rootwait 1
			re := regexp.MustCompile("(?m)rootwait 1")
			cmdLineFileString = re.ReplaceAllString(cmdLineFileString, "rootwait")
		} else if bootType == "recovery" {

			//Regex must match rootwait but NOT rootwait 1
			re := regexp.MustCompile("(?m)rootwait")
			cmdLineFileString = re.ReplaceAllString(cmdLineFileString, "rootwait 1")

		} else {
			fmt.Println("Sorry but i dont understand this boot type. Exiting")
			return
		}

		err = ioutil.WriteFile(cmdLineFile, []byte(cmdLineFileString), 0777)
		if err != nil {
			panic(err)
		}

		fmt.Println("Set boot mode to " + bootType)
	},
}

func init() {
	RootCmd.AddCommand(bootCmd)

	bootCmd.PersistentFlags().String("file", "/boot/cmdline.txt", "Path (including filename) of the cmdline file to alter")
	bootCmd.PersistentFlags().String("type", "app", "The type of boot. Accepted are app & recovery")

}
