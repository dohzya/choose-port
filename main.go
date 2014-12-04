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

var DEFAULT_FROM = 2000
var DEFAULT_TO = 40000

func main() {
	flag.Parse()

	var from int
	var to int

	if len(flag.Args()) == 0 {
		from = DEFAULT_FROM
		to = DEFAULT_TO
	} else {
		from = toInt(flag.Arg(0))
		if len(flag.Args()) == 1 {
			to = from + 1000
		} else {
			to = toInt(flag.Arg(1))
		}
	}

	println(chooseport.ChoosePort(from, to))
}
