package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func Shellout(command string) string {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.String()
	}
	return stdout.String()
}
func main() {
	os.Chdir("/tmp")
	a := Shellout("ls -l")
	fmt.Print(a)
}
