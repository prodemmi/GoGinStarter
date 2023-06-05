package utils

import "strconv"

func IStoI(value string) int {
	if value == "" {
		return 0
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return i
}
