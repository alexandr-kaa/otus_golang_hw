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
		err := errors.New("empty string passed")
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

func analyzeWord(list []string, word string) ([]string, string, error) {
	if word == "" {
		err := errors.New("empty string passed")
		return list, word, err
	}
	reg := regexp.MustCompile("(?i)^" + word + "$")
	var match = ""
	for _, wrd := range list {
		if reg.MatchString(wrd) {
			match = wrd
			break
		}
	}
	if len(match) > 0 {
		return list, match, nil
	}
	list = append(list, word)

	return list, word, nil
}

func splitAdvToMap(str string, regstr string) (map[string]int32, error) {
	if len(str) == 0 {
		err := errors.New("empty string passed")
		return nil, err
	}
	var s = make(map[string]int32)
	var strlist = make([]string, 0)
	reg := regexp.MustCompile(regstr)
	words := reg.Split(str, -1)
	for _, word := range words {
		if len(word) > 0 && word != "-" {
			reg = regexp.MustCompile(`((\w|[а-яА-Я])(\w|-|[а-яА-Я])*)`)
			ok := reg.MatchString(word)
			if ok {
				fs := reg.FindString(word)
				var err error
				var findw string
				strlist, findw, err = analyzeWord(strlist, fs)
				if err != nil {
					return nil, err
				}
				s[findw]++
			}
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
	var taskWithAsteriskIsCompleted = true
	var res map[string]int32
	var err error
	if taskWithAsteriskIsCompleted {
		res, err = splitAdvToMap(instring, `\s`)
	} else {
		res, err = splitToMap(instring, `\s`)
	}
	if err == nil {
		return Result1(res)
	}
	return nil
}
