type ByteStack struct {
	values []byte
}

func (s *ByteStack) push(b byte) {
	s.values = append(s.values, b)
}

func (s *ByteStack) pop() byte {
	if len(s.values) == 0 {
		return 0
	}
	ret := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return ret
}

func isValid(s string) bool {
    stack := &ByteStack{}
    for i :=0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            stack.push(s[i])
            continue            
        }
		switch s[i] {
            case '(', '{', '[': stack.push(s[i])
		case ')': if stack.pop() != '(' {return false}
		case '}': if stack.pop() != '{' {return false}
		case ']': if stack.pop() != '[' {return false}
		}
    }
    return stack.pop() == 0
}
