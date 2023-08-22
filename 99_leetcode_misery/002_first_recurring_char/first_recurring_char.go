package firstrecurringchar

func FirstRecurringChar(input string) (string, bool) {
	if len(input) == 0 {
		return "", false
	}

	set := make(map[rune]bool)

	for _, char := range input {
		if set[char] {
			return string(char), true
		} else {
			set[char] = true
		}
	}

	return "", false
}
