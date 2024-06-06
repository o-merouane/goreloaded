package libs

import (
	"strings"
)

func AdjustIndefiniteArticles(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			nextWord := words[i+1]
			if strings.ContainsAny(string(nextWord[0]), "aeiouAEIOU") || strings.ToLower(string(nextWord[0])) == "h" {
				words[i] = "an"
			}
		}
	}
	return strings.Join(words, " ")
}