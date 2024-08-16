package util

import (
	"strconv"
	"strings"
)

func StringIsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// StrToSliceInt splits a string and converts each element to integer
func StrToSliceInt(s string, sep string) ([]int, error) {
	var sInt []int
	brands := strings.Split(s, sep)

	for _, brand := range brands {
		i, err := strconv.Atoi(brand)
		if err != nil {
			return nil, err
		}
		sInt = append(sInt, i)
	}

	return sInt, nil
}

// MapStr slices a string using delimiter, and creates a map of substrings
func MapStr(str string, delimiter string) map[string]struct{} {
	if str == "" {
		return nil
	}

	m := map[string]struct{}{}

	chunks := strings.Split(str, delimiter)
	for _, k := range chunks {
		m[k] = struct{}{}
	}

	return m
}

// TruncString truncates a string to be limited at most 'size' length
func TruncString(s string, size int) string {
	if len([]rune(s)) > size {
		return string([]rune(s)[:size])
	}
	return s
}

// RemoveXSpace removes extra space in a string
func RemoveExtraSpace(s *string) {
	strings.Join(strings.Fields(strings.TrimSpace(*s)), " ")
}
