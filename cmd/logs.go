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

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Helper command to quickly bring up useful logs",
	Long:  ``,
}

var httpdLogsCmd = &cobra.Command{
	Use:   "httpd",
	Short: "view httpd log",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		viewLog(cmd, "/var/log/lighttpd/lighttpd.log")
	},
}

var phpLogsCmd = &cobra.Command{
	Use:   "php",
	Short: "view httpd log",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("currently php issues are logged to stdout. This will change")

	},
}

func viewLog(cmd *cobra.Command, location string) {

	tail, _ := cmd.Parent().PersistentFlags().GetBool("tail")
	if tail {
		helpers.RunScript("/bin/sh", "-c", "tail -F "+location)
	} else {
		helpers.RunScript("/bin/sh", "-c", "cat "+location)
	}

}

func init() {
	RootCmd.AddCommand(logsCmd)

	logsCmd.PersistentFlags().Bool("tail", false, "Opens a tail on the logs to see any fresh messages")

	logsCmd.AddCommand(httpdLogsCmd)
	logsCmd.AddCommand(phpLogsCmd)

}
