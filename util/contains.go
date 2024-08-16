package util

// StringSliceContains receives a slice of string and checks if contains str
func StringSliceContains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}

	return false
}

// IntSliceContains receives a slice of int and checks if contains i
func IntSliceContains(slice []int, i int) bool {
	for _, s := range slice {
		if s == i {
			return true
		}
	}
	return false
}
