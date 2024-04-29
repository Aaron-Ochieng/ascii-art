package ascii
import "fmt"
func ErrHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}