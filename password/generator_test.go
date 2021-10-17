package password

import (
	"strings"
	"testing"
)

// TestGeneratePasswordLength calls password.Generate and checks for password length.
func TestGeneratePasswordLength(t *testing.T) {
	want := 15
	pass, err := Generate(15, false)

	if len(pass) != want || err != nil {
		t.Errorf(`Generate(15, false) password length is %d, %v, want %d`, len(pass), err, want)
	}
}

// TestGeneratePasswordHasSymbol calls password.Generate and checks if generated password has symbols.
func TestGeneratePasswordHasSymbol(t *testing.T) {
	pass, err := Generate(15, true)

	containSymbol := strings.ContainsAny(pass, Symbols)

	if !containSymbol || err != nil {
		t.Errorf(`Generate(15, true) password has no symbols, %v`, err)
	}
}
