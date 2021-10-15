package password

import (
	"math"
	"math/rand"
	"time"
)

const (
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits       = "0123456789"
	Symbols      = "@&$#!%*._+-=()[]{}:?"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Generate(passLength int, hasSymbols bool) (string, error) {
	chars := LowerLetters + UpperLetters + Digits

	password := ""
	for i := 0; i < passLength; i++ {
		char := string(chars[rand.Intn(len(chars))])
		password += char
	}

	if hasSymbols {
		nofSymbols := math.Ceil(float64(passLength) / 4)

		for i := 0; i < int(nofSymbols); i++ {
			index := rand.Intn(len(password))
			randChar := string(Symbols[rand.Intn(len(Symbols))])
			password = password[:index] + randChar + password[index+1:]
		}
	}

	return password, nil
}
