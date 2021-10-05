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

	if len(os.Args) == 1 {
		fmt.Println(time.Now().Format(time.RFC3339))
	} else {
		for _, arg := range os.Args[1:] {
			if offset, err := time.ParseDuration(arg); err != nil {
				fmt.Printf("error parsing timeOffset: %s\n\n", err.Error())
				fmt.Printf("usage: %s <timeOffSet>\n", os.Args[0])
				os.Exit(1)
			} else {
				fmt.Println(time.Now().Add(offset).Format(time.RFC3339))
			}
		}
	}
}
