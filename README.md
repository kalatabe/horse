# Secure memorable password generation program

## Overview

`horse` is a command-line tool to easily and securely generate long, safe, easy-to-remember passwords.

Inspired by [XKCD 936](https://xkcd.com/936/)
and [John VDL's "Correct Horse Battery Staple"](https://bitbucket.org/jvdl/correcthorsebatterystaple)

## Download

Head over to the Releases page and download the latest release binary for your platform.

## Building from source

1. Install `go` for your platform/OS, then make sure it's in your `PATH`
2. Clone the project
3. Run `go build`. The executable `horse` will be created in the project's root directory.

## Basic usage

Simply run the program:

```bash
$ ./horse 
HighlySellersWiredDominant
```

For quick help, pass the `--help` parameter:

```bash
$ ./horse --help
Usage of ./horse:
  -c	Capitalize first letter of each word (default true)
  -l int
    	Minimum password length (default 20)
  -s string
    	Word separator
  -w int
    	Minimum number of words (default 4)

```

## Parameters

You may pass certain command-line parameters to the program to control
the password generation process and desired output.
---

    -w INTEGER

Minimum number of words to include in generated password. Default `4`.

---

    -s STRING

Separator to insert between individual words. Will be taken as-is. Default is blank.

---

    -l INTEGER

Minimum total length of password, including separator if defined. Default `20`

---
    -c

Capitalize each word in the resulting password. Default `true`

