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
	length  int
	symbols bool
	verbose bool
	encoded bool
)

func main() {
	flag.IntVar(&length, "length", 16, "Password length. Must be between 8 and 1024")
	flag.IntVar(&length, "l", 16, "Password length. Must be between 8 and 1024 [shorthand]")
	flag.BoolVar(&symbols, "symbols", true, "Toggle symbols")
	flag.BoolVar(&symbols, "s", true, "Toggle symbols [shorthand]")
	flag.BoolVar(&verbose, "verbose", false, "Verbose mode")
	flag.BoolVar(&verbose, "v", false, "Verbose mode [shorthand]")
	flag.BoolVar(&encoded, "base64", false, "Base64 encoded")
	flag.BoolVar(&encoded, "b", false, "Base64 encoded [shorthand]")

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

	generator := NewGenerator(length, symbols, encoded)

	password, err := generator.Generate()
	if err != nil {
		fmt.Println("Error generating the password with the given config", err)
		os.Exit(1)
	}

	if verbose {
		generator.PrintConfig()
	}

	fmt.Println(password)
	if encoded {
		fmt.Printf("\n%s\n", generator.Base64Encode(password))
	}
}
