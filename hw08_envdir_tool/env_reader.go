package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

type Environment map[string]string

var (
	ErrContainsEq = errors.New("file contains =")
	ErrDirectory  = errors.New("directory name file")
)

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	// Place your code here
	retval := make(Environment)
	info, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("ReadDir get error from ioutil.ReadDir %w", err)
	}
	for _, fileInfo := range info {
		if strings.ContainsRune(fileInfo.Name(), '=') {
			return nil, ErrContainsEq
		}
		if fileInfo.IsDir() {
			return nil, ErrDirectory
		}
		key, value, err := readFile(dir, fileInfo)
		if err != nil {
			return nil, fmt.Errorf("ReadDir brings error %w", err)
		}
		retval[key] = value
	}
	return retval, nil
}

func readFile(dir string, info os.FileInfo) (fileName string, value string, err error) {
	fileName = info.Name()
	content, err := ioutil.ReadFile(dir + "/" + fileName)
	if err != nil {
		return "", "", fmt.Errorf("readFile get error from ioutil.ReadFile %w", err)
	}
	if len(content) == 0 {
		return fileName, "", nil
	}
	sliceByte := make([]byte, 4)
	sizeSliceWritten := utf8.EncodeRune(sliceByte, '\n')
	sliceByte = sliceByte[:sizeSliceWritten]
	zeroslice := make([]byte, 1)
	reader := bytes.NewReader(content)
	bufioreader := bufio.NewReader(reader)
	var builder strings.Builder
	for {
		line, readFine, err := bufioreader.ReadLine()
		if err != nil {
			return "", "", fmt.Errorf("readFile ReadLine error %w", err)
		}
		line = bytes.ReplaceAll(line, zeroslice, sliceByte)
		_, err = builder.Write(line)
		if err != nil {
			return "", "", fmt.Errorf("reafFile string builder failed %w", err)
		}
		if !readFine {
			break
		}
	}
	value = builder.String()
	value = strings.TrimRight(value, " \t")

	return fileName, value, nil
}
