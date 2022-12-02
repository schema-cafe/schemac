package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/library-development/go-nameconv"
	"github.com/library-development/go-schemacafe"
)

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
		err := WriteGo(from, to)
		if err != nil {
			panic(err)
		}
	default:
		PrintUsage()
	}
}

func PrintUsage() {
	fmt.Println("Usage: schemac [go] [dir]")
}

func WriteGo(from schemacafe.Path, to string) error {
	client := schemacafe.NewClient()
	fmt.Printf("Pulling from %s%s...", client.APIURL, from.String())
	r := client.Get(from)

	if r.IsFolder {
		fmt.Println(" folder found")
		fmt.Printf("Making sure directory %s exists...", to)
		err := os.MkdirAll(to, 0755)
		if err != nil {
			return err
		}
		fmt.Println(" done")
		for _, entry := range r.Folder.Contents {
			err := WriteGo(from.Append(entry.Name), to+"/"+entry.Name.SnakeCase())
			if err != nil {
				return err
			}
		}
	}

	if r.IsSchema {
		fmt.Println(" schema found")
		if len(from) == 0 {
			return fmt.Errorf("no schema at root")
		}

		var pkg nameconv.Name
		if len(from) > 1 {
			pkg = from.SecondLast()
		} else {
			pkg = nameconv.Name{"types"}
		}

		name := from.Last()

		var typeDef bytes.Buffer
		typeDef.WriteString("type ")
		typeDef.WriteString(name.PascalCase())
		typeDef.WriteString(" struct {\n")

		imports := map[string]bool{}
		for _, field := range r.Schema.Fields {
			if field.Type.BaseType.Path.Length() > 0 {
				imports["github.com/schema-cafe/go-types"+field.Type.BaseType.Path.String()] = true
			}

			typeDef.WriteString("\t")
			typeDef.WriteString(field.Name.PascalCase())
			typeDef.WriteString(" ")
			typeDef.WriteString(field.Type.Golang(from))
			typeDef.WriteString(" `json:\"")
			typeDef.WriteString(field.Name.CamelCase())
			typeDef.WriteString("\"`\n")
		}

		typeDef.WriteString("}\n")

		var buf bytes.Buffer
		buf.WriteString("package ")
		buf.WriteString(pkg.AllLowerNoSpaces())
		buf.WriteString("\n\n")

		if len(imports) > 0 {
			buf.WriteString("import (\n")
			for imp := range imports {
				buf.WriteString("\t\"")
				buf.WriteString(imp)
				buf.WriteString("\"\n")
			}
			buf.WriteString(")\n\n")
		}

		buf.WriteString(typeDef.String())

		path := to + ".go"
		fmt.Printf("Writing Go file to %s...", path)
		err := os.WriteFile(path, buf.Bytes(), os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Println(" done")
	}

	if !r.IsFolder && !r.IsSchema {
		fmt.Println(" not found")
	}

	return nil
}
