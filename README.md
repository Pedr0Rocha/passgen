# Passgen - Password Generator CLI

Generates a password up to 1024 characters. It guarantees at least one uppercase letter,
lowercase letter, a digit and a symbol (if enabled).

## Install

```bash
go install github.com/pedr0rocha/passgen@latest
```

## Usage

```bash
passgen --help
-length  [-l] (default 16)
-symbols [-s] (default true)
-verbose [-v] (default false)
-base64  [-b] (default false)
```

```bash
passgen -length=64
# XzsOX_]Y3+O_a]}&#bJia:Qji6nWhwH==[DP:7nEw*eF#c:Joq1_Hx-bVlToM7JY

passgen -l=24 -symbols=false
# B5m9MVbWNaULj37UJgBqt0vU

passgen
# W8VrEF2h#!QWFdrP

passgen -v
# Configuration:
# - Length: 16
# - Symbols: true
# - Attempts: 1
#
# Ok0E3L?7]TSe=-Ws

passgen -b -v
# Configuration:
# - Length: 16
# - Symbols: true
# - Encoded: true
# - Attempts: 1
#
# K8ld3S.ngJc+E4[n
#
# SzhsZDNTLm5nSmMrRTRbbg==

passgen -base64
# GYPp5F?Zh=+7?rvT
#
# R1lQcDVGP1poPSs3P3J2VA==
```

## Tests

```bash
make test
```

## Disclaimer

This password generator does **NOT** implement high entropy strategies and should not be used
to extreme security requirement applications.
