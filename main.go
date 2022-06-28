package main

import (
	"fmt"
	"os"

	"github.com/LCRERGO/LALang/pkg/errors"
	grammar "github.com/LCRERGO/LALang/pkg/grammar/antlr"
	"github.com/LCRERGO/LALang/pkg/lexer"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

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

	// Creating a nil lexer
	lex := grammar.NewLALexer(cs)
	// Enabling Custom Error Handling
	lexerErrors := errors.NewLAErrorListener()
	lex.RemoveErrorListeners()
	lex.AddErrorListener(lexerErrors)

	lexer.LexerTreatment(lex, outFile)
}
