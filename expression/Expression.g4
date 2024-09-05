grammar Expression;

expression
    : expressionSingle EOF
    ;

expressionSingle
    : expressionConst                                       # ConstExpression
    | Identifier                                            # IdentifierAccessExpression
    | '(' expressionSingle ')'                              # ParenExpression
    | expressionSingle expressionMember                     # MemberAccessExpression
    | expressionSingle '[' IntegerLiteral ']'               # ArrayAccessExpression
    | expressionSingle '(' expressionArguments ')'          # FunctionCallExpression
    | '!' expressionSingle                                  # NotExpression
    | expressionSingle ('+' | '-') expressionSingle         # AdditiveExpression
    | expressionSingle ('*' | '/') expressionSingle         # MultiplicativeExpression
    | expressionSingle ('==' | '!=') expressionSingle       # EqualityExpression
    | expressionSingle 'in' expressionSingle                # InExpression
    | expressionSingle '&&' expressionSingle                # LogicalAndExpression
    | expressionSingle '||' expressionSingle                # LogicalOrExpression
    ;

expressionMember
    : '.' 'b'? Identifier?
    | '[' StringLiteral ']'
    ;

expressionConst
    : BooleanLiteral                                    # BooleanLiteral
    | 'b' StringLiteral                                 # BinaryStringLiteral
    | StringLiteral                                     # StringLiteral
    | IntegerLiteral                                    # IntegerLiteral
    | FloatingPointLiteral                              # FloatingPointLiteral
    ;

expressionArguments
    : expressionArgument (',' expressionArgument)*
    ;

expressionArgument
    :
    | Identifier
    | IntegerLiteral
    | StringLiteral
    | expressionSingle
    ;


DOT : '.';
LBRACK : '[';
RBRACK : ']';
LPAREN : '(';
RPAREN : ')';

PLUS: '+';
MINUS: '-';
MULTIPLY : '*';
DIVIDE : '/';

EQUAL : '==';
NOTEQUAL : '!=';

WhiteSpaces : [\t\u000B\u000C\u0020\u00A0]+ -> channel(HIDDEN);


BooleanLiteral
	:	'true'
	|	'false'
	;

Identifier
	: [a-zA-Z$_]+[a-zA-Z0-9$_]*
	;

IntegerLiteral
	:	DecimalIntegerLiteral
	|	HexIntegerLiteral
	|	OctalIntegerLiteral
	|	BinaryIntegerLiteral
	;


FloatingPointLiteral
	:	DecimalFloatingPointLiteral
	;


StringLiteral
	:	'"' DoubleStringCharacter* '"'
	|   '\'' SingleStringCharacter* '\''
	;

Whitespace
    : [ \t\r\n]+ -> channel(HIDDEN)
    ;

fragment
DoubleStringCharacter
    : ~["\\\r\n]
    | '\\' EscapeSequence
    ;

fragment
SingleStringCharacter
    : ~['\\\r\n]
    | '\\' EscapeSequence
    ;

fragment
EscapeSequence
    : CharacterEscapeSequence
    | HexEscapeSequence
    | UnicodeEscapeSequence
    | ExtendedUnicodeEscapeSequence
    ;

fragment
CharacterEscapeSequence
    : SingleEscapeCharacter
    | NonEscapeCharacter
    ;

fragment
HexEscapeSequence
    : 'x' HexDigit HexDigit
    ;

fragment
UnicodeEscapeSequence
    : 'u' HexDigit HexDigit HexDigit HexDigit
    | 'u' '{' HexDigit HexDigit+ '}'
    ;

fragment
ExtendedUnicodeEscapeSequence
    : 'u' '{' HexDigit+ '}'
    ;

fragment
SingleEscapeCharacter
    : ['"\\bfnrtv]
    ;

fragment
NonEscapeCharacter
    : ~['"\\bfnrtv0-9xu\r\n]
    ;

fragment
HexEscape
    :   '\\' 'x'+  HexDigit HexDigit
    ;

fragment
DecimalIntegerLiteral
	:	DecimalNumeral IntegerTypeSuffix?
	;

fragment
HexIntegerLiteral
	:	HexNumeral IntegerTypeSuffix?
	;

fragment
OctalIntegerLiteral
	:	OctalNumeral IntegerTypeSuffix?
	;

fragment
BinaryIntegerLiteral
	:	BinaryNumeral IntegerTypeSuffix?
	;

fragment
IntegerTypeSuffix
	:	[lL]
	;

fragment
DecimalNumeral
	:	'0'
	|	NonZeroDigit (Digits? | Underscores Digits)
	;

fragment
Digits
	:	Digit (DigitsAndUnderscores? Digit)?
	;

fragment
Digit
	:	'0'
	|	NonZeroDigit
	;

fragment
NonZeroDigit
	:	[1-9]
	;

fragment
DigitsAndUnderscores
	:	DigitOrUnderscore+
	;

fragment
DigitOrUnderscore
	:	Digit
	|	'_'
	;

fragment
Underscores
	:	'_'+
	;

fragment
HexNumeral
	:	'0' [xX] HexDigits
	;

fragment
HexDigits
	:	HexDigit (HexDigitsAndUnderscores? HexDigit)?
	;

fragment
HexDigit
	:	[0-9a-fA-F]
	;

fragment
HexDigitsAndUnderscores
	:	HexDigitOrUnderscore+
	;

fragment
HexDigitOrUnderscore
	:	HexDigit
	|	'_'
	;

fragment
OctalNumeral
	:	'0' Underscores? OctalDigits
	;

fragment
OctalDigits
	:	OctalDigit (OctalDigitsAndUnderscores? OctalDigit)?
	;

fragment
OctalDigit
	:	[0-7]
	;

fragment
OctalDigitsAndUnderscores
	:	OctalDigitOrUnderscore+
	;

fragment
OctalDigitOrUnderscore
	:	OctalDigit
	|	'_'
	;

fragment
BinaryNumeral
	:	'0' [bB] BinaryDigits
	;

fragment
BinaryDigits
	:	BinaryDigit (BinaryDigitsAndUnderscores? BinaryDigit)?
	;

fragment
BinaryDigit
	:	[01]
	;

fragment
BinaryDigitsAndUnderscores
	:	BinaryDigitOrUnderscore+
	;

fragment
BinaryDigitOrUnderscore
	:	BinaryDigit
	|	'_'
	;

fragment
DecimalFloatingPointLiteral
	:	Digits '.' Digits? ExponentPart? FloatTypeSuffix?
	|	'.' Digits ExponentPart? FloatTypeSuffix?
	|	Digits ExponentPart FloatTypeSuffix?
	|	Digits FloatTypeSuffix
	;

fragment
ExponentPart
	:	ExponentIndicator SignedInteger
	;

fragment
ExponentIndicator
	:	[eE]
	;

fragment
SignedInteger
	:	Sign? Digits
	;

fragment
Sign
	:	[+-]
	;

fragment
FloatTypeSuffix
	:	[fFdD]
	;

fragment
HexadecimalFloatingPointLiteral
	:	HexSignificand BinaryExponent FloatTypeSuffix?
	;

fragment
HexSignificand
	:	HexNumeral '.'?
	|	'0' [xX] HexDigits? '.' HexDigits
	;

fragment
BinaryExponent
	:	BinaryExponentIndicator SignedInteger
	;

fragment
BinaryExponentIndicator
	:	[pP]
	;