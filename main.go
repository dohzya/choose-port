package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/dohzya/choose-port/chooseport"
)

func die(msg string) {
	fmt.Fprintf(os.Stderr, "%s\n", msg)
	os.Exit(1)
}

func toInt(str string) int {
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		die(fmt.Sprintf("Bad port limit: %s", str))
	}
	return int(num)
}

func main() {
	flag.Parse()

	var port int

	if len(flag.Args()) == 0 {
		port = chooseport.Random()
	} else if len(flag.Args()) == 1 {
		from := toInt(flag.Arg(0))
		port = chooseport.From(from)
	} else {
		from := toInt(flag.Arg(0))
		to := toInt(flag.Arg(1))
		port = chooseport.Between(from, to)
	}

	println(port)
}
