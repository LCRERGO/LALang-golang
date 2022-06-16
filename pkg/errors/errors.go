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

// LA Syntax Error Struct, agregates the line number and an object
// of type RecognitionException where utilities are defined.
type LASyntaxError struct {
	line int
	msg  string
}

// Creates a new LASyntaxError
func NewLASyntaxError() *LASyntaxError {
	return &LASyntaxError{}
}

// Raises all three possible errors:
// 1. comment not closed
// 2. string not closed
// 3. Symbol not identified
func (c LASyntaxError) Error() string {
	inputLineError := strings.Split(c.msg, `'`)[1]

	if inputLineError[0] == '{' && inputLineError[len(inputLineError)-1] != '}' {
		return fmt.Sprintf("Linha %d: comentario nao fechado\n", c.line)
	} else if inputLineError[0] == '"' && inputLineError[len(inputLineError)-1] != '"' {
		return fmt.Sprintf("Linha %d: cadeia literal nao fechada\n", c.line)
	}

	return fmt.Sprintf("Linha %d: %s - simbolo nao identificado\n", c.line, inputLineError)
}

// A wrapper for errors in the LA language
type LAErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

// Creates a new LAErrorListener
func NewLAErrorListener() *LAErrorListener {
	return &LAErrorListener{}
}

// Raises an error to the error listener.
// The parsing stops if an error is raised.
func (c *LAErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	ce := LASyntaxError{line: line, msg: msg}

	panic(ce)
}
