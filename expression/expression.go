/*
 * poc-runner project
 * Copyright (C) 2024 4ra1n
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package expression

import (
	"context"
	"regexp"

	"github.com/antlr4-go/antlr/v4"

	"github.com/4ra1n/poc-runner/xerr"
)

type Environment struct {
	builtin map[string]EFunction
	vars    Vars
	Context context.Context
}

var expressionStringRegex = regexp.MustCompile(`\{\{[^{}]+?}}`)

func NewEnvironment(context context.Context, vars ...Vars) *Environment {
	if vars == nil || len(vars) == 0 {
		vars = combinedVars{}
	}
	return &Environment{
		builtin: builtin,
		vars:    combinedVars(vars),
		Context: context,
	}
}

func (e *Environment) GetString(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	return expressionStringRegex.ReplaceAllStringFunc(str, func(s string) string {
		value, err := e.Eval(s[2 : len(s)-2])
		if err == nil && value != nil {
			return string(value.ToString())
		}
		if value == nil {
			return ""
		}
		return string(value.ToString())
	}), nil
}

func (e *Environment) Verify(expression string) (err error) {
	_, err = e.parse(expression)
	return
}

func (e *Environment) Eval(expression string) (EValue, error) {
	expr, err := e.parse(expression)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	visitor := &evaluator{env: e}
	val := expr.Accept(visitor)
	if visitor.err != nil {
		return nil, xerr.Wrap(visitor.err)
	}
	return val.(EValue), nil
}

func (e *Environment) EvalWithVars(expression string, vars Vars) (EValue, error) {
	backup := e.vars
	e.vars = combinedVars{
		backup,
		vars,
	}
	value, err := e.Eval(expression)
	e.vars = backup
	return value, err
}

func (e *Environment) parse(expression string) (IExpressionContext, error) {
	input := antlr.NewInputStream(expression)
	lexer := NewExpressionLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewExpressionParser(stream)
	var parserErrorListenerVar = new(parserErrorListener)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrorListenerVar)
	expr := parser.Expression()
	if parserErrorListenerVar.Err != nil {
		return nil, xerr.Wrap(parserErrorListenerVar.Err)
	}
	return expr, nil
}
