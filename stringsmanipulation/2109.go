package stringsmanipulation

func addSpaces(s string, spaces []int) string {
	res := make([]byte, len(s)+len(spaces))
	posS := 0
	posSpaces := 0
	for i := range res {
		if posSpaces < len(spaces) && posS == spaces[posSpaces] {
			res[i] = ' '
			posSpaces++
		} else {
			res[i] = s[posS]
			posS++
		}
	}
	return string(res)
}
