# Secure memorable password generation CLI

## Overview
CLI tool to easily and securely generate long, safe, easy-to-remember passwords.

Inspired by [XKCD 936](https://xkcd.com/936/) and [John VDL's "Correct Horse Battery Staple"](https://bitbucket.org/jvdl/correcthorsebatterystaple)


## Installation
Head over to the Releases page and download the latest release binary for your platform.

## Basic usage

Simply invoke the binary:
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

You may pass a number of command-line parameters to the program to control 
the password generation logic and desired output.
---

    -w INTEGER
Minimum number of words to include in generated password. Default `4`.

---

    -l INTEGER
Minimum total length, including separator if defined, for the resulting password. Default `20`

---

    -s STRING
Separator to insert between individual words. Will be taken as-is. Default `-`

---

    -c

Switch to tell `horse` whether to capitalize each word in the resulting password. Default `true`

