package util

import "strconv"

func DiscoverDigit(code string) (string, error) {
	var (
		mult int64 = 2
		sum  int64 = 0
		i    int
	)
	for i = (len(code) - 1); i >= 0; i-- {
		c, err := strconv.ParseInt(string(code[i]), 10, 32)
		if err != nil {
			return "", err
		}

		rest := c * mult
		if rest > 9 {
			rest = (rest - 9)
		}
		sum += rest

		if mult == 2 {
			mult = 1
		} else {
			mult = 2
		}
	}

	sum = (sum % 10)
	if sum == 0 {
		return "0", nil
	}

	sum = 10 - sum
	return strconv.FormatInt(int64(sum), 10), nil
}
