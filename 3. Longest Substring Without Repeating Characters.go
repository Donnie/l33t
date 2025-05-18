func lengthOfLongestSubstring(s string) int {
	charslice := strings.Split(s, "")
	largestLen := 0 // min length could be empty string

	seenchars := map[string]int{}
	start := 0

	for i := 0; i < len(charslice); i++ {
		char := charslice[i] // current char -b

		if pos, found := seenchars[char]; found && start <= pos {
			// if prev position is found
            // and the current start is before the previous position
            // i.e. the current substring start
            // is before the recurrence of the current char
            // i.e. the current substring has a recurrence
			// reset start of current string to char right after prev position
            // for e.g. in case of bb
            // when i is at 1, pos = 0, start = 0, start needs to be set to 1
			start = pos + 1
		}

        seenchars[char] = i // a = 0; b = 1; c = 2
		currentLen := i - start + 1
		if currentLen > largestLen {
			largestLen = currentLen
		}
	}
	return largestLen
}
