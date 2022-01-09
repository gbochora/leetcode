package stringsmanipulation

import "strings"

func checkString(s string) bool {
	lastA := strings.LastIndex(s, "a")
	firstB := strings.Index(s, "b")
	return lastA < firstB || firstB == -1
}
