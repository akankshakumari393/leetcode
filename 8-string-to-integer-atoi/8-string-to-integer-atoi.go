import "math"

func myAtoi(str string) int {
	s := 0
	for s = 0; s < len(str); s++ {
        // check if the string is valid
		if !isDigitOrSignOrWhitespace(str[s]) {
			return 0
		}

		// find first +/- or digit and store in s
		if isDigitOrSign(str[s]) {
			break
		}
	}

    positive := true
	if s < len(str) {
		if str[s] == '-' {
			positive = false
			s++
		} else if str[s] == '+' {
			s++
			positive = true
		} 
    }

	result := 0
	for ; s < len(str); s++ {
		if !isDigit(str[s]) {
			break
		}
        fmt.Println("print",str[s], int(str[s]), '0')
		d := int(str[s]) - '0'
		var nextResult int
		if !positive {
			d = -d
		}

		nextResult = result*10 + d

		// overflow
		if nextResult > math.MaxInt32 || nextResult < math.MinInt32 || (nextResult-d)/10 != result {
			if positive {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = nextResult
	}

	return result
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}

func isDigitOrSign(c byte) bool {
	if isDigit(c) || c == '-' || c == '+' {
		return true
	} else {
		return false
	}
}

func isDigitOrSignOrWhitespace(c byte) bool {
	if isDigitOrSign(c) || c == ' ' {
		return true
	} else {
		return false
	}
}
