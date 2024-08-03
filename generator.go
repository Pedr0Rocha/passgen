package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits       = "0123456789"
	Symbols      = "@&$#!%*._+-=()[]{}:?<>"
)

type Generator struct {
	Length        int
	HasSymbols    bool
	Attempts      int
	AttemptsTaken int
}

func NewGenerator(length int, symbols bool) *Generator {
	return &Generator{
		Length:        length,
		HasSymbols:    symbols,
		Attempts:      10_000,
		AttemptsTaken: 0,
	}
}

func (g *Generator) Generate() (string, error) {
	tries := g.Attempts

	password := ""
	for tries > 0 {
		g.AttemptsTaken++

		password = g.buildRandomPassword()
		isValid := g.isValid(password)

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

// A password is considered valid if its length is correct and it contains at least one of each:
// upper case letter, lower case letter, digit and symbol (if enabled).
func (g *Generator) isValid(password string) bool {
	if len(password) != g.Length {
		return false
	}

	if !strings.ContainsAny(password, Digits) {
		return false
	}

	if g.HasSymbols {
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

func (g *Generator) PrintConfig() {
	fmt.Printf("Configuration:\n")
	fmt.Printf("- Length: %d\n", g.Length)
	fmt.Printf("- Symbols: %v\n", g.HasSymbols)
	fmt.Printf("- Attempts: %d\n\n", g.AttemptsTaken)
}

func (g *Generator) buildRandomPassword() string {
	chars := LowerLetters + UpperLetters + Digits

	if g.HasSymbols {
		chars += Symbols
	}

	password := ""
	for i := 0; i < g.Length; i++ {
		password += string(chars[rand.Intn(len(chars))])
	}

	return password
}
