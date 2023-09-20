package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func randomFormat() string {
	formats := []string{
		"Hello, %s!",
		"Hi, %s!",
		"Hey, %s!",
		"Howdy, %s!",
		"Yo, %s!",
	}
	return formats[rand.Intn(len(formats))]
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		return "", errors.New("name undefined")
	}

	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat())
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        messages[name] = message
    }
    return messages, nil
}