
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	lp := s[:1]
	for i := 0; i < len(s); i++ {
		left1, right1 := expandFromCenter(s, i, i)
		if right1+1-left1 > len(lp) {
			lp = s[left1 : right1+1]
		}

		left2, right2 := expandFromCenter(s, i, i+1)
		if right2+1-left2 > len(lp) {
			lp = s[left2 : right2+1]
		}
	}
	return lp
}

func expandFromCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
