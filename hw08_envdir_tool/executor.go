package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env...
func RunCmd(cmd []string, env Environment) (returnCode int) {
	// Place your code here
	env.applyEnvironment()
	cmdName := cmd[0]
	command := exec.Command(cmdName)
	command.Args = cmd
	command.Env = os.Environ()
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
	return
}
func (env Environment) applyEnvironment() {
	for key, value := range env {
		if checkVariable(key) {
			os.Unsetenv(key)
		}
		if value != "" {
			os.Setenv(key, value)
		}
	}
}

func checkVariable(key string) bool {
	for _, str := range os.Environ() {
		if strings.Contains(str, key) {
			return true
		}
	}
	return false
}
