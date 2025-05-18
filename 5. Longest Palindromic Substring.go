func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
	lens := len(s)

	lp := s[:1]
    lenlp := 1

	for i := 0; i < lens; i++ {
		for j := lens - 1; j >= i && (j+1-i) > lenlp; j-- {
			if s[i] == s[j] && isPal(s, i+1, j-1) {
				lp = s[i : j+1]
                lenlp = j+1-i
			}
		}
	}
	return lp
}

func isPal(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
        i++;
        j--;
	}
	return true
}
