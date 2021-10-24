package genparanthesis

func generateParenthesis(n int) []string {
	list := make([]string, 0)
	comb := make([]byte, 2*n)
	generateParenthesisRec(0, n, n, 0, comb, &list)
	return list
}

func generateParenthesisRec(balance, leftNum, rightNum, index int, comb []byte, list *[]string) {
	if balance < 0 || leftNum < 0 || rightNum < 0 {
		return
	}
	if balance == 0 && leftNum == 0 && rightNum == 0 {
		*list = append(*list, string(comb))
		return
	}
	comb[index] = '('
	generateParenthesisRec(balance+1, leftNum-1, rightNum, index+1, comb, list)
	comb[index] = ')'
	generateParenthesisRec(balance-1, leftNum, rightNum-1, index+1, comb, list)
}
