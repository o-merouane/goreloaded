package libs

func ToLower(s string) string {
	str := []rune(s)
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			str[i] = str[i] + 32
		}
	}
	return string(str)
}