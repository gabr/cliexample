package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Available options:\n")
		flag.PrintDefaults()
	}

	number := flag.Int("number", -1, "natural number")
	//fmt.Println("Parsed(): ", flag.Parsed()) // false
	flag.Parse()
	//fmt.Println("Parsed(): ", flag.Parsed()) // true

	fmt.Println("Given number has value:", *number)

	/*
		// this prints args after the '--' or any non '-' arg
		go run .\cliexample.go -number 8 -- test
		Given number has value: 8
		Args():  [test]
	*/
	fmt.Println("Args(): ", flag.Args())

	/*
		There is no way to specify mandatory arguments,
		but if some arguments were not set then this
		Visit() function will not list them.
	*/
	numberWasSet := false
	flag.Visit(func (f *flag.Flag) {
		fmt.Println("Visit: ", f)
		if f.Name == "number" {
			numberWasSet = true
		}
	})

	// number parameter is required
	if numberWasSet == false {
		flag.Usage()
		return
	}

	flag.VisitAll(func (f *flag.Flag) {
		fmt.Println("VisitAll: ", f)
	})
}
