package password

import (
	"strings"
	"testing"
)

// TestGeneratePasswordLength calls password.Generate and checks for password length.
func TestGeneratePasswordLength(t *testing.T) {
	want := 15
	password, err := Generate(15, false)

	if len(password) != want || err != nil {
		t.Errorf(`Generate(15, false) password length is %d, %v, want %d`, len(password), err, want)
	}
}

// TestGeneratePasswordHasSymbol calls password.Generate and checks if generated password has symbols.
func TestGeneratePasswordHasSymbols(t *testing.T) {
	password, err := Generate(30, true)

	containsSymbols := strings.ContainsAny(password, Symbols)

	if !containsSymbols || err != nil {
		t.Errorf(`Generate(30, true) password has no symbols, %v`, err)
	}
}

// TestGeneratePasswordHasSymbol calls password.Generate and checks if generated password has symbols.
func TestGeneratePasswordHasNoSymbols(t *testing.T) {
	password, err := Generate(30, false)

	containsSymbols := strings.ContainsAny(password, Symbols)

	if containsSymbols || err != nil {
		t.Errorf(`Generate(30, false) password has symbols, %v`, err)
	}
}
