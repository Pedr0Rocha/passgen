package password

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits       = "0123456789"
	Symbols      = "@&$#!%*._+-=()[]{}:?"
)

func Generate(passLength int, hasSymbols bool) (string, error) {
	tries := 10_000

	password := ""
	for tries > 0 {
		password = buildRandomPassword(passLength, hasSymbols)
		isValid := isValid(password, passLength, hasSymbols)

		if isValid {
			break
		}

		tries--
	}

	if tries == 0 {
		return password, fmt.Errorf("retries exausted, could not generate a password with the set requirements")
	}

	return password, nil
}

func isValid(password string, length int, hasSymbols bool) bool {
	if len(password) != length {
		return false
	}

	if !strings.ContainsAny(password, Digits) {
		return false
	}

	if hasSymbols {
		if !strings.ContainsAny(password, Symbols) {
			return false
		}
	}

	if !strings.ContainsAny(password, LowerLetters) {
		return false
	}

	if !strings.ContainsAny(password, UpperLetters) {
		return false
	}

	return true
}

func buildRandomPassword(length int, hasSymbols bool) string {
	chars := LowerLetters + UpperLetters + Digits

	if hasSymbols {
		chars += Symbols
	}

	password := ""
	for i := 0; i < length; i++ {
		password += string(chars[rand.Intn(len(chars))])
	}

	return password
}
