package stringsmanipulation

import (
	"strings"
	"unicode"
)

func capitalizeTitle(title string) string {
	title = strings.ToLower(title)
	chars := []byte(title)
	start := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] == ' ' {
			if i-start > 2 {
				chars[start] = byte(unicode.ToUpper(rune(chars[start])))
			}
			start = i + 1
		}
	}
	if len(chars)-start > 2 {
		chars[start] = byte(unicode.ToUpper(rune(chars[start])))
	}

	return string(chars)
}
