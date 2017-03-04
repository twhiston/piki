package helpers

import (
	"strings"
	"io/ioutil"
	"fmt"
	"os"
)

func StringExists(needle string, haystack string) bool {
	//check if env variable exists
	return strings.Contains(haystack, needle)
}

func GetFileAsString(pathToFile string) (string) {
	envFileData, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		fmt.Println("Failed to open file. You may need to run with sudo")
		os.Exit(1)
	}
	return string(envFileData)
}
