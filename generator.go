package main

import (
	"crypto/rand"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"math/big"
	"strings"
)

func getRandomInt(max int) int64 {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max+1)))
	if err != nil {
		log.Fatalf("Failed to initialize secure random number generator! %s", err)
	}
	return num.Int64()
}

func getWord(WordList *[]string, capitalize bool) string {
	vWords := *WordList
	index := getRandomInt(len(vWords))
	newWord := vWords[index]
	if capitalize {
		newWord = cases.Title(language.English, cases.Compact).String(newWord)
	}
	return newWord
}

func generate(minWords int, minLength int, separator string, capitalize bool) string {
	var words []string
	dictionary := &FullDictionary
	for i := 1; i <= minWords; i++ {
		words = append(words, getWord(dictionary, capitalize))
	}

	password := strings.Join(words, separator)

	// If password still doesn't meet the minimum length requirement, keep adding words until it does
	for len(password) < minLength {
		words = append(words, getWord(dictionary, capitalize))
		password = strings.Join(words, separator)
	}
	return password
}
