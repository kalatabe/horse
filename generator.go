package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func getRandomInt(max int) int64 {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
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

func readWordlist(filepath string) ([]string, error) {
	if filepath == "" {
		return nil, fmt.Errorf("wordlist file path is empty")
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open wordlist file: %v", err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading wordlist file: %v", err)
	}

	if len(words) == 0 {
		return nil, fmt.Errorf("wordlist file contains no valid words. Words must not consist entirely of whitespace")
	}

	return words, nil
}

func generateDictionary() *[]string {
	var dictionary *[]string

	if len(wordlistFlagValue) == 0 {
		dictionary = &FullDictionary
		return dictionary
	}

	var allWords []string
	for _, filepath := range wordlistFlagValue {
		customWords, err := readWordlist(filepath)
		if err != nil {
			log.Fatalf("Error reading wordlist %s: %v", filepath, err)
		}
		allWords = append(allWords, customWords...)
	}
	dictionary = &allWords
	return dictionary
}

func generate(minWords int, minLength int, separator string, capitalize bool) string {
	var words []string
	dictionary := generateDictionary()

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
