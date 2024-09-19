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
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/4ra1n/poc-runner/util"
)

type evaluator struct {
	env *Environment
	err error
}

func (e *evaluator) Visits(trees []antlr.ParseTree) []EValue {
	var values []EValue
	for _, vc := range trees {
		value := e.Visit(vc)
		if e.err != nil {
			return nil
		}
		values = append(values, value.(EValue))
	}
	return values
}

func (e *evaluator) Visit(tree antlr.ParseTree) any {
	if e.err != nil {
		return nil
	}
	switch tree.(type) {
	case *AdditiveExpressionContext:
		return e.visitAdditiveExpressionContext(tree.(*AdditiveExpressionContext))
	case *MultiplicativeExpressionContext:
		return e.visitMultiplicativeExpressionContext(tree.(*MultiplicativeExpressionContext))
	case *ConstExpressionContext:
		return e.Visit(tree.(*ConstExpressionContext).ExpressionConst())
	case *IntegerLiteralContext:
		return e.visitIntegerLiteralContext(tree.(*IntegerLiteralContext))
	case *StringLiteralContext:
		return e.visitStringLiteralContext(tree.(*StringLiteralContext))
	case *BinaryStringLiteralContext:
		return e.visitBinaryStringLiteralContext(tree.(*BinaryStringLiteralContext))
	case *FunctionCallExpressionContext:
		return e.visitFunctionCallExpressionContext(tree.(*FunctionCallExpressionContext))
	case *IdentifierAccessExpressionContext:
		return e.visitIdentifierAccessExpressionContext(tree.(*IdentifierAccessExpressionContext))
	case *FloatingPointLiteralContext:
		return e.visitFloatingPointLiteralContext(tree.(*FloatingPointLiteralContext))
	case *BooleanLiteralContext:
		return e.visitBooleanLiteralContext(tree.(*BooleanLiteralContext))
	case *EqualityExpressionContext:
		return e.visitEqualityExpressionContext(tree.(*EqualityExpressionContext))
	case *MemberAccessExpressionContext:
		return e.visitMemberAccessExpressionContext(tree.(*MemberAccessExpressionContext))
	case *LogicalAndExpressionContext:
		return e.visitLogicalAndExpressionContext(tree.(*LogicalAndExpressionContext))
	case *LogicalOrExpressionContext:
		return e.visitLogicalOrExpressionContext(tree.(*LogicalOrExpressionContext))
	case *ParenExpressionContext:
		return e.visitParenExpressionContext(tree.(*ParenExpressionContext))
	case *InExpressionContext:
		return e.visitInExpressionContext(tree.(*InExpressionContext))
	case *NotExpressionContext:
		return e.visitNotExpressionContext(tree.(*NotExpressionContext))
	}
	panic(reflect.TypeOf(tree))
}

func (e *evaluator) visitIdentifierAccessExpressionContext(c *IdentifierAccessExpressionContext) any {
	return e.visitIdentifier(c.Identifier())
}

func (e *evaluator) visitIdentifier(n antlr.TerminalNode) any {
	variable := n.GetText()
	if value, ok := e.env.builtin[variable]; ok {
		return value
	}
	value, err := e.env.vars.GetValue(e.env, variable)
	if err != nil {
		return e.Error(err)
	}
	return value
}

func (e *evaluator) visitFunctionCallExpressionContext(c *FunctionCallExpressionContext) any {
	callee, ok := e.Visit(c.ExpressionSingle()).(EFunction)
	if !ok {
		return e.Error(fmt.Errorf("%s is not function", c.ExpressionSingle().GetText()))
	}
	expressionArguments := c.ExpressionArguments()
	var args []EValue
	if expressionArguments.GetChildCount() == 1 && expressionArguments.ExpressionArgument(0).GetText() == "" {
		args = make([]EValue, 0, 0)
	} else {
		args = util.Map(expressionArguments.AllExpressionArgument(),
			func(item IExpressionArgumentContext, index int) EValue {
				if v := item.StringLiteral(); v != nil {
					return e.visitStringLiteral(v).(EValue)
				}
				if v := item.IntegerLiteral(); v != nil {
					return e.visitIntegerLiteral(v).(EValue)
				}
				if v := item.Identifier(); v != nil {
					return e.visitIdentifier(v).(EValue)
				}
				if v := item.ExpressionSingle(); v != nil {
					return e.Visit(v).(EValue)
				}
				panic("unreached code")
			})
	}
	if e.err != nil {
		return nil
	}
	value, err := callee.Call(args)
	if err != nil {
		return e.Error(err)
	}
	return value
}

func (e *evaluator) visitIntegerLiteralContext(c *IntegerLiteralContext) any {
	return e.visitIntegerLiteral(c.IntegerLiteral())
}

func (e *evaluator) visitIntegerLiteral(n antlr.TerminalNode) any {
	val, _ := strconv.Atoi(n.GetText())
	return EInt(val)
}

func (e *evaluator) visitFloatingPointLiteralContext(c *FloatingPointLiteralContext) any {
	return e.visitFloatLiteral(c.FloatingPointLiteral())
}

func (e *evaluator) visitFloatLiteral(n antlr.TerminalNode) any {
	val, _ := strconv.ParseFloat(n.GetText(), 64)
	return EFloat(val)
}

func (e *evaluator) visitBinaryStringLiteralContext(c *BinaryStringLiteralContext) any {
	return e.visitBinaryStringLiteral(c.StringLiteral())
}

func (e *evaluator) visitBinaryStringLiteral(n antlr.TerminalNode) any {
	text := n.GetText()
	text, err := strconv.Unquote(`"` + text[1:len(text)-1] + `"`)
	if err != nil {
		return e.Error(err)
	}
	return EBytes(text)
}

func (e *evaluator) visitStringLiteralContext(c *StringLiteralContext) any {
	return e.visitStringLiteral(c.StringLiteral())
}

func (e *evaluator) visitStringLiteral(n antlr.TerminalNode) any {
	text := n.GetText()
	var err error
	if text[0] == byte('\'') {
		text, err = strconv.Unquote(`"` + strings.ReplaceAll(text[1:len(text)-1], `"`, `\"`) + `"`)
	} else {
		text, err = strconv.Unquote(text)
	}
	if err != nil {
		return e.Error(err)
	}
	return EString(text)
}

func (e *evaluator) visitBooleanLiteralContext(c *BooleanLiteralContext) any {
	return e.visitBooleanLiteral(c.BooleanLiteral())
}

func (e *evaluator) visitBooleanLiteral(n antlr.TerminalNode) any {
	if n.GetText() == "true" {
		return EBool(true)
	}
	return EBool(false)
}

func (e *evaluator) visitAdditiveExpressionContext(c *AdditiveExpressionContext) any {
	left := e.Visit(c.ExpressionSingle(0))
	right := e.Visit(c.ExpressionSingle(1))
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return e.Error(errors.New(fmt.Sprintf("%s <OP> %s", reflect.TypeOf(left), reflect.TypeOf(right))))
	}
	switch {
	case c.PLUS() != nil:
		switch left.(type) {
		case EInt:
			return eValuePlus[EInt](left, right)
		case EFloat:
			return eValuePlus[EFloat](left, right)
		case EString:
			return eValuePlus[EString](left, right)
		case EBytes:
			return append(left.(EBytes), right.(EBytes)...)
		}
	case c.MINUS() != nil:
		switch left.(type) {
		case EInt:
			return eValueMinus[EInt](left, right)
		case EFloat:
			return eValueMinus[EFloat](left, right)
		}
	}
	panic("unreached code")
}

func (e *evaluator) visitMultiplicativeExpressionContext(c *MultiplicativeExpressionContext) any {
	left := e.Visit(c.ExpressionSingle(0))
	right := e.Visit(c.ExpressionSingle(1))
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return e.Error(errors.New(
			fmt.Sprintf("%s <OP> %s", reflect.TypeOf(left), reflect.TypeOf(right))))
	}
	switch {
	case c.MULTIPLY() != nil:
		switch left.(type) {
		case EInt:
			return eValueMultiply[EInt](left, right)
		case EFloat:
			return eValueMultiply[EFloat](left, right)
		}
	case c.DIVIDE() != nil:
		switch left.(type) {
		case EInt:
			return eValueDivide[EInt](left, right)
		case EFloat:
			return eValueDivide[EFloat](left, right)
		}
	}
	panic("unreached code")
}

func (e *evaluator) visitEqualityExpressionContext(c *EqualityExpressionContext) any {
	left := e.Visit(c.ExpressionSingle(0))
	right := e.Visit(c.ExpressionSingle(1))

	// FIX BUG FOR NOT EQUAL
	switch {
	case c.EQUAL() != nil:
		return EBool(reflect.DeepEqual(left, right))
	case c.NOTEQUAL() != nil:
		return !EBool(reflect.DeepEqual(left, right))
	}
	panic("unreached code")
}

func (e *evaluator) visitLogicalAndExpressionContext(c *LogicalAndExpressionContext) any {
	left := e.Visit(c.ExpressionSingle(0))
	leftValue, ok := left.(EBool)
	if !ok {
		return e.Error(fmt.Errorf("expect EBool, got %s", reflect.TypeOf(left)))
	}
	if !leftValue {
		return EBool(false)
	}
	right := e.Visit(c.ExpressionSingle(1))
	rightValue, ok := right.(EBool)
	if !ok {
		return e.Error(fmt.Errorf("expect EBool, got %s", reflect.TypeOf(right)))
	}
	return rightValue
}

func (e *evaluator) visitLogicalOrExpressionContext(c *LogicalOrExpressionContext) any {
	left := e.Visit(c.ExpressionSingle(0))
	leftValue, ok := left.(EBool)
	if !ok {
		return e.Error(fmt.Errorf(
			"expect EBool, got %s", reflect.TypeOf(left)))
	}
	if leftValue {
		return EBool(true)
	}
	right := e.Visit(c.ExpressionSingle(1))
	rightValue, ok := right.(EBool)
	if !ok {
		return e.Error(fmt.Errorf("expect EBool, got %s", reflect.TypeOf(right)))
	}
	return rightValue
}

func (e *evaluator) visitMemberAccessExpressionContext(c *MemberAccessExpressionContext) any {
	left := e.Visit(c.ExpressionSingle())
	if e.err != nil {
		return nil
	}
	object, ok := left.(EObject)
	if !ok {
		return e.Error(
			fmt.Errorf("expect EObject, got %s", reflect.TypeOf(left)))
	}
	expressionMember := c.ExpressionMember()
	var member string
	if s := expressionMember.StringLiteral(); s != nil {
		text := s.GetText()
		member = text[1 : len(text)-1]
	} else {
		member = expressionMember.GetText()[1:]
	}
	value, err := object.Get(member)
	if err != nil {
		return e.Error(err)
	}
	return value
}

func (e *evaluator) visitParenExpressionContext(c *ParenExpressionContext) any {
	return e.Visit(c.ExpressionSingle())
}

func (e *evaluator) visitInExpressionContext(c *InExpressionContext) any {
	left := e.Visit(c.ExpressionSingle(0))
	if e.err != nil {
		return nil
	}
	key, ok := left.(EString)
	if !ok {
		return e.Error(
			fmt.Errorf("expect EString, got %s", reflect.TypeOf(left)))
	}
	right := e.Visit(c.ExpressionSingle(1))
	if e.err != nil {
		return nil
	}
	object, ok := right.(EObject)
	if !ok {
		return e.Error(
			fmt.Errorf("expect EObject, got %s", reflect.TypeOf(right)))
	}
	return EBool(util.Contains(object.Keys(), string(key)))
}

func (e *evaluator) visitNotExpressionContext(c *NotExpressionContext) any {
	expr := e.Visit(c.ExpressionSingle())
	value, ok := expr.(EBool)
	if !ok {
		return e.Error(fmt.Errorf("expect EBool, got %s", reflect.TypeOf(expr)))
	}
	return !value
}

func (e *evaluator) Error(err error) any {
	if e.err == nil {
		e.err = err
	}
	return nil
}

func (e *evaluator) VisitChildren(node antlr.RuleNode) any {
	return e.Visit(node.GetChild(0).(antlr.ParseTree))
}

func (e *evaluator) VisitTerminal(node antlr.TerminalNode) any {
	return nil
}

func (e *evaluator) VisitErrorNode(node antlr.ErrorNode) any {
	return nil
}

func eValuePlus[T EInt | EFloat | EString](left any, right any) T {
	return left.(T) + right.(T)
}

func eValueMinus[T EInt | EFloat](left any, right any) T {
	return left.(T) - right.(T)
}

func eValueMultiply[T EInt | EFloat](left any, right any) T {
	return left.(T) * right.(T)
}

func eValueDivide[T EInt | EFloat](left any, right any) T {
	return left.(T) / right.(T)
}
