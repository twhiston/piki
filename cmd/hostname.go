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
	"github.com/twhiston/piki/helpers"
	"io/ioutil"
	"regexp"
)

// hostnameCmd represents the hostname command
var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "Set your Piki machine hostname",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		currentHost := helpers.GetFileAsString("/etc/hostname")
		hostname, err := cmd.PersistentFlags().GetString("name")

		hostsFileString := helpers.GetFileAsString("/etc/hosts")

		//Write the hostname to the hostname file
		err = ioutil.WriteFile("/etc/hostname", []byte(hostname), 0644)
		if err != nil {
			panic(err)
		}

		re := regexp.MustCompile("(?m)" + currentHost)
		res := re.ReplaceAllString(hostsFileString, hostname)

		err = ioutil.WriteFile("/etc/hosts", []byte(res), 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println("Set hostname to " + hostname)
		fmt.Println("You should run piki reboot")

	},
}

func init() {
	RootCmd.AddCommand(hostnameCmd)

	hostnameCmd.PersistentFlags().String("name", "piki", "Hostname for the device")
}
