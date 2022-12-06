package helper

import "strconv"

func IntCoverter(val string) int {
	ret, err := strconv.Atoi(val)
	if err != nil {
		return -1
	}
	return ret
}
