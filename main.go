package main

import "flag"

func main() {
	flag.Parse()
	lang := flag.Arg(0)
	path := flag.Arg(1)
	switch lang {
	case "go":
		WriteGo(path)
	case "ts":
		WriteTS(path)
	default:
		PrintUsage()
	}
}

func PrintUsage() {
	println("Usage: shemac [go|ts] [path]")
}

func WriteGo(path string) {
	println("Writing Go code to", path)
}

func WriteTS(path string) {
	println("Writing TypeScript code to", path)
}
