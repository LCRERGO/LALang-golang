package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/LCRERGO/LALang/pkg/errors"
	grammar "github.com/LCRERGO/LALang/pkg/grammar/antlr"
	"github.com/LCRERGO/LALang/pkg/lexer"
	"github.com/LCRERGO/LALang/pkg/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func debugValue() int {
	debug, err := strconv.Atoi(os.Getenv("DEBUG"))
	if err != nil {
		debug = 0
	}

	return debug
}

/*
 * The main runner of the task
 */
func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <input> <output>\n", os.Args[0])
		os.Exit(1)
	}
	inFname, outFname := os.Args[1], os.Args[2]

	// Opens input an output files
	cs, err := antlr.NewFileStream(inFname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	outFile, err := os.OpenFile(outFname, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	debug := debugValue()

	// A defered routine to panic at errors and stop parsing symbols
	defer func(outFile *os.File) {
		if err := recover(); err != nil {
			// Type casting a inteface{} into an error
			outFile.WriteString(fmt.Errorf("%v", err).Error())
		}
		os.Exit(1)
	}(outFile)

	// Creating a new lexer
	lex := grammar.NewLALexer(cs)
	// Enabling Custom Error Handling
	lex.RemoveErrorListeners()
	lex.AddErrorListener(errors.NewLALexerErrorListener())

	lexer.RunLexer(lex, outFile, debug)

	tokStream := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)

	// Creating a new Parser
	par := grammar.NewLAParser(tokStream)
	// Enabling Custom Error Handling
	par.RemoveErrorListeners()
	par.AddErrorListener(errors.NewLASyntaxErrorListener())

	parser.RunParser(par, outFile, debug)
}
