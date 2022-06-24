package main

import (
	"fmt"
	"os"

	"github.com/LCRERGO/LALang/pkg/errors"
	grammar "github.com/LCRERGO/LALang/pkg/grammar/antlr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Recognizes each token in a specific way and writes them to out file
func LexerTreatment(lex *grammar.LA, outFile *os.File) {
	// A defered routine to panic at errors and stop parsing symbols
	defer func(outFile *os.File) {
		if err := recover(); err != nil {
			// Type casting a inteface{} into an error
			outFile.WriteString(fmt.Errorf("%v", err).Error())
		}
		os.Exit(1)
	}(outFile)
	for t := lex.NextToken(); t.GetTokenType() != antlr.TokenEOF; t = lex.NextToken() {
		var outstr string
		tokenType := t.GetTokenType()

		switch tokenType {
		case grammar.LAPALAVRA_CHAVE,
			grammar.LAVIRGULA,
			grammar.LAABREPARENTESES,
			grammar.LAFECHAPARENTESES,
			grammar.LADELIM,
			grammar.LAOP_ARIT,
			grammar.LAOP_PON,
			grammar.LAOP_REL,
			grammar.LAATRIB:
			// return output string on format <'match','match'> e.g.: <'algoritmo','algoritmo'>
			outstr = fmt.Sprintf("<'%s','%s'>\n", t.GetText(), t.GetText())
		case grammar.LACADEIA,
			grammar.LAIDENT,
			grammar.LANUM_INT,
			grammar.LANUM_REAL:
			// return output string on format <'match',token> e.g.: <'idade',IDENT>
			outstr = fmt.Sprintf("<'%s',%s>\n", t.GetText(), lex.SymbolicNames[t.GetTokenType()])
		default:
			// return output string on format <'match','token'> e.g.:
			outstr = fmt.Sprintf("<'%s','%s'>\n", t.GetText(), lex.SymbolicNames[t.GetTokenType()])
		}
		outFile.WriteString(outstr)
	}
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

	// Creating a nil lexer
	lex := grammar.NewLA(cs)
	// Enabling Custom Error Handling
	lexerErrors := errors.NewLAErrorListener()
	lex.RemoveErrorListeners()
	lex.AddErrorListener(lexerErrors)

	LexerTreatment(lex, outFile)
}
