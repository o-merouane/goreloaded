package libs

import (
	"strconv"
	"strings"
)

func FAdjustPunctuation(text string) string {

	for i := 0; i < 20; i++ {
		x := strconv.Itoa(i)
		text = strings.Replace(text, x+"),", x+") ,", -1)
	}

	text = strings.Replace(text, "cap, ", "cap,", -1)
	text = strings.Replace(text, "low, ", "low,", -1)
	text = strings.Replace(text, "up, ", "up,", -1)

	return text
}
