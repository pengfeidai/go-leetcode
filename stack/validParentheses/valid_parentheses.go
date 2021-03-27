package stack

func isValid(s string) bool {
	// 空字符串直接返回true
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '{') || (v == '(') {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			// 栈顶元素
			top := stack[len(stack)-1]
			if top == '(' && v == ')' || top == '[' && v == ']' || top == '{' && v == '}' {
				// 匹配完后离开栈
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
