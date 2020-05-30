package piglatin

import (
	"strings"
)

const (
	suffix  string = "ay"
	vowels  string = "aeiou"
	ySuffix string = "y" + suffix
)

func Translate(text string) string {
	if len(strings.TrimSpace(text)) == 0 {
		return text
	}

	words := strings.Split(text, " ")

	var pigLatinWords []string

	for _, word := range words {
		if strings.Contains(vowels, word[0:1]) {
			pigLatinWords = append(pigLatinWords, word+ySuffix)
		} else {
			var lastConsonantIndex int
			for i := 1; i < len(word[:]); i++ {
				if strings.Contains(vowels, word[i:i+1]) {
					break
				} else {
					lastConsonantIndex = i
				}
			}

			pigLatinWords = append(pigLatinWords, word[lastConsonantIndex+1:]+word[:lastConsonantIndex+1]+suffix)
		}
	}

	return strings.Join(pigLatinWords, " ")
}
