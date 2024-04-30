package ascii

// GetStartingIndex is a function used to get the starting postion
// using the ascii value of the charcter passed as argument
// in the ascii art fil.
func GetStartingIndex(ascii int) int {
	return (ascii-32)*9 + 1
}
