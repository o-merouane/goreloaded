package libs

import (
	"regexp"
	"strings"
)

func AdjustSingleQuotes(text string) string {
    // This regular expression will match single quotes around one or more words, allowing for spaces around them
    re := regexp.MustCompile(`'\s*([^']+?)\s*'`)
    text = re.ReplaceAllStringFunc(text, func(match string) string {
        // Remove spaces after the opening quote and before the closing quote
        return "'" + strings.TrimSpace(match[1:len(match)-1]) + "'"
    })
    return text
}