package main

import (
	"strings"
	"testing"
)

func TestGenerator(t *testing.T) {
	inputs := []Generator{
		NewGenerator(1024, true),
		NewGenerator(15, false),
		NewGenerator(30, true),
		NewGenerator(30, false),
	}

	for _, gen := range inputs {
		t.Run("Test generator", func(t *testing.T) {
			password, err := gen.Generate()
			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}

			containsSymbols := strings.ContainsAny(password, Symbols)
			containsLowerCase := strings.ContainsAny(password, LowerLetters)
			containsUpperCase := strings.ContainsAny(password, UpperLetters)
			containsDigits := strings.ContainsAny(password, Digits)

			if gen.Length != len(password) {
				t.Errorf("wrong length. expected %d, got %d", gen.Length, len(password))
			}

			if gen.HasSymbols != containsSymbols {
				t.Errorf("wrong password composition (symbols). expected %v, got %v", gen.HasSymbols, containsSymbols)
			}

			if !containsLowerCase {
				t.Errorf("wrong password composition (lowercase letters). expected %v, got %v", true, containsLowerCase)
			}

			if !containsUpperCase {
				t.Errorf("wrong password composition (uppercase letters). expected %v, got %v", true, containsUpperCase)
			}

			if !containsDigits {
				t.Errorf("wrong password composition (digits). expected %v, got %v", true, containsDigits)
			}
		})
	}
}
