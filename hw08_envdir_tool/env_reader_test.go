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
			t.Error(err)
		}
		name, value, err := readFile("./testdata/env/", info)
		if err != nil {
			t.Error(err)
		}
		if name != "BAR" {
			t.Errorf("Name != BAR %s", name)
		}
		if value != "bar" {
			t.Errorf("%s!=bar", value)
		}
	})
	t.Run("readFileFOO", func(t *testing.T) {
		info, err := os.Stat("./testdata/env/FOO")
		if err != nil {
			t.Error(err)
		}
		name, value, err := readFile("./testdata/env/", info)
		if err != nil {
			t.Error(err)
		}
		if name != "FOO" {
			t.Errorf("Name != FOO %s", name)
		}
		if value != "   foo\nwith new line" {
			t.Errorf("WRONG %s!=   foo\nwith new line", value)
		}
	})
	t.Run("readEmptyFile", func(t *testing.T) {
		info, err := os.Stat("./testdata/env/UNSET")
		if err != nil {
			t.Error(err)
		}
		name, value, err := readFile("./testdata/env/", info)
		if err != nil {
			t.Error(err)
		}
		if name != "UNSET" {
			t.Errorf("Name != FOO %s", name)
		}
		if value != "" {
			t.Errorf("no empty line %s", value)
		}
	})
}
