package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func getParams() (string, string, string) {

	account := os.Args[1]
	project := os.Args[2]
	version := os.Args[3]

	return account, project, version
}

func appendC(args ...string) []string {
	cmdArgs := args
	diskFlag := []string{"/c"}

	cmdArgs = append(diskFlag, cmdArgs...)
	return cmdArgs
}

func runCmd(args ...string) {
	cmdArgs := appendC(args...)
	cmd := exec.Command("cmd", cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println(string(output))
}

func main() {
	account, project, version := getParams()
	runCmd("docker", "build", "-t", fmt.Sprintf("ghcr/io/%v/%v/v%v", account, project, version), ".")
	runCmd("docker", "push", fmt.Sprintf("ghcr/io/%v/%v/v%v", account, project, version))
}
