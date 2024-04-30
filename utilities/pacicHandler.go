package ascii

import "log"

func HandlePanic() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
}
