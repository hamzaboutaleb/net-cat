package utils

func IsValidPort(port string) bool {
	return port != "0" && isNumeric(port)
}
