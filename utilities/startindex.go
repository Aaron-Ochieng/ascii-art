package ascii

func GetStartingIndex(ascii int) int {
	return (ascii-32)*9 + 2
}
