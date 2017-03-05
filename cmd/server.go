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
	"os"
	"os/exec"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "stop start or restart the server",
	Long:  `simple alias for /etc/init.d/lighttp`,
	Run: func(cmd *cobra.Command, args []string) {

		stop, _ := cmd.PersistentFlags().GetBool("stop")
		start, _ := cmd.PersistentFlags().GetBool("start")
		restart, _ := cmd.PersistentFlags().GetBool("restart")

		if !stop && !start && !restart {
			fmt.Println("You must specify a --stop --start or --restart flag")
		}

		if stop {
			fmt.Print(helpers.RunScript("/bin/sh", "-c", "/etc/init.d/lighttpd", "stop"))
		}
		if start {
			fmt.Print(helpers.RunScript("/bin/sh", "-c", "/etc/init.d/lighttpd", "start"))
		}
		if restart {
			fmt.Print(helpers.RunScript("/bin/sh", "-c", "/etc/init.d/lighttpd", "restart"))
		}

	},
}

var editServerCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit the server config",
	Long:  `simple alias for sudo vi /etc/lighttpd/lighttpd.conf`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.PersistentFlags().GetString("conf")
		editor, err := cmd.PersistentFlags().GetString("editor")
		eCmd := exec.Command(editor, file)
		eCmd.Stdin = os.Stdin
		eCmd.Stdout = os.Stdout
		err = eCmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(editServerCmd)

	serverCmd.PersistentFlags().Bool("restart", false, "A help for foo")
	serverCmd.PersistentFlags().Bool("start", false, "A help for foo")
	serverCmd.PersistentFlags().Bool("stop", false, "A help for foo")

	editServerCmd.PersistentFlags().String("conf", "/etc/lighttpd/lighttpd.conf", "Path and filename of lighttpd conf file")
	editServerCmd.PersistentFlags().String("editor", "vi", "Name of editor to use")

}
