package feature

//var ValidParenthesesAnsiNums = []int{40, 41, 123, 125, 91, 93}
//var ValidParenthesesPairs = map[byte]byte{40: 41, 123: 125, 91: 93}

func ValidParentheses(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	} else if verifySymmetricalSlice(s) {
		return true
	} else {
		//givenParenthesesPairs := length / 2
		for i := 0; i < length; i += 2 {
			if validParenthesesAnsiNums(s[i], s[i+1]) {
				continue
			} else {
				return false
			}
		}
		return true
	}

}
func validParenthesesAnsiNums(firstBracketAnsiNum byte, lastBracketAnsiNum byte) bool {
	if firstBracketAnsiNum == 40 && lastBracketAnsiNum == 41 {
		return true
	} else if firstBracketAnsiNum == 91 && lastBracketAnsiNum == 93 {
		return true
	} else if firstBracketAnsiNum == 123 && lastBracketAnsiNum == 125 {
		return true
	} else {
		return false
	}

}

func verifySymmetricalSlice(s string) bool {
	middle_end := len(s) / 2
	for i := 0; i < middle_end; i++ {
		if !validParenthesesAnsiNums(s[i], s[len(s)-i-1]) {
			return false
		}
	}
	return true
}

func ifEachSegmentsValid(s string) bool {
	var segment_part int = 1
	for i := 0; i < len(s); i++ {
		if s[i] == 40 || s[i] == 91 || s[i] == 123 {
			segment_part++
		}
	}
	return true
}
