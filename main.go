package main

import (
	"fmt"
	"os"

	grammar "github.com/LCRERGO/LALang/pkg/grammar/antlr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

/*
 * The main runner of the task
 */

func main() {
	// TODO: alter for len < 3
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <input>\n", os.Args[0])
		os.Exit(1)
	}
	cs, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	lex := grammar.NewLA(cs)
	//outFile := os.OpenFile(os.Args[2], os.O_RDWR, 0766)

	for t := lex.NextToken(); t.GetTokenType() != antlr.TokenEOF; t = lex.NextToken() {
		var outstr string
		//errDispach := lex.GetErrorListenerDispatch()
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
		print(outstr)
	}

}
