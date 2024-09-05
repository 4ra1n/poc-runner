// Code generated from Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package expression // Expression
import "github.com/antlr4-go/antlr/v4"

// ExpressionListener is a complete listener for a parse tree produced by ExpressionParser.
type ExpressionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAdditiveExpression is called when entering the AdditiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterMemberAccessExpression is called when entering the MemberAccessExpression production.
	EnterMemberAccessExpression(c *MemberAccessExpressionContext)

	// EnterLogicalAndExpression is called when entering the LogicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterConstExpression is called when entering the ConstExpression production.
	EnterConstExpression(c *ConstExpressionContext)

	// EnterInExpression is called when entering the InExpression production.
	EnterInExpression(c *InExpressionContext)

	// EnterLogicalOrExpression is called when entering the LogicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterNotExpression is called when entering the NotExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterFunctionCallExpression is called when entering the FunctionCallExpression production.
	EnterFunctionCallExpression(c *FunctionCallExpressionContext)

	// EnterParenExpression is called when entering the ParenExpression production.
	EnterParenExpression(c *ParenExpressionContext)

	// EnterEqualityExpression is called when entering the EqualityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterMultiplicativeExpression is called when entering the MultiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterArrayAccessExpression is called when entering the ArrayAccessExpression production.
	EnterArrayAccessExpression(c *ArrayAccessExpressionContext)

	// EnterIdentifierAccessExpression is called when entering the IdentifierAccessExpression production.
	EnterIdentifierAccessExpression(c *IdentifierAccessExpressionContext)

	// EnterExpressionMember is called when entering the expressionMember production.
	EnterExpressionMember(c *ExpressionMemberContext)

	// EnterBooleanLiteral is called when entering the BooleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterBinaryStringLiteral is called when entering the BinaryStringLiteral production.
	EnterBinaryStringLiteral(c *BinaryStringLiteralContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterIntegerLiteral is called when entering the IntegerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterFloatingPointLiteral is called when entering the FloatingPointLiteral production.
	EnterFloatingPointLiteral(c *FloatingPointLiteralContext)

	// EnterExpressionArguments is called when entering the expressionArguments production.
	EnterExpressionArguments(c *ExpressionArgumentsContext)

	// EnterExpressionArgument is called when entering the expressionArgument production.
	EnterExpressionArgument(c *ExpressionArgumentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAdditiveExpression is called when exiting the AdditiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitMemberAccessExpression is called when exiting the MemberAccessExpression production.
	ExitMemberAccessExpression(c *MemberAccessExpressionContext)

	// ExitLogicalAndExpression is called when exiting the LogicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitConstExpression is called when exiting the ConstExpression production.
	ExitConstExpression(c *ConstExpressionContext)

	// ExitInExpression is called when exiting the InExpression production.
	ExitInExpression(c *InExpressionContext)

	// ExitLogicalOrExpression is called when exiting the LogicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitNotExpression is called when exiting the NotExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitFunctionCallExpression is called when exiting the FunctionCallExpression production.
	ExitFunctionCallExpression(c *FunctionCallExpressionContext)

	// ExitParenExpression is called when exiting the ParenExpression production.
	ExitParenExpression(c *ParenExpressionContext)

	// ExitEqualityExpression is called when exiting the EqualityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the MultiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitArrayAccessExpression is called when exiting the ArrayAccessExpression production.
	ExitArrayAccessExpression(c *ArrayAccessExpressionContext)

	// ExitIdentifierAccessExpression is called when exiting the IdentifierAccessExpression production.
	ExitIdentifierAccessExpression(c *IdentifierAccessExpressionContext)

	// ExitExpressionMember is called when exiting the expressionMember production.
	ExitExpressionMember(c *ExpressionMemberContext)

	// ExitBooleanLiteral is called when exiting the BooleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitBinaryStringLiteral is called when exiting the BinaryStringLiteral production.
	ExitBinaryStringLiteral(c *BinaryStringLiteralContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitIntegerLiteral is called when exiting the IntegerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitFloatingPointLiteral is called when exiting the FloatingPointLiteral production.
	ExitFloatingPointLiteral(c *FloatingPointLiteralContext)

	// ExitExpressionArguments is called when exiting the expressionArguments production.
	ExitExpressionArguments(c *ExpressionArgumentsContext)

	// ExitExpressionArgument is called when exiting the expressionArgument production.
	ExitExpressionArgument(c *ExpressionArgumentContext)
}
