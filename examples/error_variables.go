package examples

import (
	"errors"
	"fmt"
	"math/rand"
)

const MAX_TIMEOUTS = 5

var ErrTimeout = errors.New("the request timed out")
var ErrRejected = errors.New("the request was rejected")

func sendRequest() (string, error) {
	switch rand.Intn(3) % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}

func ErrorVariables() {
	response, err := sendRequest()
	if errors.Is(err, ErrTimeout) {
		timeouts := 0
		for errors.Is(err, ErrTimeout) {
			timeouts++
			fmt.Println("Timeout. Retrying.")
			if timeouts == MAX_TIMEOUTS {
				panic("too many timeouts")
			}
			response, err = sendRequest()
		}
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
