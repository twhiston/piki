package helpers

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func RunScript(name string, args ...string) string {
	return RunScriptInDirectory(name, "", args...)
}

func RunScriptInDirectory(name string, dir string, args ...string) string {

	cmd := exec.Command(name, args...)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if dir != "" {
		cmd.Dir = dir
	}
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error() + ":" + stderr.String())
		fmt.Println(out.String())
		os.Exit(1)
	}
	output := out.String()
	//fmt.Print(output)
	return output
}
