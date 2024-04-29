package ascii

import "errors"

// CheckIllegalChar function is used to handle the charcharters not in range
// of printable ascii characters.
func CheckIllegalChar(arr []string) ([]string, error) {
	for _, s := range arr {
		for _, r := range s {
			if r < 32 || r > 126 {
				return []string{}, errors.New("the string must contain characters in the range 32 - 126")
			}
		}
	}
	return arr, nil
}
