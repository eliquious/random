# random
A command line tool for generating random data

## Installation

```
go install github.com/eliquious/random
```

### Bash Completions

Run this command to generate a completion script. Copy the output into `.bash_profile`, etc.

```
random --completion-script-bash
```

### Zsh Completions

Run this command to generate a completion script. Copy the output into `.bash_profile`, etc.

```
random --completion-script-zsh
```

## Usage

```
usage: random [<flags>] <command> [<args> ...]

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  silly
    Generate a random silly name.

  first <GENDER>
    Generate a random first name.

  last
    Generate a random last name.

  full <GENDER>
    Generate a random full name.

  email
    Generate a random email.

  country [<flags>]
    Generate a random country.

  currency
    Generates a random currency.

  city
    Generates a random city.

  state [<flags>]
    Generates a random state.

  street
    Generates a random street.

  address
    Generates a random address.

  number <START> <STOP>
    Generates a random number.

  decimal <START> <STOP> <PREC>
    Generates a random decimal.

  bool
    Generates a random bool.

  paragraph
    Generates a random paragraph.

  postal <COUNTRY>
    Generates a random postal.

  ip
    Generates a random ipv4.

  day
    Generates a random day.

  hex [<LENGTH>]
    Generates a random hex.

  base64 [<LENGTH>]
    Generates a random base64.
```
