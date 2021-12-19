package stringsmanipulation

import (
	"strconv"
	"strings"
)

func DecodeString(s string) string {
	pos := 0
	return decodeStringRec(s, &pos)
}

func decodeStringRec(s string, pos *int) string {
	var sb strings.Builder
	for *pos < len(s) {
		if isLetter(s[*pos]) {
			sb.WriteByte(s[*pos])
			*pos++
			continue
		}
		if s[*pos] == ']' {
			return sb.String()
		}
		end, k := extractNum(s, *pos)
		*pos = end + 1
		decoded := decodeStringRec(s, pos)
		writeStringKTimes(sb, decoded, k)
		*pos++
	}
	return sb.String()
}

func writeStringKTimes(sb strings.Builder, s string, k int) {
	for k > 0 {
		sb.WriteString(s)
		k--
	}
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z'
}

func extractNum(s string, pos int) (num, end int) {
	end = pos + 1
	for isDigit(s[end]) {
		end++
	}
	num, _ = strconv.Atoi(s[pos:end])
	return
}
