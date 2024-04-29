package ascii

import "strings"

// HandleBackspace function is used in handldling the backspace escape charcter
// in the input string.
func HandleBackspace(s string) string {
	index := strings.Index(s, "\\b") // obtaining the index of the first occurance of \b character

	if index == -1 { // Base case: no "\b" found, return the string as is
		return s
	}
	if index == 0 { // If "\b" is at the beginning, remove it and return the rest of the string
		return HandleBackspace(s[1:])
	}
	return HandleBackspace(s[:index-1] + s[index+2:]) // Remove the "\b" and the character before it recursively
}
