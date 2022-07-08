package parser

import (
	"os"

	grammar "github.com/LCRERGO/LALang/pkg/grammar/antlr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type LAListener struct {
	*grammar.BaseLAListener
}

func RunParser(par *grammar.LAParser, outFile *os.File, debug int) {

	antlr.ParseTreeWalkerDefault.Walk(&LAListener{}, par.Programa())
}
