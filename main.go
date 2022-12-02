package main

import (
	"flag"
	"fmt"

	"github.com/library-development/go-schemacafe"
)

const usage = "Usage: schemac [go|ts] [dir]"

func main() {
	flag.Parse()
	lang := flag.Arg(0)
	from := schemacafe.Path{}
	to := flag.Arg(1)
	if to == "" {
		to = "."
	}
	switch lang {
	case "go":
		err := schemacafe.WriteGo(from, to)
		if err != nil {
			panic(err)
		}
	case "ts":
		err := schemacafe.WriteTypescript(from, to)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println(usage)
	}
}
