package main

import (
	"flag"
	"fmt"

	"github.com/library-development/go-schemacafe"
)

const usage = "Usage: schemac [py|rb|java|cs|ts|go|rs] <dir>"

func main() {
	flag.Parse()
	lang := flag.Arg(0)
	from := schemacafe.Path{}
	to := flag.Arg(1)
	if to == "" {
		to = "."
	}
	var err error
	switch lang {
	case "py":
		err = schemacafe.WritePython(from, to)
	case "rb":
		err = schemacafe.WriteRuby(from, to)
	case "java":
		err = schemacafe.WriteJava(from, to)
	case "cs":
		err = schemacafe.WriteCSharp(from, to)
	case "ts":
		err = schemacafe.WriteTypescript(from, to)
	case "go":
		err = schemacafe.WriteGo(from, to)
	case "rs":
		err = schemacafe.WriteRust(from, to)
	default:
		fmt.Println(usage)
	}
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
