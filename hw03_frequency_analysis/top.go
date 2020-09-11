package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
)

func splitToMap(str string, regstr string) (map[string]int32, error) {
	if len(str) == 0 {
		return nil, error.Error()
	}
	var s = make(map[string]int32)
	reg := regexp.MustCompile(regstr)
	words := reg.Split(str, -1)
	for _, word := range words {
		s[word]++
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
		delete(mapin, keyStored)
		return keyStored
	}
	for i := 0; i < 10; i++ {
		res[i] = findMax()
	}
	return res
}

func Top10(instring string) []string {
	return Result1(splitToMap(instring, `\s`))
}
