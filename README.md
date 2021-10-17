# CLI tool to generate simple passwords

Generates a simple password up to 64 characters, using uppercase and lowercase letters, digits and symbols (if enabled).

## Install

`go build -o $GOPATH/passgen`

## Usage

`passgen --help`

`passgen generate --help`

## Examples

Default flags, password length 15 and has symbols:

`passgen generate`

`Ouput: J3x7_FF]E4HYd3.`


Password length 30 and no symbols:

`passgen generate -l=30 -s=false`

`Ouput: 4bPkIAjshaoiJqowtY3qQqm5GNe6DD`



## Disclaimer

This password generator does **NOT** implement high entropy strategies and should not be used to extreme security requirement applications. It's highly unlikely but the tool does not garantee to generate a password with upper + lower + digit, increase the password length or generate another one.
