package ascii

import "fmt"

// ErrHandler function is used in handiling errors from the main programme
// printing the errors accordigly.
func ErrHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
