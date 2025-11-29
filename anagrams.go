package main

func IsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	count := make(map[rune]int)

	for _, c := range a {
		count[c]++
	}

	for _, c := range b {
		count[c]--
		if count[c] < 0 {
			return false
		}
	}

	return true
}
