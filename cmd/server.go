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
	Short: "interact with the lighttpd server",
}

// serverCmd represents the server command
var ctrlServerCmd = &cobra.Command{
	Use:   "ctrl",
	Short: "stop start or restart the server",
	Long: `simple alias for /etc/init.d/lighttp.
Arguments could be
	start
	stop
	restart`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("argument must be start / stop / restart")
			return
		}

		command := args[0]

		if command == "stop" {
			fmt.Print(helpers.RunScript("/bin/sh", "-c", "sudo /etc/init.d/lighttpd stop"))
		} else if command == "start" {
			fmt.Print(helpers.RunScript("/bin/sh", "-c", "sudo /etc/init.d/lighttpd start"))
		} else if command == "restart" {
			fmt.Print(helpers.RunScript("/bin/sh", "-c", "sudo /etc/init.d/lighttpd restart"))
		} else {
			fmt.Println("argument must be start / stop / restart")
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
		eCmd := exec.Command("/bin/sh", "-c", "sudo "+editor+" "+file)
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
	serverCmd.AddCommand(ctrlServerCmd)

	editServerCmd.PersistentFlags().String("conf", "/etc/lighttpd/lighttpd.conf", "Path and filename of lighttpd conf file")
	editServerCmd.PersistentFlags().String("editor", "vi", "Name of editor to use")

}
