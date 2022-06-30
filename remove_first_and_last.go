package main

// RemoveChar returns string with first and last letters removed
func RemoveChar(word string) string {
	switch len(word) {
	case 0, 1, 2:
		return ""
	default:
		return word[1 : len(word)-1]
	}
}
