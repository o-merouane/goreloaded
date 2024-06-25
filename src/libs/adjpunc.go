package libs

import (
	"regexp"
	"strings"
)

func AdjustPunctuation(text string) string {
	punctuations := []string{".", ",", "!", "?", ":", ";"}
	for _, p := range punctuations {
		re := regexp.MustCompile(`\s*` + regexp.QuoteMeta(p) + `\s*`)
		text = re.ReplaceAllString(text, p+" ")
	}

	// Handle groups of punctuation like ...

	text = strings.ReplaceAll(text, ". . .", "...")

	return text
}
