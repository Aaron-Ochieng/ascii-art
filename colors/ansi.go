package utils

import "fmt"

func GetColor(colorName string) string {
	res := GetRGBValue(colorName)

	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm ", res.R, res.G, res.B)
}