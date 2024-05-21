package main

import (
	"flag"
	"fmt"
)

var minWordsFlag = flag.Int("w", 4, "Minimum number of words")
var minLengthFlag = flag.Int("l", 20, "Minimum password length, including separator")
var separatorFlag = flag.String("s", "", "Word separator")
var capitalizeFlag = flag.Bool("c", false, "Capitalize first letter of each word")

func main() {
	flag.Parse()
	password := generate(*minWordsFlag, *minLengthFlag, *separatorFlag, *capitalizeFlag)
	fmt.Println(password)
}
