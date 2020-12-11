package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	digitRegexp, err := regexp.Compile("[0-9]+")

	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	content = digitRegexp.Find(content)

	// Copy the matching data to a new slice.
	digits := make([]byte, len(content))
	copy(digits, content)

	text := string(digits)

	fmt.Println(text)
}
