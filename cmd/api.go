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
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Interact with the defaut dashboard app api",
	Long:  `Only works with the default app api, so if you use something else on your piki you should ignore this command`,
}

func init() {
	dashCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	apiCmd.PersistentFlags().String("base", "http://localhost/", "Set the base path, MUST end with /")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func validateBasePath(cmd *cobra.Command) string {
	base, err := cmd.Flags().GetString("base")

	if err != nil {
		fmt.Println("no base path is available")
		os.Exit(1)
	}

	return base
}

func RenderApiGetCall(base string, path string) {

	requestAddress := base + path
	resp, err := http.Get(requestAddress)
	fmt.Println("GET: " + requestAddress)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

}

func RenderApiPostCall(base string, path string, payload string) {

	body := strings.NewReader(payload)

	requestAddress := base + path
	resp, err := http.Post(requestAddress, "application/json", body)
	fmt.Println("POST: " + requestAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

}
