package utils

func IsPrint(s string) bool {
	for _, val := range s {
		if val < ' ' || val > '~' {
			return false
		}
	}
	return true
}
