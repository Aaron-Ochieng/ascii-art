package ascii
import "strings"
func StringContain(s string) []string {
	if strings.Contains(s, "o") {
		return strings.Split(s, "\r\n")
	}
	return strings.Split(s, "\n")
}