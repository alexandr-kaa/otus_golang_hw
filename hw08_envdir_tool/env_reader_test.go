package main

import (
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {
	// Place your code here
	_, err := ReadDir("./testdata/env")
	if err != nil {
		t.Errorf("Error! %v", err.Error())
	}
}
func TestReadFile(t *testing.T) {

	t.Run("readFileBAR", func(t *testing.T) {
		info, err := os.Stat("./testdata/env/BAR")
		if err != nil {
			t.Fatal(err)
		}
		name, value, err := readFile("./testdata/env/", info)
		if err != nil {
			t.Fatal(err)
		}
		if name != "BAR" {
			t.Fatalf("Name != BAR %s", name)
		}
		if value != "bar" {
			t.Fatalf("%s!=bar", value)
		}
	})
	t.Run("readFileFOO", func(t *testing.T) {
		info, err := os.Stat("./testdata/env/FOO")
		if err != nil {
			t.Fatal(err)
		}
		name, value, err := readFile("./testdata/env/", info)
		if err != nil {
			t.Fatal(err)
		}
		if name != "FOO" {
			t.Fatalf("Name != FOO %s", name)
		}
		if value != "   foo\nwith new line" {
			t.Fatalf("WRONG %s!=   foo\nwith new line", value)
		}
	})
}
