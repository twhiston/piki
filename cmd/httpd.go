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
	"io/ioutil"
	"os"
)

// httpdCmd represents the httpd command
var httpdCmd = &cobra.Command{
	Use:   "httpd",
	Short: "Turn on or off httpd checking when starting up OS",
	Long:  `Set flag --disabled to turn off the checking and always allow the browser to start`,
	Run: func(cmd *cobra.Command, args []string) {
		httpdFile, err := cmd.PersistentFlags().GetString("file")

		if _, err := os.Stat(httpdFile); os.IsNotExist(err) {
			f, _ := os.Create(httpdFile)
			f.Close()
		}

		disabled, err := cmd.PersistentFlags().GetBool("disabled")

		var httpdFileString = ""
		if disabled {
			httpdFileString = "disabled"
		} else {
			httpdFileString = "enabled"
		}

		err = ioutil.WriteFile(httpdFile, []byte(httpdFileString), 0777)
		if err != nil {
			panic(err)
		}

		fmt.Println("Set check_for_httpd to " + httpdFileString)

	},
}

func init() {
	RootCmd.AddCommand(httpdCmd)
	httpdCmd.PersistentFlags().String("file", "/boot/check_for_httpd", "Path (including filename) of the check_for_httpd file to alter")
	httpdCmd.PersistentFlags().Bool("disabled", false, "set the disable flag to set the value to disabled")
}
