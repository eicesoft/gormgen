package main

import (
	"flag"
	"github.com/eicesoft/gormgen/gormgen"
	"log"
	"strings"
)

var (
	input      string
	structName string
)

func init() {
	flag.StringVar(&input, "i", "", "[Required] The name of schema struct to generate structs for, comma seperated")
	flag.StringVar(&structName, "s", "", "[Required] The name of the input file dir")

	if !flag.Parsed() {
		flag.Parse()
	}
}

func main() {
	if input == "" {
		flag.Usage()
		return
	}

	if structName == "" {
		flag.Usage()
		return
	}

	structs := strings.Split(structName, ",")
	gen := gormgen.NewGenerator(input)
	p := gormgen.NewParser(input)

	if err := gen.ParserAST(p, structs).Generate().Flush(); err != nil {
		log.Fatalln(err)
	}
}
