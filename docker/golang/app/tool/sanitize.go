package tool

import (
	"strings"
)

func SaniAtSign(s string) string {
	replaced := strings.Replace(s, "@", "(a)", 1)
	return replaced
}

func RevSaniAtSign(s string) string {
	replaced := strings.Replace(s, "(a)", "@", 1)
	return replaced
}
