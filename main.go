package main

import (
	"flag"
	"fmt"
	"strings"
)

var minWordsFlag = flag.Int("w", 4, "Minimum number of words")
var minLengthFlag = flag.Int("l", 20, "Minimum password length, including separator")
var separatorFlag = flag.String("s", "", "Word separator")
var capitalizeFlag = flag.Bool("c", false, "Capitalize first letter of each word")

type wordlistFlag []string

func (w *wordlistFlag) String() string {
	return strings.Join(*w, ", ")
}

func (w *wordlistFlag) Set(value string) error {
	*w = append(*w, value)
	return nil
}

var wordlistFlagValue wordlistFlag

func init() {
	flag.Var(&wordlistFlagValue, "f", "Path to wordlist file (one word per line). Can be specified multiple times.")
}

func main() {
	flag.Parse()
	password := generate(*minWordsFlag, *minLengthFlag, *separatorFlag, *capitalizeFlag)
	fmt.Println(password)
}
