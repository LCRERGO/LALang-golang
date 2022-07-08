package lexer

import (
	"fmt"
	"os"

	grammar "github.com/LCRERGO/LALang/pkg/grammar/antlr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Recognizes each token in a specific way and writes them to out file
func RunLexer(lex *grammar.LALexer, outFile *os.File, debug int) {
	if debug > 0 {
		for t := lex.NextToken(); t.GetTokenType() != antlr.TokenEOF; t = lex.NextToken() {
			var outstr string
			tokenType := t.GetTokenType()

			switch tokenType {
			case grammar.LALexerCADEIA,
				grammar.LALexerIDENT,
				grammar.LALexerNUM_INT,
				grammar.LALexerNUM_REAL:
				// return output string on format <'match',token> e.g.: <'idade',IDENT>
				outstr = fmt.Sprintf("<'%s',%s>\n", t.GetText(), lex.SymbolicNames[t.GetTokenType()])
			default:
				//  case:
				// grammar.LALexerPALAVRA_CHAVE,
				// grammar.LALexerVIRGULA,
				// grammar.LALexerABREPARENTESES,
				// grammar.LALexerFECHAPARENTESES,
				// grammar.LALexerDELIM,
				// grammar.LALexerOP_ARIT,
				// grammar.LALexerOP_PON,
				// grammar.LALexerOP_REL,
				// grammar.LALexerATRIB:

				// return output string on format <'match','match'> e.g.: <'algoritmo','algoritmo'>
				outstr = fmt.Sprintf("<'%s','%s'>\n", t.GetText(), t.GetText())
			}
			outFile.WriteString(outstr)
		}
	}
}
