package lexer

import (
	"fmt"
	"os"

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
