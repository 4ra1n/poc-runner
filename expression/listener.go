package expression

import (
	"errors"
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"github.com/antlr4-go/antlr/v4"
)

type parserErrorListener struct {
	antlr.DefaultErrorListener
	Err error
}

func (l *parserErrorListener) SyntaxError(
	_ antlr.Recognizer,
	_ interface{},
	line, column int,
	msg string,
	_ antlr.RecognitionException) {
	if l.Err == nil {
		l.Err = xerr.Wrap(errors.New(fmt.Sprintf("syntax error %d:%d - %s\n", line, column, msg)))
	}
}
