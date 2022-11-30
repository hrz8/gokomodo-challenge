package parser

import "strconv"

func ParseStringToInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return result
}
