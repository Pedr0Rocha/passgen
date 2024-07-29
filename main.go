package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	MAX_LENGTH = 1024
	MIN_LENGTH = 8
)

var (
	length     int
	hasSymbols bool
)

func main() {
	flag.IntVar(&length, "length", 16, "Password length. Must be between 8 and 1024")
	flag.IntVar(&length, "l", 16, "Password length. Must be between 8 and 1024 [shorthand]")
	flag.BoolVar(&hasSymbols, "symbols", true, "Toggle symbols")
	flag.BoolVar(&hasSymbols, "s", true, "Toggle symbols [shorthand]")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if length > MAX_LENGTH || length < MIN_LENGTH {
		fmt.Fprintf(os.Stderr, "Error: password length must be between %d and %d\n", MIN_LENGTH, MAX_LENGTH)
		flag.Usage()
		os.Exit(1)
	}

	generator := Generator{
		Length:     length,
		HasSymbols: hasSymbols,
		Attempts:   10_000,
	}

	password, err := generator.Generate()
	if err != nil {
		fmt.Println("Error generating the password with the given config", err)
		os.Exit(1)
	}

	fmt.Printf("%d character password generated:\n", generator.Length)
	fmt.Println(password)
}
