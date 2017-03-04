package helpers

import (
	"os/exec"
	"bytes"
	"fmt"
	"os"
)

func RunScript(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error() + ":" + stderr.String())
		os.Exit(1)
	}
	output := out.String()
	//fmt.Print(output)
	return output
}
