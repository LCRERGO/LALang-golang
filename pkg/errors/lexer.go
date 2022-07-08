package errors

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

/*
 * Custom errors definition as stated in:
 * https://stackoverflow.com/questions/66067549/how-to-write-a-custom-error-reporter-in-go-target-of-antlr
 */

// LA Lexer Error Struct, agregates the line number and the
// original error message of type RecognitionException
// where utilities are defined.
type LALexerError struct {
	line int
	msg  string
}

// Creates a new LALexerError
func NewLALexerError() *LALexerError {
	return &LALexerError{}
}

// Raises all three possible errors:
// 1. comment not closed
// 2. string not closed
// 3. Symbol not identified
func (l LALexerError) Error() string {
	inputLineError := strings.Split(l.msg, `'`)[1]
	var ret string

	if inputLineError[0] == '{' && inputLineError[len(inputLineError)-1] != '}' {

		ret = fmt.Sprintf("Linha %d: comentario nao fechado\n", l.line)
	} else if inputLineError[0] == '"' && inputLineError[len(inputLineError)-1] != '"' {
		ret = fmt.Sprintf("Linha %d: cadeia literal nao fechada\n", l.line)
	} else {
		ret = fmt.Sprintf("Linha %d: %s - simbolo nao identificado\n", l.line, inputLineError)
	}
	ret += "Fim da compilacao\n"

	return ret
}

// A wrapper for lexer errors in the LA language
type LALexerErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

// Creates a new LALexerErrorListener
func NewLALexerErrorListener() *LALexerErrorListener {
	return &LALexerErrorListener{}
}

// Raises a lexer error to the error listener.
// The parsing stops if an error is raised.
func (c *LALexerErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	le := LALexerError{line: line, msg: msg}

	panic(le)
}
