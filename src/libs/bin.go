package libs

import (
	"strconv"
)

func Bin(s string) string {
	d, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	t := strconv.FormatInt(d, 10)
	return t
}
