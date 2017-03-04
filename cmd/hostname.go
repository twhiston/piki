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


	"io/ioutil"
	"github.com/twhiston/piki/helpers"
	"regexp"
	"fmt"
)

// hostnameCmd represents the hostname command
var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		fmt.Println("Set hostname to "+ hostname)
		fmt.Println("You should run piki reboot")


	},
}

func init() {
	RootCmd.AddCommand(hostnameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	hostnameCmd.PersistentFlags().String("name", "piki", "Hostname for the device")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostnameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
