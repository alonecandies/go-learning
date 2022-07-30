package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string,error) {
    // Return a greeting that embeds the name in a message.
    if (name == "") {
	return "", errors.New("empty name")
    }
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages:= make(map[string]string)
	for _, name := range names {
		message, err :=Hello(name)
		if err!=nil {
			return nil,err
		}
		messages[name] = message
	}
	return messages, nil
}

//init sets intial values for variables used in the func
func init() {
	rand.Seed(time.Now().UnixNano())
}

//randomFormat returns one of a set of greeting messages. The returned message is selected at random.
func randomFormat() string {
	//A slice of msg formats.
	formats:=[]string {
		"Hi, %v. Welcome!",
		"Great to see u, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
