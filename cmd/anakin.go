package anakin

import (
	"fmt"
	"rsc.io/quote"
)

func callName(name string) string {
	message := fmt.Sprintf("Hi, %s! %s", name, quote.Go())
	return message
}

func getNumber(number int64) int64 {
	return number
}
