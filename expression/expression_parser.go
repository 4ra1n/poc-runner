// Code generated from Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package expression // Expression
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ExpressionParser struct {
	*antlr.BaseParser
}

var ExpressionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func expressionParserInit() {
	staticData := &ExpressionParserStaticData
	staticData.LiteralNames = []string{
		"", "'!'", "'in'", "'&&'", "'||'", "'b'", "','", "'.'", "'['", "']'",
		"'('", "')'", "'+'", "'-'", "'*'", "'/'", "'=='", "'!='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "DOT", "LBRACK", "RBRACK", "LPAREN", "RPAREN",
		"PLUS", "MINUS", "MULTIPLY", "DIVIDE", "EQUAL", "NOTEQUAL", "WhiteSpaces",
		"BooleanLiteral", "Identifier", "IntegerLiteral", "FloatingPointLiteral",
		"StringLiteral", "Whitespace",
	}
	staticData.RuleNames = []string{
		"expression", "expressionSingle", "expressionMember", "expressionConst",
		"expressionArguments", "expressionArgument",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 24, 96, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 25, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 56, 8,
		1, 10, 1, 12, 1, 59, 9, 1, 1, 2, 1, 2, 3, 2, 63, 8, 2, 1, 2, 3, 2, 66,
		8, 2, 1, 2, 1, 2, 1, 2, 3, 2, 71, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 3, 3, 79, 8, 3, 1, 4, 1, 4, 1, 4, 5, 4, 84, 8, 4, 10, 4, 12, 4, 87,
		9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 94, 8, 5, 1, 5, 0, 1, 2, 6, 0,
		2, 4, 6, 8, 10, 0, 3, 1, 0, 12, 13, 1, 0, 14, 15, 1, 0, 16, 17, 113, 0,
		12, 1, 0, 0, 0, 2, 24, 1, 0, 0, 0, 4, 70, 1, 0, 0, 0, 6, 78, 1, 0, 0, 0,
		8, 80, 1, 0, 0, 0, 10, 93, 1, 0, 0, 0, 12, 13, 3, 2, 1, 0, 13, 14, 5, 0,
		0, 1, 14, 1, 1, 0, 0, 0, 15, 16, 6, 1, -1, 0, 16, 25, 3, 6, 3, 0, 17, 25,
		5, 20, 0, 0, 18, 19, 5, 10, 0, 0, 19, 20, 3, 2, 1, 0, 20, 21, 5, 11, 0,
		0, 21, 25, 1, 0, 0, 0, 22, 23, 5, 1, 0, 0, 23, 25, 3, 2, 1, 7, 24, 15,
		1, 0, 0, 0, 24, 17, 1, 0, 0, 0, 24, 18, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0,
		25, 57, 1, 0, 0, 0, 26, 27, 10, 6, 0, 0, 27, 28, 7, 0, 0, 0, 28, 56, 3,
		2, 1, 7, 29, 30, 10, 5, 0, 0, 30, 31, 7, 1, 0, 0, 31, 56, 3, 2, 1, 6, 32,
		33, 10, 4, 0, 0, 33, 34, 7, 2, 0, 0, 34, 56, 3, 2, 1, 5, 35, 36, 10, 3,
		0, 0, 36, 37, 5, 2, 0, 0, 37, 56, 3, 2, 1, 4, 38, 39, 10, 2, 0, 0, 39,
		40, 5, 3, 0, 0, 40, 56, 3, 2, 1, 3, 41, 42, 10, 1, 0, 0, 42, 43, 5, 4,
		0, 0, 43, 56, 3, 2, 1, 2, 44, 45, 10, 10, 0, 0, 45, 56, 3, 4, 2, 0, 46,
		47, 10, 9, 0, 0, 47, 48, 5, 8, 0, 0, 48, 49, 5, 21, 0, 0, 49, 56, 5, 9,
		0, 0, 50, 51, 10, 8, 0, 0, 51, 52, 5, 10, 0, 0, 52, 53, 3, 8, 4, 0, 53,
		54, 5, 11, 0, 0, 54, 56, 1, 0, 0, 0, 55, 26, 1, 0, 0, 0, 55, 29, 1, 0,
		0, 0, 55, 32, 1, 0, 0, 0, 55, 35, 1, 0, 0, 0, 55, 38, 1, 0, 0, 0, 55, 41,
		1, 0, 0, 0, 55, 44, 1, 0, 0, 0, 55, 46, 1, 0, 0, 0, 55, 50, 1, 0, 0, 0,
		56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 3, 1, 0,
		0, 0, 59, 57, 1, 0, 0, 0, 60, 62, 5, 7, 0, 0, 61, 63, 5, 5, 0, 0, 62, 61,
		1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65, 1, 0, 0, 0, 64, 66, 5, 20, 0, 0,
		65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 71, 1, 0, 0, 0, 67, 68, 5,
		8, 0, 0, 68, 69, 5, 23, 0, 0, 69, 71, 5, 9, 0, 0, 70, 60, 1, 0, 0, 0, 70,
		67, 1, 0, 0, 0, 71, 5, 1, 0, 0, 0, 72, 79, 5, 19, 0, 0, 73, 74, 5, 5, 0,
		0, 74, 79, 5, 23, 0, 0, 75, 79, 5, 23, 0, 0, 76, 79, 5, 21, 0, 0, 77, 79,
		5, 22, 0, 0, 78, 72, 1, 0, 0, 0, 78, 73, 1, 0, 0, 0, 78, 75, 1, 0, 0, 0,
		78, 76, 1, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 7, 1, 0, 0, 0, 80, 85, 3, 10,
		5, 0, 81, 82, 5, 6, 0, 0, 82, 84, 3, 10, 5, 0, 83, 81, 1, 0, 0, 0, 84,
		87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 9, 1, 0, 0,
		0, 87, 85, 1, 0, 0, 0, 88, 94, 1, 0, 0, 0, 89, 94, 5, 20, 0, 0, 90, 94,
		5, 21, 0, 0, 91, 94, 5, 23, 0, 0, 92, 94, 3, 2, 1, 0, 93, 88, 1, 0, 0,
		0, 93, 89, 1, 0, 0, 0, 93, 90, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 92,
		1, 0, 0, 0, 94, 11, 1, 0, 0, 0, 9, 24, 55, 57, 62, 65, 70, 78, 85, 93,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ExpressionParserInit initializes any static state used to implement ExpressionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExpressionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExpressionParserInit() {
	staticData := &ExpressionParserStaticData
	staticData.once.Do(expressionParserInit)
}

// NewExpressionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExpressionParser(input antlr.TokenStream) *ExpressionParser {
	ExpressionParserInit()
	this := new(ExpressionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ExpressionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Expression.g4"

	return this
}

// ExpressionParser tokens.
const (
	ExpressionParserEOF                  = antlr.TokenEOF
	ExpressionParserT__0                 = 1
	ExpressionParserT__1                 = 2
	ExpressionParserT__2                 = 3
	ExpressionParserT__3                 = 4
	ExpressionParserT__4                 = 5
	ExpressionParserT__5                 = 6
	ExpressionParserDOT                  = 7
	ExpressionParserLBRACK               = 8
	ExpressionParserRBRACK               = 9
	ExpressionParserLPAREN               = 10
	ExpressionParserRPAREN               = 11
	ExpressionParserPLUS                 = 12
	ExpressionParserMINUS                = 13
	ExpressionParserMULTIPLY             = 14
	ExpressionParserDIVIDE               = 15
	ExpressionParserEQUAL                = 16
	ExpressionParserNOTEQUAL             = 17
	ExpressionParserWhiteSpaces          = 18
	ExpressionParserBooleanLiteral       = 19
	ExpressionParserIdentifier           = 20
	ExpressionParserIntegerLiteral       = 21
	ExpressionParserFloatingPointLiteral = 22
	ExpressionParserStringLiteral        = 23
	ExpressionParserWhitespace           = 24
)

// ExpressionParser rules.
const (
	ExpressionParserRULE_expression          = 0
	ExpressionParserRULE_expressionSingle    = 1
	ExpressionParserRULE_expressionMember    = 2
	ExpressionParserRULE_expressionConst     = 3
	ExpressionParserRULE_expressionArguments = 4
	ExpressionParserRULE_expressionArgument  = 5
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionSingle() IExpressionSingleContext
	EOF() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ExpressionSingle() IExpressionSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExpressionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ExpressionParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExpressionParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.expressionSingle(0)
	}
	{
		p.SetState(13)
		p.Match(ExpressionParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionSingleContext is an interface to support dynamic dispatch.
type IExpressionSingleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionSingleContext differentiates from other interfaces.
	IsExpressionSingleContext()
}

type ExpressionSingleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSingleContext() *ExpressionSingleContext {
	var p = new(ExpressionSingleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionSingle
	return p
}

func InitEmptyExpressionSingleContext(p *ExpressionSingleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionSingle
}

func (*ExpressionSingleContext) IsExpressionSingleContext() {}

func NewExpressionSingleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSingleContext {
	var p = new(ExpressionSingleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expressionSingle

	return p
}

func (s *ExpressionSingleContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSingleContext) CopyAll(ctx *ExpressionSingleContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionSingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSingleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AdditiveExpressionContext struct {
	ExpressionSingleContext
}

func NewAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) AllExpressionSingle() []IExpressionSingleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			len++
		}
	}

	tst := make([]IExpressionSingleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionSingleContext); ok {
			tst[i] = t.(IExpressionSingleContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) ExpressionSingle(i int) IExpressionSingleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *AdditiveExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ExpressionParserPLUS, 0)
}

func (s *AdditiveExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ExpressionParserMINUS, 0)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

type MemberAccessExpressionContext struct {
	ExpressionSingleContext
}

func NewMemberAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberAccessExpressionContext {
	var p = new(MemberAccessExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *MemberAccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberAccessExpressionContext) ExpressionSingle() IExpressionSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *MemberAccessExpressionContext) ExpressionMember() IExpressionMemberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionMemberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionMemberContext)
}

func (s *MemberAccessExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterMemberAccessExpression(s)
	}
}

func (s *MemberAccessExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitMemberAccessExpression(s)
	}
}

type LogicalAndExpressionContext struct {
	ExpressionSingleContext
}

func NewLogicalAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) AllExpressionSingle() []IExpressionSingleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			len++
		}
	}

	tst := make([]IExpressionSingleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionSingleContext); ok {
			tst[i] = t.(IExpressionSingleContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) ExpressionSingle(i int) IExpressionSingleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

type ConstExpressionContext struct {
	ExpressionSingleContext
}

func NewConstExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstExpressionContext {
	var p = new(ConstExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *ConstExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstExpressionContext) ExpressionConst() IExpressionConstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionConstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionConstContext)
}

func (s *ConstExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterConstExpression(s)
	}
}

func (s *ConstExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitConstExpression(s)
	}
}

type InExpressionContext struct {
	ExpressionSingleContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) AllExpressionSingle() []IExpressionSingleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			len++
		}
	}

	tst := make([]IExpressionSingleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionSingleContext); ok {
			tst[i] = t.(IExpressionSingleContext)
			i++
		}
	}

	return tst
}

func (s *InExpressionContext) ExpressionSingle(i int) IExpressionSingleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitInExpression(s)
	}
}

type LogicalOrExpressionContext struct {
	ExpressionSingleContext
}

func NewLogicalOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) AllExpressionSingle() []IExpressionSingleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			len++
		}
	}

	tst := make([]IExpressionSingleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionSingleContext); ok {
			tst[i] = t.(IExpressionSingleContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) ExpressionSingle(i int) IExpressionSingleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

type NotExpressionContext struct {
	ExpressionSingleContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) ExpressionSingle() IExpressionSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

type FunctionCallExpressionContext struct {
	ExpressionSingleContext
}

func NewFunctionCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionContext {
	var p = new(FunctionCallExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *FunctionCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionContext) ExpressionSingle() IExpressionSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *FunctionCallExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLPAREN, 0)
}

func (s *FunctionCallExpressionContext) ExpressionArguments() IExpressionArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionArgumentsContext)
}

func (s *FunctionCallExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRPAREN, 0)
}

func (s *FunctionCallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterFunctionCallExpression(s)
	}
}

func (s *FunctionCallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitFunctionCallExpression(s)
	}
}

type ParenExpressionContext struct {
	ExpressionSingleContext
}

func NewParenExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpressionContext {
	var p = new(ParenExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *ParenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLPAREN, 0)
}

func (s *ParenExpressionContext) ExpressionSingle() IExpressionSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *ParenExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRPAREN, 0)
}

func (s *ParenExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterParenExpression(s)
	}
}

func (s *ParenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitParenExpression(s)
	}
}

type EqualityExpressionContext struct {
	ExpressionSingleContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllExpressionSingle() []IExpressionSingleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			len++
		}
	}

	tst := make([]IExpressionSingleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionSingleContext); ok {
			tst[i] = t.(IExpressionSingleContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) ExpressionSingle(i int) IExpressionSingleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *EqualityExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ExpressionParserEQUAL, 0)
}

func (s *EqualityExpressionContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(ExpressionParserNOTEQUAL, 0)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

type MultiplicativeExpressionContext struct {
	ExpressionSingleContext
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllExpressionSingle() []IExpressionSingleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			len++
		}
	}

	tst := make([]IExpressionSingleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionSingleContext); ok {
			tst[i] = t.(IExpressionSingleContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) ExpressionSingle(i int) IExpressionSingleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *MultiplicativeExpressionContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(ExpressionParserMULTIPLY, 0)
}

func (s *MultiplicativeExpressionContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(ExpressionParserDIVIDE, 0)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

type ArrayAccessExpressionContext struct {
	ExpressionSingleContext
}

func NewArrayAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayAccessExpressionContext {
	var p = new(ArrayAccessExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *ArrayAccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAccessExpressionContext) ExpressionSingle() IExpressionSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *ArrayAccessExpressionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLBRACK, 0)
}

func (s *ArrayAccessExpressionContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIntegerLiteral, 0)
}

func (s *ArrayAccessExpressionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRBRACK, 0)
}

func (s *ArrayAccessExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterArrayAccessExpression(s)
	}
}

func (s *ArrayAccessExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitArrayAccessExpression(s)
	}
}

type IdentifierAccessExpressionContext struct {
	ExpressionSingleContext
}

func NewIdentifierAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierAccessExpressionContext {
	var p = new(IdentifierAccessExpressionContext)

	InitEmptyExpressionSingleContext(&p.ExpressionSingleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSingleContext))

	return p
}

func (s *IdentifierAccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierAccessExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIdentifier, 0)
}

func (s *IdentifierAccessExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterIdentifierAccessExpression(s)
	}
}

func (s *IdentifierAccessExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitIdentifierAccessExpression(s)
	}
}

func (p *ExpressionParser) ExpressionSingle() (localctx IExpressionSingleContext) {
	return p.expressionSingle(0)
}

func (p *ExpressionParser) expressionSingle(_p int) (localctx IExpressionSingleContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionSingleContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionSingleContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ExpressionParserRULE_expressionSingle, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ExpressionParserT__4, ExpressionParserBooleanLiteral, ExpressionParserIntegerLiteral, ExpressionParserFloatingPointLiteral, ExpressionParserStringLiteral:
		localctx = NewConstExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(16)
			p.ExpressionConst()
		}

	case ExpressionParserIdentifier:
		localctx = NewIdentifierAccessExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)
			p.Match(ExpressionParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ExpressionParserLPAREN:
		localctx = NewParenExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(18)
			p.Match(ExpressionParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.expressionSingle(0)
		}
		{
			p.SetState(20)
			p.Match(ExpressionParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ExpressionParserT__0:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(22)
			p.Match(ExpressionParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(23)
			p.expressionSingle(7)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAdditiveExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(26)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(27)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ExpressionParserPLUS || _la == ExpressionParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(28)
					p.expressionSingle(7)
				}

			case 2:
				localctx = NewMultiplicativeExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(30)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ExpressionParserMULTIPLY || _la == ExpressionParserDIVIDE) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(31)
					p.expressionSingle(6)
				}

			case 3:
				localctx = NewEqualityExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(33)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ExpressionParserEQUAL || _la == ExpressionParserNOTEQUAL) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(34)
					p.expressionSingle(5)
				}

			case 4:
				localctx = NewInExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(36)
					p.Match(ExpressionParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(37)
					p.expressionSingle(4)
				}

			case 5:
				localctx = NewLogicalAndExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(39)
					p.Match(ExpressionParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(40)
					p.expressionSingle(3)
				}

			case 6:
				localctx = NewLogicalOrExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(42)
					p.Match(ExpressionParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(43)
					p.expressionSingle(2)
				}

			case 7:
				localctx = NewMemberAccessExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(45)
					p.ExpressionMember()
				}

			case 8:
				localctx = NewArrayAccessExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(46)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(47)
					p.Match(ExpressionParserLBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(48)
					p.Match(ExpressionParserIntegerLiteral)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(49)
					p.Match(ExpressionParserRBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 9:
				localctx = NewFunctionCallExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(50)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(51)
					p.Match(ExpressionParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(52)
					p.ExpressionArguments()
				}
				{
					p.SetState(53)
					p.Match(ExpressionParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionMemberContext is an interface to support dynamic dispatch.
type IExpressionMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	RBRACK() antlr.TerminalNode

	// IsExpressionMemberContext differentiates from other interfaces.
	IsExpressionMemberContext()
}

type ExpressionMemberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionMemberContext() *ExpressionMemberContext {
	var p = new(ExpressionMemberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionMember
	return p
}

func InitEmptyExpressionMemberContext(p *ExpressionMemberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionMember
}

func (*ExpressionMemberContext) IsExpressionMemberContext() {}

func NewExpressionMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionMemberContext {
	var p = new(ExpressionMemberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expressionMember

	return p
}

func (s *ExpressionMemberContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionMemberContext) DOT() antlr.TerminalNode {
	return s.GetToken(ExpressionParserDOT, 0)
}

func (s *ExpressionMemberContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIdentifier, 0)
}

func (s *ExpressionMemberContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLBRACK, 0)
}

func (s *ExpressionMemberContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserStringLiteral, 0)
}

func (s *ExpressionMemberContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRBRACK, 0)
}

func (s *ExpressionMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionMemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterExpressionMember(s)
	}
}

func (s *ExpressionMemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitExpressionMember(s)
	}
}

func (p *ExpressionParser) ExpressionMember() (localctx IExpressionMemberContext) {
	localctx = NewExpressionMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExpressionParserRULE_expressionMember)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ExpressionParserDOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Match(ExpressionParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(61)
				p.Match(ExpressionParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(64)
				p.Match(ExpressionParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case ExpressionParserLBRACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Match(ExpressionParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(ExpressionParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(ExpressionParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionConstContext is an interface to support dynamic dispatch.
type IExpressionConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionConstContext differentiates from other interfaces.
	IsExpressionConstContext()
}

type ExpressionConstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionConstContext() *ExpressionConstContext {
	var p = new(ExpressionConstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionConst
	return p
}

func InitEmptyExpressionConstContext(p *ExpressionConstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionConst
}

func (*ExpressionConstContext) IsExpressionConstContext() {}

func NewExpressionConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionConstContext {
	var p = new(ExpressionConstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expressionConst

	return p
}

func (s *ExpressionConstContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionConstContext) CopyAll(ctx *ExpressionConstContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryStringLiteralContext struct {
	ExpressionConstContext
}

func NewBinaryStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryStringLiteralContext {
	var p = new(BinaryStringLiteralContext)

	InitEmptyExpressionConstContext(&p.ExpressionConstContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionConstContext))

	return p
}

func (s *BinaryStringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryStringLiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserStringLiteral, 0)
}

func (s *BinaryStringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterBinaryStringLiteral(s)
	}
}

func (s *BinaryStringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitBinaryStringLiteral(s)
	}
}

type StringLiteralContext struct {
	ExpressionConstContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyExpressionConstContext(&p.ExpressionConstContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionConstContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserStringLiteral, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

type FloatingPointLiteralContext struct {
	ExpressionConstContext
}

func NewFloatingPointLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatingPointLiteralContext {
	var p = new(FloatingPointLiteralContext)

	InitEmptyExpressionConstContext(&p.ExpressionConstContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionConstContext))

	return p
}

func (s *FloatingPointLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatingPointLiteralContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserFloatingPointLiteral, 0)
}

func (s *FloatingPointLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterFloatingPointLiteral(s)
	}
}

func (s *FloatingPointLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitFloatingPointLiteral(s)
	}
}

type BooleanLiteralContext struct {
	ExpressionConstContext
}

func NewBooleanLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	InitEmptyExpressionConstContext(&p.ExpressionConstContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionConstContext))

	return p
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserBooleanLiteral, 0)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

type IntegerLiteralContext struct {
	ExpressionConstContext
}

func NewIntegerLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	InitEmptyExpressionConstContext(&p.ExpressionConstContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionConstContext))

	return p
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIntegerLiteral, 0)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (p *ExpressionParser) ExpressionConst() (localctx IExpressionConstContext) {
	localctx = NewExpressionConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExpressionParserRULE_expressionConst)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ExpressionParserBooleanLiteral:
		localctx = NewBooleanLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Match(ExpressionParserBooleanLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ExpressionParserT__4:
		localctx = NewBinaryStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Match(ExpressionParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(ExpressionParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ExpressionParserStringLiteral:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)
			p.Match(ExpressionParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ExpressionParserIntegerLiteral:
		localctx = NewIntegerLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(76)
			p.Match(ExpressionParserIntegerLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ExpressionParserFloatingPointLiteral:
		localctx = NewFloatingPointLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(77)
			p.Match(ExpressionParserFloatingPointLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionArgumentsContext is an interface to support dynamic dispatch.
type IExpressionArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpressionArgument() []IExpressionArgumentContext
	ExpressionArgument(i int) IExpressionArgumentContext

	// IsExpressionArgumentsContext differentiates from other interfaces.
	IsExpressionArgumentsContext()
}

type ExpressionArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionArgumentsContext() *ExpressionArgumentsContext {
	var p = new(ExpressionArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionArguments
	return p
}

func InitEmptyExpressionArgumentsContext(p *ExpressionArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionArguments
}

func (*ExpressionArgumentsContext) IsExpressionArgumentsContext() {}

func NewExpressionArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionArgumentsContext {
	var p = new(ExpressionArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expressionArguments

	return p
}

func (s *ExpressionArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionArgumentsContext) AllExpressionArgument() []IExpressionArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionArgumentContext); ok {
			len++
		}
	}

	tst := make([]IExpressionArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionArgumentContext); ok {
			tst[i] = t.(IExpressionArgumentContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionArgumentsContext) ExpressionArgument(i int) IExpressionArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionArgumentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionArgumentContext)
}

func (s *ExpressionArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterExpressionArguments(s)
	}
}

func (s *ExpressionArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitExpressionArguments(s)
	}
}

func (p *ExpressionParser) ExpressionArguments() (localctx IExpressionArgumentsContext) {
	localctx = NewExpressionArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ExpressionParserRULE_expressionArguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.ExpressionArgument()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ExpressionParserT__5 {
		{
			p.SetState(81)
			p.Match(ExpressionParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.ExpressionArgument()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionArgumentContext is an interface to support dynamic dispatch.
type IExpressionArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	IntegerLiteral() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	ExpressionSingle() IExpressionSingleContext

	// IsExpressionArgumentContext differentiates from other interfaces.
	IsExpressionArgumentContext()
}

type ExpressionArgumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionArgumentContext() *ExpressionArgumentContext {
	var p = new(ExpressionArgumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionArgument
	return p
}

func InitEmptyExpressionArgumentContext(p *ExpressionArgumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionArgument
}

func (*ExpressionArgumentContext) IsExpressionArgumentContext() {}

func NewExpressionArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionArgumentContext {
	var p = new(ExpressionArgumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expressionArgument

	return p
}

func (s *ExpressionArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionArgumentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIdentifier, 0)
}

func (s *ExpressionArgumentContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIntegerLiteral, 0)
}

func (s *ExpressionArgumentContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserStringLiteral, 0)
}

func (s *ExpressionArgumentContext) ExpressionSingle() IExpressionSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *ExpressionArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterExpressionArgument(s)
	}
}

func (s *ExpressionArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitExpressionArgument(s)
	}
}

func (p *ExpressionParser) ExpressionArgument() (localctx IExpressionArgumentContext) {
	localctx = NewExpressionArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ExpressionParserRULE_expressionArgument)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Match(ExpressionParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Match(ExpressionParserIntegerLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.Match(ExpressionParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.expressionSingle(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ExpressionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionSingleContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionSingleContext)
		}
		return p.ExpressionSingle_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ExpressionParser) ExpressionSingle_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
