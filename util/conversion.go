package util

import (
	"fmt"
	"strconv"
)

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil{
		fmt.Printf("Couldn't convert %s to int", s)
	}

	return i
}
