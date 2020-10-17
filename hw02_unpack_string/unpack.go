package hw02_unpack_string //nolint:golint,stylecheck
//package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func CheckString(source string) bool {
	//С помощью регулярного выражения определение нужной строки
	reg := regexp.MustCompile(`^(([[:alpha:]]|\s)\d?|\\{2}\d?|\\\d{1,2})*$`)
	res := reg.MatchString(source)
	return res
}

func translateString(source string) (string, error) {
	var escapeSymbol bool
	var resultString strings.Builder
	var lastSymbol rune
	escapeSymbol = false
	if source == "" {
		return "", nil
	}
	for _, r := range source {
		var localbuilder strings.Builder
		localbuilder.WriteRune(r)
		if escapeSymbol {
			escapeSymbol = false
			resultString.WriteRune(r)
			lastSymbol = r
			continue
		}
		if r == '\\' {
			escapeSymbol = true
			continue
		}
		isdigit := unicode.IsDigit(r)
		if isdigit {
			count, err := strconv.Atoi(localbuilder.String())
			if err != nil {
				return "", err
			}
			var lastSymbolBldr strings.Builder
			lastSymbolBldr.WriteRune(lastSymbol)
			lastSymbolStr := lastSymbolBldr.String()
			if count > 0 {
				lastUnpackString := strings.Repeat(lastSymbolStr, count-1)
				resultString.WriteString(lastUnpackString)
			} else {
				lastUnpackString := resultString.String()
				resultString.Reset()
				resultString.WriteString(strings.TrimSuffix(lastUnpackString, lastSymbolStr))
			}
			continue
		}
		lastSymbol = r
		resultString.WriteRune(r)
	}
	return resultString.String(), nil
}

func Unpack(str string) (string, error) {
	// Place your code here
	sourceString := str
	if CheckString(sourceString) {
		translateString, err := translateString(sourceString)
		return translateString, err
	}
	return "", ErrInvalidString
}

func UnpackWithEscape(str string) (string, error) {
	return Unpack(str)
}
