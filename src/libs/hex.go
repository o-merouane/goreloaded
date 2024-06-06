package libs

import (
	"strconv"
)

func Hex(s string) string {
	d, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		panic(err)
	}
	t := strconv.FormatInt(d, 10)

	return t
}
