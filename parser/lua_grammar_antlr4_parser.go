// Code generated from lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // lua_grammar_antlr4

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

type lua_grammar_antlr4Parser struct {
	*antlr.BaseParser
}

var Lua_grammar_antlr4ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lua_grammar_antlr4ParserInit() {
	staticData := &Lua_grammar_antlr4ParserStaticData
	staticData.LiteralNames = []string{
		"", "'local'", "'='", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'true'",
		"'false'", "'nil'", "','", "'function'", "'end'", "'if'", "'then'",
		"'elseif'", "'else'", "'for'", "'do'", "'while'", "'{'", "'}'", "'['",
		"']'", "'--[['", "'--]]'", "']]'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "NUMBER", "STRING",
		"LETTER", "DIGIT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "assignment", "expression", "literal", "function_call",
		"function_declaration", "block", "if_statement", "for_statement", "while_statement",
		"table", "field", "table_access", "multi_line_comment", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 33, 198, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 5, 0, 35, 8, 0, 10, 0, 12, 0, 38, 9, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 46, 8, 1, 1, 2, 3, 2, 49, 8, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		65, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 73, 8, 3, 10, 3, 12,
		3, 76, 9, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 85, 8, 5,
		10, 5, 12, 5, 88, 9, 5, 3, 5, 90, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 5, 6, 100, 8, 6, 10, 6, 12, 6, 103, 9, 6, 3, 6, 105, 8,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 5, 7, 112, 8, 7, 10, 7, 12, 7, 115, 9,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 126, 8,
		8, 10, 8, 12, 8, 129, 9, 8, 1, 8, 1, 8, 3, 8, 133, 8, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 145, 8, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 5, 11, 161, 8, 11, 10, 11, 12, 11, 164, 9, 11, 3, 11, 166,
		8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 175, 8,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 184, 8, 14,
		10, 14, 12, 14, 187, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 193, 8,
		15, 10, 15, 12, 15, 196, 9, 15, 1, 15, 0, 1, 6, 16, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 5, 1, 0, 5, 6, 1, 0, 7, 8, 2,
		0, 9, 11, 29, 30, 1, 0, 27, 27, 1, 0, 31, 32, 209, 0, 36, 1, 0, 0, 0, 2,
		45, 1, 0, 0, 0, 4, 48, 1, 0, 0, 0, 6, 64, 1, 0, 0, 0, 8, 77, 1, 0, 0, 0,
		10, 79, 1, 0, 0, 0, 12, 93, 1, 0, 0, 0, 14, 113, 1, 0, 0, 0, 16, 116, 1,
		0, 0, 0, 18, 136, 1, 0, 0, 0, 20, 150, 1, 0, 0, 0, 22, 156, 1, 0, 0, 0,
		24, 174, 1, 0, 0, 0, 26, 176, 1, 0, 0, 0, 28, 181, 1, 0, 0, 0, 30, 190,
		1, 0, 0, 0, 32, 35, 3, 2, 1, 0, 33, 35, 3, 28, 14, 0, 34, 32, 1, 0, 0,
		0, 34, 33, 1, 0, 0, 0, 35, 38, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37,
		1, 0, 0, 0, 37, 1, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 39, 46, 3, 4, 2, 0,
		40, 46, 3, 6, 3, 0, 41, 46, 3, 16, 8, 0, 42, 46, 3, 18, 9, 0, 43, 46, 3,
		20, 10, 0, 44, 46, 3, 12, 6, 0, 45, 39, 1, 0, 0, 0, 45, 40, 1, 0, 0, 0,
		45, 41, 1, 0, 0, 0, 45, 42, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 44, 1,
		0, 0, 0, 46, 3, 1, 0, 0, 0, 47, 49, 5, 1, 0, 0, 48, 47, 1, 0, 0, 0, 48,
		49, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 3, 30, 15, 0, 51, 52, 5, 2,
		0, 0, 52, 53, 3, 6, 3, 0, 53, 5, 1, 0, 0, 0, 54, 55, 6, 3, -1, 0, 55, 65,
		3, 8, 4, 0, 56, 65, 3, 30, 15, 0, 57, 65, 3, 10, 5, 0, 58, 65, 3, 22, 11,
		0, 59, 65, 3, 26, 13, 0, 60, 61, 5, 3, 0, 0, 61, 62, 3, 6, 3, 0, 62, 63,
		5, 4, 0, 0, 63, 65, 1, 0, 0, 0, 64, 54, 1, 0, 0, 0, 64, 56, 1, 0, 0, 0,
		64, 57, 1, 0, 0, 0, 64, 58, 1, 0, 0, 0, 64, 59, 1, 0, 0, 0, 64, 60, 1,
		0, 0, 0, 65, 74, 1, 0, 0, 0, 66, 67, 10, 2, 0, 0, 67, 68, 7, 0, 0, 0, 68,
		73, 3, 6, 3, 3, 69, 70, 10, 1, 0, 0, 70, 71, 7, 1, 0, 0, 71, 73, 3, 6,
		3, 2, 72, 66, 1, 0, 0, 0, 72, 69, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72,
		1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 7, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0,
		77, 78, 7, 2, 0, 0, 78, 9, 1, 0, 0, 0, 79, 80, 3, 30, 15, 0, 80, 89, 5,
		3, 0, 0, 81, 86, 3, 6, 3, 0, 82, 83, 5, 12, 0, 0, 83, 85, 3, 6, 3, 0, 84,
		82, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0,
		0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 81, 1, 0, 0, 0, 89, 90,
		1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 5, 4, 0, 0, 92, 11, 1, 0, 0, 0,
		93, 94, 5, 13, 0, 0, 94, 95, 3, 30, 15, 0, 95, 104, 5, 3, 0, 0, 96, 101,
		3, 30, 15, 0, 97, 98, 5, 12, 0, 0, 98, 100, 3, 30, 15, 0, 99, 97, 1, 0,
		0, 0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0,
		102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 96, 1, 0, 0, 0, 104, 105,
		1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 5, 4, 0, 0, 107, 108, 3, 14,
		7, 0, 108, 109, 5, 14, 0, 0, 109, 13, 1, 0, 0, 0, 110, 112, 3, 2, 1, 0,
		111, 110, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113,
		114, 1, 0, 0, 0, 114, 15, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 117, 5,
		15, 0, 0, 117, 118, 3, 6, 3, 0, 118, 119, 5, 16, 0, 0, 119, 127, 3, 14,
		7, 0, 120, 121, 5, 17, 0, 0, 121, 122, 3, 6, 3, 0, 122, 123, 5, 16, 0,
		0, 123, 124, 3, 14, 7, 0, 124, 126, 1, 0, 0, 0, 125, 120, 1, 0, 0, 0, 126,
		129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 132,
		1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 131, 5, 18, 0, 0, 131, 133, 3, 14,
		7, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0,
		134, 135, 5, 14, 0, 0, 135, 17, 1, 0, 0, 0, 136, 137, 5, 19, 0, 0, 137,
		138, 3, 30, 15, 0, 138, 139, 5, 2, 0, 0, 139, 140, 3, 6, 3, 0, 140, 141,
		5, 12, 0, 0, 141, 144, 3, 6, 3, 0, 142, 143, 5, 12, 0, 0, 143, 145, 3,
		6, 3, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0, 0,
		0, 146, 147, 5, 20, 0, 0, 147, 148, 3, 14, 7, 0, 148, 149, 5, 14, 0, 0,
		149, 19, 1, 0, 0, 0, 150, 151, 5, 21, 0, 0, 151, 152, 3, 6, 3, 0, 152,
		153, 5, 20, 0, 0, 153, 154, 3, 14, 7, 0, 154, 155, 5, 14, 0, 0, 155, 21,
		1, 0, 0, 0, 156, 165, 5, 22, 0, 0, 157, 162, 3, 24, 12, 0, 158, 159, 5,
		12, 0, 0, 159, 161, 3, 24, 12, 0, 160, 158, 1, 0, 0, 0, 161, 164, 1, 0,
		0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 166, 1, 0, 0, 0,
		164, 162, 1, 0, 0, 0, 165, 157, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166,
		167, 1, 0, 0, 0, 167, 168, 5, 23, 0, 0, 168, 23, 1, 0, 0, 0, 169, 170,
		3, 30, 15, 0, 170, 171, 5, 2, 0, 0, 171, 172, 3, 6, 3, 0, 172, 175, 1,
		0, 0, 0, 173, 175, 3, 6, 3, 0, 174, 169, 1, 0, 0, 0, 174, 173, 1, 0, 0,
		0, 175, 25, 1, 0, 0, 0, 176, 177, 3, 30, 15, 0, 177, 178, 5, 24, 0, 0,
		178, 179, 3, 6, 3, 0, 179, 180, 5, 25, 0, 0, 180, 27, 1, 0, 0, 0, 181,
		185, 5, 26, 0, 0, 182, 184, 8, 3, 0, 0, 183, 182, 1, 0, 0, 0, 184, 187,
		1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 188, 1, 0,
		0, 0, 187, 185, 1, 0, 0, 0, 188, 189, 5, 28, 0, 0, 189, 29, 1, 0, 0, 0,
		190, 194, 5, 31, 0, 0, 191, 193, 7, 4, 0, 0, 192, 191, 1, 0, 0, 0, 193,
		196, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 31, 1,
		0, 0, 0, 196, 194, 1, 0, 0, 0, 20, 34, 36, 45, 48, 64, 72, 74, 86, 89,
		101, 104, 113, 127, 132, 144, 162, 165, 174, 185, 194,
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

// lua_grammar_antlr4ParserInit initializes any static state used to implement lua_grammar_antlr4Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// Newlua_grammar_antlr4Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Lua_grammar_antlr4ParserInit() {
	staticData := &Lua_grammar_antlr4ParserStaticData
	staticData.once.Do(lua_grammar_antlr4ParserInit)
}

// Newlua_grammar_antlr4Parser produces a new parser instance for the optional input antlr.TokenStream.
func Newlua_grammar_antlr4Parser(input antlr.TokenStream) *lua_grammar_antlr4Parser {
	Lua_grammar_antlr4ParserInit()
	this := new(lua_grammar_antlr4Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Lua_grammar_antlr4ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "lua_grammar_antlr4.g4"

	return this
}

// lua_grammar_antlr4Parser tokens.
const (
	lua_grammar_antlr4ParserEOF    = antlr.TokenEOF
	lua_grammar_antlr4ParserT__0   = 1
	lua_grammar_antlr4ParserT__1   = 2
	lua_grammar_antlr4ParserT__2   = 3
	lua_grammar_antlr4ParserT__3   = 4
	lua_grammar_antlr4ParserT__4   = 5
	lua_grammar_antlr4ParserT__5   = 6
	lua_grammar_antlr4ParserT__6   = 7
	lua_grammar_antlr4ParserT__7   = 8
	lua_grammar_antlr4ParserT__8   = 9
	lua_grammar_antlr4ParserT__9   = 10
	lua_grammar_antlr4ParserT__10  = 11
	lua_grammar_antlr4ParserT__11  = 12
	lua_grammar_antlr4ParserT__12  = 13
	lua_grammar_antlr4ParserT__13  = 14
	lua_grammar_antlr4ParserT__14  = 15
	lua_grammar_antlr4ParserT__15  = 16
	lua_grammar_antlr4ParserT__16  = 17
	lua_grammar_antlr4ParserT__17  = 18
	lua_grammar_antlr4ParserT__18  = 19
	lua_grammar_antlr4ParserT__19  = 20
	lua_grammar_antlr4ParserT__20  = 21
	lua_grammar_antlr4ParserT__21  = 22
	lua_grammar_antlr4ParserT__22  = 23
	lua_grammar_antlr4ParserT__23  = 24
	lua_grammar_antlr4ParserT__24  = 25
	lua_grammar_antlr4ParserT__25  = 26
	lua_grammar_antlr4ParserT__26  = 27
	lua_grammar_antlr4ParserT__27  = 28
	lua_grammar_antlr4ParserNUMBER = 29
	lua_grammar_antlr4ParserSTRING = 30
	lua_grammar_antlr4ParserLETTER = 31
	lua_grammar_antlr4ParserDIGIT  = 32
	lua_grammar_antlr4ParserWS     = 33
)

// lua_grammar_antlr4Parser rules.
const (
	lua_grammar_antlr4ParserRULE_program              = 0
	lua_grammar_antlr4ParserRULE_statement            = 1
	lua_grammar_antlr4ParserRULE_assignment           = 2
	lua_grammar_antlr4ParserRULE_expression           = 3
	lua_grammar_antlr4ParserRULE_literal              = 4
	lua_grammar_antlr4ParserRULE_function_call        = 5
	lua_grammar_antlr4ParserRULE_function_declaration = 6
	lua_grammar_antlr4ParserRULE_block                = 7
	lua_grammar_antlr4ParserRULE_if_statement         = 8
	lua_grammar_antlr4ParserRULE_for_statement        = 9
	lua_grammar_antlr4ParserRULE_while_statement      = 10
	lua_grammar_antlr4ParserRULE_table                = 11
	lua_grammar_antlr4ParserRULE_field                = 12
	lua_grammar_antlr4ParserRULE_table_access         = 13
	lua_grammar_antlr4ParserRULE_multi_line_comment   = 14
	lua_grammar_antlr4ParserRULE_identifier           = 15
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllMulti_line_comment() []IMulti_line_commentContext
	Multi_line_comment(i int) IMulti_line_commentContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *ProgramContext) AllMulti_line_comment() []IMulti_line_commentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMulti_line_commentContext); ok {
			len++
		}
	}

	tst := make([]IMulti_line_commentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMulti_line_commentContext); ok {
			tst[i] = t.(IMulti_line_commentContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Multi_line_comment(i int) IMulti_line_commentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulti_line_commentContext); ok {
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

	return t.(IMulti_line_commentContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, lua_grammar_antlr4ParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3832065546) != 0 {
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case lua_grammar_antlr4ParserT__0, lua_grammar_antlr4ParserT__2, lua_grammar_antlr4ParserT__8, lua_grammar_antlr4ParserT__9, lua_grammar_antlr4ParserT__10, lua_grammar_antlr4ParserT__12, lua_grammar_antlr4ParserT__14, lua_grammar_antlr4ParserT__18, lua_grammar_antlr4ParserT__20, lua_grammar_antlr4ParserT__21, lua_grammar_antlr4ParserNUMBER, lua_grammar_antlr4ParserSTRING, lua_grammar_antlr4ParserLETTER:
			{
				p.SetState(32)
				p.Statement()
			}

		case lua_grammar_antlr4ParserT__25:
			{
				p.SetState(33)
				p.Multi_line_comment()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(38)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	Expression() IExpressionContext
	If_statement() IIf_statementContext
	For_statement() IFor_statementContext
	While_statement() IWhile_statementContext
	Function_declaration() IFunction_declarationContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *StatementContext) For_statement() IFor_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_statementContext)
}

func (s *StatementContext) While_statement() IWhile_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_statementContext)
}

func (s *StatementContext) Function_declaration() IFunction_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_declarationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, lua_grammar_antlr4ParserRULE_statement)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(41)
			p.If_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(42)
			p.For_statement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(43)
			p.While_statement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(44)
			p.Function_declaration()
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Expression() IExpressionContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, lua_grammar_antlr4ParserRULE_assignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserT__0 {
		{
			p.SetState(47)
			p.Match(lua_grammar_antlr4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(50)
		p.Identifier()
	}
	{
		p.SetState(51)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.expression(0)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArith_operator returns the arith_operator token.
	GetArith_operator() antlr.Token

	// SetArith_operator sets the arith_operator token.
	SetArith_operator(antlr.Token)

	// Getter signatures
	Literal() ILiteralContext
	Identifier() IIdentifierContext
	Function_call() IFunction_callContext
	Table() ITableContext
	Table_access() ITable_accessContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	arith_operator antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetArith_operator() antlr.Token { return s.arith_operator }

func (s *ExpressionContext) SetArith_operator(v antlr.Token) { s.arith_operator = v }

func (s *ExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExpressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *ExpressionContext) Table() ITableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableContext)
}

func (s *ExpressionContext) Table_access() ITable_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_accessContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *lua_grammar_antlr4Parser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, lua_grammar_antlr4ParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(55)
			p.Literal()
		}

	case 2:
		{
			p.SetState(56)
			p.Identifier()
		}

	case 3:
		{
			p.SetState(57)
			p.Function_call()
		}

	case 4:
		{
			p.SetState(58)
			p.Table()
		}

	case 5:
		{
			p.SetState(59)
			p.Table_access()
		}

	case 6:
		{
			p.SetState(60)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.expression(0)
		}
		{
			p.SetState(62)
			p.Match(lua_grammar_antlr4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_expression)
				p.SetState(66)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(67)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).arith_operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == lua_grammar_antlr4ParserT__4 || _la == lua_grammar_antlr4ParserT__5) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).arith_operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(68)
					p.expression(3)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_expression)
				p.SetState(69)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(70)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).arith_operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == lua_grammar_antlr4ParserT__6 || _la == lua_grammar_antlr4ParserT__7) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).arith_operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(71)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserNUMBER, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserSTRING, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, lua_grammar_antlr4ParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1610616320) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Function_callContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Function_callContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (s *Function_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitFunction_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, lua_grammar_antlr4ParserRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Identifier()
	}
	{
		p.SetState(80)
		p.Match(lua_grammar_antlr4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3762294280) != 0 {
		{
			p.SetState(81)
			p.expression(0)
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__11 {
			{
				p.SetState(82)
				p.Match(lua_grammar_antlr4ParserT__11)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(83)
				p.expression(0)
			}

			p.SetState(88)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(91)
		p.Match(lua_grammar_antlr4ParserT__3)
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

// IFunction_declarationContext is an interface to support dynamic dispatch.
type IFunction_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	Block() IBlockContext

	// IsFunction_declarationContext differentiates from other interfaces.
	IsFunction_declarationContext()
}

type Function_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_declarationContext() *Function_declarationContext {
	var p = new(Function_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_declaration
	return p
}

func InitEmptyFunction_declarationContext(p *Function_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_declaration
}

func (*Function_declarationContext) IsFunction_declarationContext() {}

func NewFunction_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declarationContext {
	var p = new(Function_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_declaration

	return p
}

func (s *Function_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declarationContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *Function_declarationContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *Function_declarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Function_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterFunction_declaration(s)
	}
}

func (s *Function_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitFunction_declaration(s)
	}
}

func (s *Function_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitFunction_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Function_declaration() (localctx IFunction_declarationContext) {
	localctx = NewFunction_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, lua_grammar_antlr4ParserRULE_function_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(lua_grammar_antlr4ParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Identifier()
	}
	{
		p.SetState(95)
		p.Match(lua_grammar_antlr4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserLETTER {
		{
			p.SetState(96)
			p.Identifier()
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__11 {
			{
				p.SetState(97)
				p.Match(lua_grammar_antlr4ParserT__11)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(98)
				p.Identifier()
			}

			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(106)
		p.Match(lua_grammar_antlr4ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Block()
	}
	{
		p.SetState(108)
		p.Match(lua_grammar_antlr4ParserT__13)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, lua_grammar_antlr4ParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3764956682) != 0 {
		{
			p.SetState(110)
			p.Statement()
		}

		p.SetState(115)
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

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *If_statementContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (s *If_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitIf_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, lua_grammar_antlr4ParserRULE_if_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(lua_grammar_antlr4ParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.expression(0)
	}
	{
		p.SetState(118)
		p.Match(lua_grammar_antlr4ParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Block()
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == lua_grammar_antlr4ParserT__16 {
		{
			p.SetState(120)
			p.Match(lua_grammar_antlr4ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.expression(0)
		}
		{
			p.SetState(122)
			p.Match(lua_grammar_antlr4ParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Block()
		}

		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserT__17 {
		{
			p.SetState(130)
			p.Match(lua_grammar_antlr4ParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Block()
		}

	}
	{
		p.SetState(134)
		p.Match(lua_grammar_antlr4ParserT__13)
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

// IFor_statementContext is an interface to support dynamic dispatch.
type IFor_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	Block() IBlockContext

	// IsFor_statementContext differentiates from other interfaces.
	IsFor_statementContext()
}

type For_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_statementContext() *For_statementContext {
	var p = new(For_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_for_statement
	return p
}

func InitEmptyFor_statementContext(p *For_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_for_statement
}

func (*For_statementContext) IsFor_statementContext() {}

func NewFor_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_statementContext {
	var p = new(For_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_for_statement

	return p
}

func (s *For_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *For_statementContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *For_statementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *For_statementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *For_statementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *For_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterFor_statement(s)
	}
}

func (s *For_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitFor_statement(s)
	}
}

func (s *For_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitFor_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) For_statement() (localctx IFor_statementContext) {
	localctx = NewFor_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, lua_grammar_antlr4ParserRULE_for_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(lua_grammar_antlr4ParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Identifier()
	}
	{
		p.SetState(138)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.expression(0)
	}
	{
		p.SetState(140)
		p.Match(lua_grammar_antlr4ParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.expression(0)
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserT__11 {
		{
			p.SetState(142)
			p.Match(lua_grammar_antlr4ParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.expression(0)
		}

	}
	{
		p.SetState(146)
		p.Match(lua_grammar_antlr4ParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Block()
	}
	{
		p.SetState(148)
		p.Match(lua_grammar_antlr4ParserT__13)
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

// IWhile_statementContext is an interface to support dynamic dispatch.
type IWhile_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Block() IBlockContext

	// IsWhile_statementContext differentiates from other interfaces.
	IsWhile_statementContext()
}

type While_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_statementContext() *While_statementContext {
	var p = new(While_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_while_statement
	return p
}

func InitEmptyWhile_statementContext(p *While_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_while_statement
}

func (*While_statementContext) IsWhile_statementContext() {}

func NewWhile_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_statementContext {
	var p = new(While_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_while_statement

	return p
}

func (s *While_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *While_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *While_statementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *While_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterWhile_statement(s)
	}
}

func (s *While_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitWhile_statement(s)
	}
}

func (s *While_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitWhile_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) While_statement() (localctx IWhile_statementContext) {
	localctx = NewWhile_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, lua_grammar_antlr4ParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(lua_grammar_antlr4ParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.expression(0)
	}
	{
		p.SetState(152)
		p.Match(lua_grammar_antlr4ParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Block()
	}
	{
		p.SetState(154)
		p.Match(lua_grammar_antlr4ParserT__13)
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

// ITableContext is an interface to support dynamic dispatch.
type ITableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsTableContext differentiates from other interfaces.
	IsTableContext()
}

type TableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableContext() *TableContext {
	var p = new(TableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table
	return p
}

func InitEmptyTableContext(p *TableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table
}

func (*TableContext) IsTableContext() {}

func NewTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableContext {
	var p = new(TableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table

	return p
}

func (s *TableContext) GetParser() antlr.Parser { return s.parser }

func (s *TableContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *TableContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *TableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterTable(s)
	}
}

func (s *TableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitTable(s)
	}
}

func (s *TableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitTable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Table() (localctx ITableContext) {
	localctx = NewTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, lua_grammar_antlr4ParserRULE_table)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(lua_grammar_antlr4ParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3762294280) != 0 {
		{
			p.SetState(157)
			p.Field()
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__11 {
			{
				p.SetState(158)
				p.Match(lua_grammar_antlr4ParserT__11)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(159)
				p.Field()
			}

			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(167)
		p.Match(lua_grammar_antlr4ParserT__22)
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

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Expression() IExpressionContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FieldContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, lua_grammar_antlr4ParserRULE_field)
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Identifier()
		}
		{
			p.SetState(170)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(173)
			p.expression(0)
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

// ITable_accessContext is an interface to support dynamic dispatch.
type ITable_accessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Expression() IExpressionContext

	// IsTable_accessContext differentiates from other interfaces.
	IsTable_accessContext()
}

type Table_accessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_accessContext() *Table_accessContext {
	var p = new(Table_accessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table_access
	return p
}

func InitEmptyTable_accessContext(p *Table_accessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table_access
}

func (*Table_accessContext) IsTable_accessContext() {}

func NewTable_accessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_accessContext {
	var p = new(Table_accessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table_access

	return p
}

func (s *Table_accessContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_accessContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Table_accessContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Table_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_accessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_accessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterTable_access(s)
	}
}

func (s *Table_accessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitTable_access(s)
	}
}

func (s *Table_accessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitTable_access(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Table_access() (localctx ITable_accessContext) {
	localctx = NewTable_accessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, lua_grammar_antlr4ParserRULE_table_access)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Identifier()
	}
	{
		p.SetState(177)
		p.Match(lua_grammar_antlr4ParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.expression(0)
	}
	{
		p.SetState(179)
		p.Match(lua_grammar_antlr4ParserT__24)
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

// IMulti_line_commentContext is an interface to support dynamic dispatch.
type IMulti_line_commentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMulti_line_commentContext differentiates from other interfaces.
	IsMulti_line_commentContext()
}

type Multi_line_commentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulti_line_commentContext() *Multi_line_commentContext {
	var p = new(Multi_line_commentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_multi_line_comment
	return p
}

func InitEmptyMulti_line_commentContext(p *Multi_line_commentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_multi_line_comment
}

func (*Multi_line_commentContext) IsMulti_line_commentContext() {}

func NewMulti_line_commentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Multi_line_commentContext {
	var p = new(Multi_line_commentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_multi_line_comment

	return p
}

func (s *Multi_line_commentContext) GetParser() antlr.Parser { return s.parser }
func (s *Multi_line_commentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Multi_line_commentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Multi_line_commentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterMulti_line_comment(s)
	}
}

func (s *Multi_line_commentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitMulti_line_comment(s)
	}
}

func (s *Multi_line_commentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitMulti_line_comment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Multi_line_comment() (localctx IMulti_line_commentContext) {
	localctx = NewMulti_line_commentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, lua_grammar_antlr4ParserRULE_multi_line_comment)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(lua_grammar_antlr4ParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(182)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || _la == lua_grammar_antlr4ParserT__26 {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.Match(lua_grammar_antlr4ParserT__27)
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLETTER() []antlr.TerminalNode
	LETTER(i int) antlr.TerminalNode
	AllDIGIT() []antlr.TerminalNode
	DIGIT(i int) antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(lua_grammar_antlr4ParserLETTER)
}

func (s *IdentifierContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserLETTER, i)
}

func (s *IdentifierContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(lua_grammar_antlr4ParserDIGIT)
}

func (s *IdentifierContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserDIGIT, i)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, lua_grammar_antlr4ParserRULE_identifier)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(lua_grammar_antlr4ParserLETTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(191)
				_la = p.GetTokenStream().LA(1)

				if !(_la == lua_grammar_antlr4ParserLETTER || _la == lua_grammar_antlr4ParserDIGIT) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *lua_grammar_antlr4Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *lua_grammar_antlr4Parser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
