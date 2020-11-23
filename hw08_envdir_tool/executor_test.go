package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	c := []string{"ls", "-l"}
	testMap := Environment{"FOO": "foo"}
	err := RunCmd(c, testMap)
	if err != 0 {
		t.Error(err)
	}
	// Place your code here

}
func TestApplyEnvironment(t *testing.T) {
	testMap := Environment{"FOO": "foo", "BAR": "bar"}
	testMap.applyEnvironment()
	t.Run("checkNewVars", func(t *testing.T) {
		pairs := os.Environ()
		var match int
		for i, pair := range pairs {
			if pair == "FOO=foo" || pair == "BAR=bar" {
				fmt.Print(i)
				match++
			}
		}
		if match != 2 {
			t.Fatal("Mismatch!!!")
		}
	})
	testMap = Environment{"FOO": ""}
	testMap.applyEnvironment()
	t.Run("removeVar", func(t *testing.T) {
		for _, pair := range os.Environ() {
			if strings.Index(pair, "FOO") != -1 {
				t.Fatal("Find FOO")
			}
		}
	})

}
