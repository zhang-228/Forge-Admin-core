package xstrings

import "unicode"

// ToUpperFirst 首字母转大写
func ToUpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
}

// ToLowerFirst 首字母转小写
func ToLowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
