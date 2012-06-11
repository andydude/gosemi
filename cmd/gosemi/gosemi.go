package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/scanner"
	"go/token"
	"io"
	"os"
)

func main () {

	// read
	flag.Parse()
	filename := flag.Arg(0)
	fptr, err := os.Open(filename)
	if err != nil { panic(err) }
	frdr := bufio.NewReaderSize(fptr, 0x1000000)
	src, err := frdr.ReadSlice('\x7F')
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	// scan
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile(filename, fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		_, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}

		switch tok {
		case token.SEMICOLON:
			fmt.Print(";")

		case token.COMMENT:
			continue

		case token.IDENT:
			fallthrough
		case token.INT:
			fallthrough
		case token.FLOAT:
			fallthrough
		case token.IMAG:
			fallthrough
		case token.CHAR:
			fallthrough
		case token.STRING:
			fmt.Printf("%s ", lit)
		default:
			fmt.Printf("%s ", tok.String())
		}
	}
}