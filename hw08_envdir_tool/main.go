package main

import (
	"log"
	"os"
)

func main() {
	// Place your code here
	args := os.Args[1:]
	var err error
	if env, err := ReadDir(args[0]); err == nil {
		RunCmd(args[1:], env)
		return
	}
	log.Fatal(err)
}
