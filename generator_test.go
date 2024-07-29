package main

import (
	"strings"
	"testing"
)

func TestGeneratePasswordLength(t *testing.T) {
	gen := Generator{
		Length:     15,
		HasSymbols: false,
		Attempts:   10_000,
	}
	password, err := gen.Generate()
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	if len(password) != gen.Length {
		t.Errorf("expected %d, got %d", gen.Length, len(password))
	}
}

func TestGeneratePasswordHasSymbols(t *testing.T) {
	gen := Generator{
		Length:     30,
		HasSymbols: true,
		Attempts:   10_000,
	}
	password, err := gen.Generate()
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	containsSymbols := strings.ContainsAny(password, Symbols)

	if !containsSymbols {
		t.Errorf("expected symbols, got %s", password)
	}
}

func TestGeneratePasswordHasNoSymbols(t *testing.T) {
	gen := Generator{
		Length:     30,
		HasSymbols: false,
		Attempts:   10_000,
	}
	password, err := gen.Generate()
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	containsSymbols := strings.ContainsAny(password, Symbols)

	if containsSymbols {
		t.Errorf("expected no symbols, got %s", password)
	}
}
