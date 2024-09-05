// Code generated from Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package expression // Expression
import "github.com/antlr4-go/antlr/v4"

// BaseExpressionListener is a complete listener for a parse tree produced by ExpressionParser.
type BaseExpressionListener struct{}

var _ ExpressionListener = &BaseExpressionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExpressionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExpressionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseExpressionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseExpressionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAdditiveExpression is called when production AdditiveExpression is entered.
func (s *BaseExpressionListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production AdditiveExpression is exited.
func (s *BaseExpressionListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMemberAccessExpression is called when production MemberAccessExpression is entered.
func (s *BaseExpressionListener) EnterMemberAccessExpression(ctx *MemberAccessExpressionContext) {}

// ExitMemberAccessExpression is called when production MemberAccessExpression is exited.
func (s *BaseExpressionListener) ExitMemberAccessExpression(ctx *MemberAccessExpressionContext) {}

// EnterLogicalAndExpression is called when production LogicalAndExpression is entered.
func (s *BaseExpressionListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production LogicalAndExpression is exited.
func (s *BaseExpressionListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterConstExpression is called when production ConstExpression is entered.
func (s *BaseExpressionListener) EnterConstExpression(ctx *ConstExpressionContext) {}

// ExitConstExpression is called when production ConstExpression is exited.
func (s *BaseExpressionListener) ExitConstExpression(ctx *ConstExpressionContext) {}

// EnterInExpression is called when production InExpression is entered.
func (s *BaseExpressionListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production InExpression is exited.
func (s *BaseExpressionListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterLogicalOrExpression is called when production LogicalOrExpression is entered.
func (s *BaseExpressionListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production LogicalOrExpression is exited.
func (s *BaseExpressionListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterNotExpression is called when production NotExpression is entered.
func (s *BaseExpressionListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production NotExpression is exited.
func (s *BaseExpressionListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterFunctionCallExpression is called when production FunctionCallExpression is entered.
func (s *BaseExpressionListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// ExitFunctionCallExpression is called when production FunctionCallExpression is exited.
func (s *BaseExpressionListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// EnterParenExpression is called when production ParenExpression is entered.
func (s *BaseExpressionListener) EnterParenExpression(ctx *ParenExpressionContext) {}

// ExitParenExpression is called when production ParenExpression is exited.
func (s *BaseExpressionListener) ExitParenExpression(ctx *ParenExpressionContext) {}

// EnterEqualityExpression is called when production EqualityExpression is entered.
func (s *BaseExpressionListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production EqualityExpression is exited.
func (s *BaseExpressionListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterMultiplicativeExpression is called when production MultiplicativeExpression is entered.
func (s *BaseExpressionListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production MultiplicativeExpression is exited.
func (s *BaseExpressionListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterArrayAccessExpression is called when production ArrayAccessExpression is entered.
func (s *BaseExpressionListener) EnterArrayAccessExpression(ctx *ArrayAccessExpressionContext) {}

// ExitArrayAccessExpression is called when production ArrayAccessExpression is exited.
func (s *BaseExpressionListener) ExitArrayAccessExpression(ctx *ArrayAccessExpressionContext) {}

// EnterIdentifierAccessExpression is called when production IdentifierAccessExpression is entered.
func (s *BaseExpressionListener) EnterIdentifierAccessExpression(ctx *IdentifierAccessExpressionContext) {
}

// ExitIdentifierAccessExpression is called when production IdentifierAccessExpression is exited.
func (s *BaseExpressionListener) ExitIdentifierAccessExpression(ctx *IdentifierAccessExpressionContext) {
}

// EnterExpressionMember is called when production expressionMember is entered.
func (s *BaseExpressionListener) EnterExpressionMember(ctx *ExpressionMemberContext) {}

// ExitExpressionMember is called when production expressionMember is exited.
func (s *BaseExpressionListener) ExitExpressionMember(ctx *ExpressionMemberContext) {}

// EnterBooleanLiteral is called when production BooleanLiteral is entered.
func (s *BaseExpressionListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production BooleanLiteral is exited.
func (s *BaseExpressionListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterBinaryStringLiteral is called when production BinaryStringLiteral is entered.
func (s *BaseExpressionListener) EnterBinaryStringLiteral(ctx *BinaryStringLiteralContext) {}

// ExitBinaryStringLiteral is called when production BinaryStringLiteral is exited.
func (s *BaseExpressionListener) ExitBinaryStringLiteral(ctx *BinaryStringLiteralContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseExpressionListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseExpressionListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterIntegerLiteral is called when production IntegerLiteral is entered.
func (s *BaseExpressionListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production IntegerLiteral is exited.
func (s *BaseExpressionListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterFloatingPointLiteral is called when production FloatingPointLiteral is entered.
func (s *BaseExpressionListener) EnterFloatingPointLiteral(ctx *FloatingPointLiteralContext) {}

// ExitFloatingPointLiteral is called when production FloatingPointLiteral is exited.
func (s *BaseExpressionListener) ExitFloatingPointLiteral(ctx *FloatingPointLiteralContext) {}

// EnterExpressionArguments is called when production expressionArguments is entered.
func (s *BaseExpressionListener) EnterExpressionArguments(ctx *ExpressionArgumentsContext) {}

// ExitExpressionArguments is called when production expressionArguments is exited.
func (s *BaseExpressionListener) ExitExpressionArguments(ctx *ExpressionArgumentsContext) {}

// EnterExpressionArgument is called when production expressionArgument is entered.
func (s *BaseExpressionListener) EnterExpressionArgument(ctx *ExpressionArgumentContext) {}

// ExitExpressionArgument is called when production expressionArgument is exited.
func (s *BaseExpressionListener) ExitExpressionArgument(ctx *ExpressionArgumentContext) {}
