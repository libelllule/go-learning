package validparentheses

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]byte, 0, len(s)/2)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if s[i] == ')' && last != '(' ||
				s[i] == ']' && last != '[' ||
				s[i] == '}' && last != '{' {
				return false
			}
		}
	}
	return len(stack) == 0
}
