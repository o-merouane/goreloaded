package libs

import (
	"strings"
)

func FAdjustPunctuation(text string) string {
	text = strings.Replace(text, "cap, ", "cap,", -1)
	text = strings.Replace(text, "low, ", "low,", -1)
	text = strings.Replace(text, "up, ", "up,", -1)

	return text
}
