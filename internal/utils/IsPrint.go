package utils

import "fmt"

func IsPrint(s string) bool {
	for _, val := range s {
		if val < ' ' || val > '~' {
			fmt.Println("error here :", val)
			return false
		}
	}
	return true
}
