package examples

import (
	"errors"
	"fmt"
	"log"
)

var errTimeout = errors.New("the request timed out")

func sendRequest2() (string, error) {
	return "", fmt.Errorf("we got an error: %w ", errTimeout) // %w is for unwrapping the error
}

func WrappingErrors() {
	if _, err := sendRequest2(); err != nil {
		if errors.Is(err, errTimeout) {
			log.Println("we got a timeout error")
		} else {
			log.Println("we got some other error")
		}
	}
}
