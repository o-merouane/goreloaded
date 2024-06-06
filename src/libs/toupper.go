package libs

func ToUpper(s string) string {
	str := []rune(s)
	for i := 0; i < len(s); i++ {
		if s[i] >= 97 && s[i] <= 122 {
			str[i] = str[i] - 32
		}
	}
	return string(str)
}
