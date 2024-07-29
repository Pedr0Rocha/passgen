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
	Symbols      = "@&$#!%*._+-=()[]{}:?"
)

type Generator struct {
	Length     int
	HasSymbols bool
	Attempts   int
}

func (g Generator) Generate() (string, error) {
	tries := g.Attempts

	password := ""
	for tries > 0 {
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

func (g Generator) isValid(password string) bool {
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

func (g Generator) buildRandomPassword() string {
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