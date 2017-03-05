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
	"os"
	"regexp"
)

// netCmd represents the net command
var netCmd = &cobra.Command{
	Use:   "net",
	Short: "set up network info",
	Long:  `Argument must be wpa, wep, open, eth`,
	Run: func(cmd *cobra.Command, args []string) {

		net, _ := cmd.Flags().GetString("network")
		pass, _ := cmd.Flags().GetString("pass")

		address, _ := cmd.Flags().GetString("address")
		netmask, _ := cmd.Flags().GetString("netmask")
		broadcast, _ := cmd.Flags().GetString("broadcast")

		output := ""

		if len(args) == 0 {
			fmt.Println("You must specify the type of network connection you wish to establish")
			os.Exit(1)
		}

		switch args[0] {

		case "wpa":
			output = `
iface wlan0-fullpageos inet manual
    wpa-ssid "` + net + `"
    wpa-psk "` + pass + `"`
			break
		case "wep":
			output = `
iface wlan0-fullpageos inet manual
    wireless-essid "` + net + `"
    wireless-key "` + pass + `"`
			break
		case "open":
			output = `
iface wlan0-fullpageos inet manual
    wireless-essid "` + net + `"
    wireless-mode managed`
			break
		case "eth":
			output = `
auto eth0:1
iface eth0:1 inet static
  address ` + address + `
  netmask ` + netmask + `
  broadcast ` + broadcast
			break
		default:
			fmt.Println("Argument must be wpa, wep, open, eth")
			return
		}

		filePath, _ := cmd.PersistentFlags().GetString("file")
		networkFileString := helpers.GetFileAsString(filePath)

		//Find in the file a specific string
		headerString := "## PIKI AUTO GENERATED ENTRY - DO NOT EDIT MANUALLY ##"

		if !helpers.StringExists(headerString, networkFileString) {
			networkFileString = networkFileString + headerString
		}

		re := regexp.MustCompile(headerString + "(.|\n)*")
		networkFileString = re.ReplaceAllString(networkFileString, headerString+`
		`+output)

		err := ioutil.WriteFile(filePath, []byte(networkFileString), 0777)
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	RootCmd.AddCommand(netCmd)

	netCmd.PersistentFlags().String("file", "/boot/piki-network.txt", "The network include file")
	netCmd.Flags().String("network", "", "The name of the wireless network to connect to")
	netCmd.Flags().String("pass", "", "The password of the wireless network to connect to")
	netCmd.Flags().String("address", "192.168.250.1", "The address of the wired network to connect to")
	netCmd.Flags().String("netmask", "255.255.255.0", "The netmask of the wired network to connect to")
	netCmd.Flags().String("broadcast", "192.168.250.255", "The broadcast ip of the wireless network to connect to")
	//netCmd.Flags().Bool("add",false,"if true then the network described will be added to the network list rather than replaced")

}
