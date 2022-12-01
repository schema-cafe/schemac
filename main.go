package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/library-development/go-schemacafe"
	"github.com/library-development/go-util"
)

func main() {
	flag.Parse()
	lang := flag.Arg(0)
	from := []string{}
	to := "."
	switch lang {
	case "go":
		WriteGo(from, to)
	case "ts":
		WriteTS(from, to)
	default:
		PrintUsage()
	}
}

func PrintUsage() {
	println("Usage: shemac [go|ts]")
}

func WriteGo(from []string, to string) error {
	path := path(from)

	fmt.Printf("Pulling from schema.cafe%s\n", path)

	client := schemacafe.NewClient()
	r := client.Get(path)

	if r.IsFolder {
		os.MkdirAll(to, 0755)
		for _, entry := range r.Folder.Contents {
			err := WriteGo(append(from, entry.Name), to+"/"+entry.Name)
			if err != nil {
				return err
			}
		}
	}

	if r.IsSchema {
		var buf bytes.Buffer
		buf.WriteString("package ")
		buf.WriteString(pkgName(to))
		buf.WriteString("\n")

		buf.WriteString("type ")
		buf.WriteString(name(to))
		buf.WriteString(" struct {\n")

		for _, field := range r.Schema.Fields {
			buf.WriteString("\t")
			buf.WriteString(field.Name)
			buf.WriteString(" ")
			buf.WriteString(field.Type)
			buf.WriteString(" `json:\"")
			buf.WriteString(field.Name)
			buf.WriteString("\"`\n")
		}

		buf.WriteString("}\n")

		err := os.WriteFile(to+".go", buf.Bytes(), os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeGoFile(path string) error {
	fmt.Printf("Writing %s\n", path)
	return nil
}

func WriteTS(from []string, to string) {
	fmt.Printf("Pulling from schema.cafe%s\n", from)
}

func path(p []string) string {
	return strings.Join(util.ArrayMap(p, func(s string) string {
		return "/" + s
	}), "")
}
