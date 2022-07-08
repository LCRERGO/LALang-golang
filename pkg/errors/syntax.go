package errors

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// LA Syntax Error Struct, agregates the line number and an object
// of type RecognitionException where utilities are defined.
type LASyntaxError struct {
	line            int
	msg             string
	offendingSymbol interface{}
}

// Creates a new LALexerError
func NewLASyntaxError() *LALexerError {
	return &LALexerError{}
}

func (c LASyntaxError) Error() string {
	var ret string

	t := c.offendingSymbol.(antlr.Token)
	switch t.GetTokenType() {
	case antlr.TokenEOF:
		ret = fmt.Sprintf("Linha %d: erro sintatico proximo a EOF\n", c.line)
	default:
		ret = fmt.Sprintf("Linha %d: erro sintatico proximo a %s\n", c.line, t.GetText())
	}
	ret += "Fim da compilacao\n"

	return ret
}

// A wrapper for syntax errors in the LA language
type LASyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

// Creates a new LASyntaxErrorListener
func NewLASyntaxErrorListener() *LASyntaxErrorListener {
	return &LASyntaxErrorListener{}
}

// Raises a lexer error to the error listener.
// The parsing stops if an error is raised.
func (c *LASyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	tok := offendingSymbol.(antlr.Token)
	println(tok.GetText())
	se := LASyntaxError{line: line, msg: msg, offendingSymbol: offendingSymbol}

	panic(se)
}
