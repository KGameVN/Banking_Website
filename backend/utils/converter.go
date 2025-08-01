package utils

import "strconv"

func StringToInt(in string) (int,error) {
	number, err := strconv.Atoi(in)
	if err != nil {
		return -1, err
	}
	return number, nil
}

func StringToInt64(in string) int64 {
	return 0
}