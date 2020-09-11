package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"errors"
	"regexp"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

func splitToMap(str string, regstr string) (map[string]int32, error) {
	if len(str) == 0 {
		err := errors.New("")
		return nil, err
	}
	var s = make(map[string]int32)
	reg := regexp.MustCompile(regstr)
	words := reg.Split(str, -1)
	for _, word := range words {
		if len(word) > 0 {
			s[word]++
		}
	}
	return s, nil
}
func splitAdvToMap(str string, regstr string) (map[string]int32, error) {
	if len(str) == 0 {
		err := errors.New("")
		return nil, err
	}
	var s = make(map[string]int32)
	reg := regexp.MustCompile(regstr)
	words := reg.Split(str, -1)
	for _, word := range words {
		if len(word) > 0 {
			s[word]++
		}
	}
	return s, nil
}

func Result1(mapin map[string]int32) []string {
	res := make([]string, 10)

	findMax := func() string {
		var max int32
		var keyStored string
		for key, value := range mapin {
			if max < value {
				keyStored = key
				max = value
			}
		}
		mapin[keyStored] = 0
		return keyStored
	}
	for i := 0; i < 10; i++ {
		res[i] = findMax()
	}
	return res
}

func Top10(instring string) []string {
	res, err := splitToMap(instring, `\s`)
	if err == nil {
		return Result1(res)
	}
	//log.Fatalf("Error occurred!!!")
	return nil
}
