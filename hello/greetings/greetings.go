package greetings

import (
	"errors"
	"fmt"

	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Name is required")
	}

	message := fmt.Sprintf(randomname(), name)
	return message, nil
}

func randomname() string {
	format := []string{
		"Hi %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! well met.",
	}

	return format[rand.Intn(len(format))]
}
