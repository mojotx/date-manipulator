package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// Set local time to UTC
	if utc, err := time.LoadLocation("UTC"); err != nil {
		panic("Cannot set time.Local to UTC: " + err.Error())
	} else {
		time.Local = utc
	}

	// If no arguments, just print now
	if len(os.Args) == 1 {
		fmt.Println(time.Now().Format(time.RFC3339))
		return
	}

	// process each duration
	for _, arg := range os.Args[1:] {
		offset, err := time.ParseDuration(arg)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error parsing time offset: %s\n", err.Error())
			continue
		}

		fmt.Println(time.Now().Add(offset).Format(time.RFC3339))
	}
}
