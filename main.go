package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	err := setTZtoUTC()
	if err != nil {
		panic("could not set time.Local to UTC")
	}

	// If a command-line argument was passed,
	// convert to a time.Duration value
	var offset time.Duration
	if len(os.Args) == 2 {
		offset, err = time.ParseDuration(os.Args[1])
		if err != nil {
			fmt.Printf("error: %s\n\n", err.Error())
			fmt.Printf("usage: %s [ timeOffset ]\n", os.Args[0])
			os.Exit(1)
		}
	}
	fmt.Println(time.Now().Add(offset).Format(time.RFC3339))
}

// setTZtoUTC sets the time.Local value to UTC
func setTZtoUTC() error {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		return err
	}

	time.Local = utc
	return nil
}
