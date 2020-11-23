package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	// Place your code here
	retval := make(Environment)
	info, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range info {
		key, value, err := readFile(dir, fileInfo)
		if err != nil {
			return nil, fmt.Errorf("ReadDir btings error %v", err)
		}
		retval[key] = value
	}
	return retval, nil
}

func readFile(dir string, info os.FileInfo) (fileName string, value string, err error) {
	fileName = info.Name()
	content, err := ioutil.ReadFile(dir + "/" + fileName)
	if err != nil {
		return "", "", fmt.Errorf("readFile get error from ioutil.ReadFile %v", err)
	}
	sliceByte := make([]byte, 4)
	utf8.EncodeRune(sliceByte, '\n')

	var builder strings.Builder

	_, err = builder.Write(content)
	if err != nil {
		return "", "", fmt.Errorf("readFile get error from string builder write method %v", err)
	}

	str := strings.Split(builder.String(), "\n")[0]

	builder.Reset()

	for i := 0; i < len(str); i++ {
		if byteValue := str[i]; byteValue == 0 {
			_, err = builder.WriteRune('\n')
		} else {
			err = builder.WriteByte(byteValue)
		}
		if err != nil {
			return "", "", fmt.Errorf("readFile get error trying write to string string builder %v", err)
		}
	}
	value = builder.String()
	value = strings.TrimRight(value, " \t")

	return fileName, value, nil
}
