package main

import (
	"os"
	"strings"
	"testing"
)

func TestReadWordlist(t *testing.T) {
	// Create a temporary test file
	content := "word1\nword2\nword3\n"
	tmpfile, err := os.CreateTemp("", "test-wordlist-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}

	words, err := readWordlist(tmpfile.Name())
	if err != nil {
		t.Fatalf("readWordlist failed: %v", err)
	}

	expectedWords := []string{"word1", "word2", "word3"}
	if len(words) != len(expectedWords) {
		t.Errorf("Expected %d words, got %d", len(expectedWords), len(words))
	}

	for i, word := range words {
		if word != expectedWords[i] {
			t.Errorf("Word %d: expected %q, got %q", i, expectedWords[i], word)
		}
	}
}

func TestGeneratePassword(t *testing.T) {
	testWords := []string{"short", "medium", "longerword"}

	tests := []struct {
		name       string
		minWords   int
		minLength  int
		separator  string
		capitalize bool
	}{
		{
			name:       "basic generation",
			minWords:   2,
			minLength:  10,
			separator:  "-",
			capitalize: false,
		},
		{
			name:       "capitalized words",
			minWords:   2,
			minLength:  10,
			separator:  "-",
			capitalize: true,
		},
		{
			name:       "no separator",
			minWords:   2,
			minLength:  10,
			separator:  "",
			capitalize: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			originalWordlistFlag := make(wordlistFlag, len(wordlistFlagValue))
			copy(originalWordlistFlag, wordlistFlagValue)

			tmpfile, err := os.CreateTemp("", "test-wordlist-*.txt")
			if err != nil {
				t.Fatalf("Failed to create temporary file: %v", err)
			}
			defer os.Remove(tmpfile.Name())
			wordlistFlagValue = wordlistFlag{tmpfile.Name()}
			defer func() { wordlistFlagValue = originalWordlistFlag }()

			for _, word := range testWords {
				if _, err := tmpfile.WriteString(word + "\n"); err != nil {
					t.Fatalf("Failed to write to temporary file: %v", err)
				}
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatalf("Failed to close temporary file: %v", err)
			}

			// Generate password
			password := generate(tt.minWords, tt.minLength, tt.separator, tt.capitalize)

			//minimum length
			if len(password) < tt.minLength {
				t.Errorf("Password length %d is less than minimum length %d", len(password), tt.minLength)
			}

			//minimum words
			words := strings.Split(password, tt.separator)
			if len(words) < tt.minWords {
				t.Errorf("Number of words %d is less than minimum words %d", len(words), tt.minWords)
			}

			//capitalization
			if tt.capitalize {
				for _, word := range words {
					if word != "" && !strings.HasPrefix(word, strings.ToUpper(string(word[0]))) {
						t.Errorf("Word %q is not capitalized", word)
					}
				}
			}
		})
	}
}
