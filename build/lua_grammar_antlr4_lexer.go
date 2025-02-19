// Code generated from lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type lua_grammar_antlr4Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var Lua_grammar_antlr4LexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lua_grammar_antlr4lexerLexerInit() {
	staticData := &Lua_grammar_antlr4LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "NUMBER", "STRING", "ESC", "LETTER", "DIGIT",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 33, 210, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 27, 1, 28, 4, 28, 175, 8, 28, 11, 28, 12, 28, 176, 1, 28, 1, 28, 4,
		28, 181, 8, 28, 11, 28, 12, 28, 182, 3, 28, 185, 8, 28, 1, 29, 1, 29, 1,
		29, 5, 29, 190, 8, 29, 10, 29, 12, 29, 193, 9, 29, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 4, 33, 205, 8, 33, 11,
		33, 12, 33, 206, 1, 33, 1, 33, 0, 0, 34, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47,
		24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 0, 63, 31, 65,
		32, 67, 33, 1, 0, 5, 1, 0, 48, 57, 2, 0, 34, 34, 92, 92, 8, 0, 34, 34,
		47, 47, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 65,
		90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 214, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0,
		0, 0, 67, 1, 0, 0, 0, 1, 69, 1, 0, 0, 0, 3, 75, 1, 0, 0, 0, 5, 77, 1, 0,
		0, 0, 7, 79, 1, 0, 0, 0, 9, 81, 1, 0, 0, 0, 11, 83, 1, 0, 0, 0, 13, 85,
		1, 0, 0, 0, 15, 87, 1, 0, 0, 0, 17, 89, 1, 0, 0, 0, 19, 94, 1, 0, 0, 0,
		21, 100, 1, 0, 0, 0, 23, 104, 1, 0, 0, 0, 25, 106, 1, 0, 0, 0, 27, 115,
		1, 0, 0, 0, 29, 119, 1, 0, 0, 0, 31, 122, 1, 0, 0, 0, 33, 127, 1, 0, 0,
		0, 35, 134, 1, 0, 0, 0, 37, 139, 1, 0, 0, 0, 39, 143, 1, 0, 0, 0, 41, 146,
		1, 0, 0, 0, 43, 152, 1, 0, 0, 0, 45, 154, 1, 0, 0, 0, 47, 156, 1, 0, 0,
		0, 49, 158, 1, 0, 0, 0, 51, 160, 1, 0, 0, 0, 53, 165, 1, 0, 0, 0, 55, 170,
		1, 0, 0, 0, 57, 174, 1, 0, 0, 0, 59, 186, 1, 0, 0, 0, 61, 196, 1, 0, 0,
		0, 63, 199, 1, 0, 0, 0, 65, 201, 1, 0, 0, 0, 67, 204, 1, 0, 0, 0, 69, 70,
		5, 108, 0, 0, 70, 71, 5, 111, 0, 0, 71, 72, 5, 99, 0, 0, 72, 73, 5, 97,
		0, 0, 73, 74, 5, 108, 0, 0, 74, 2, 1, 0, 0, 0, 75, 76, 5, 61, 0, 0, 76,
		4, 1, 0, 0, 0, 77, 78, 5, 40, 0, 0, 78, 6, 1, 0, 0, 0, 79, 80, 5, 41, 0,
		0, 80, 8, 1, 0, 0, 0, 81, 82, 5, 42, 0, 0, 82, 10, 1, 0, 0, 0, 83, 84,
		5, 47, 0, 0, 84, 12, 1, 0, 0, 0, 85, 86, 5, 43, 0, 0, 86, 14, 1, 0, 0,
		0, 87, 88, 5, 45, 0, 0, 88, 16, 1, 0, 0, 0, 89, 90, 5, 116, 0, 0, 90, 91,
		5, 114, 0, 0, 91, 92, 5, 117, 0, 0, 92, 93, 5, 101, 0, 0, 93, 18, 1, 0,
		0, 0, 94, 95, 5, 102, 0, 0, 95, 96, 5, 97, 0, 0, 96, 97, 5, 108, 0, 0,
		97, 98, 5, 115, 0, 0, 98, 99, 5, 101, 0, 0, 99, 20, 1, 0, 0, 0, 100, 101,
		5, 110, 0, 0, 101, 102, 5, 105, 0, 0, 102, 103, 5, 108, 0, 0, 103, 22,
		1, 0, 0, 0, 104, 105, 5, 44, 0, 0, 105, 24, 1, 0, 0, 0, 106, 107, 5, 102,
		0, 0, 107, 108, 5, 117, 0, 0, 108, 109, 5, 110, 0, 0, 109, 110, 5, 99,
		0, 0, 110, 111, 5, 116, 0, 0, 111, 112, 5, 105, 0, 0, 112, 113, 5, 111,
		0, 0, 113, 114, 5, 110, 0, 0, 114, 26, 1, 0, 0, 0, 115, 116, 5, 101, 0,
		0, 116, 117, 5, 110, 0, 0, 117, 118, 5, 100, 0, 0, 118, 28, 1, 0, 0, 0,
		119, 120, 5, 105, 0, 0, 120, 121, 5, 102, 0, 0, 121, 30, 1, 0, 0, 0, 122,
		123, 5, 116, 0, 0, 123, 124, 5, 104, 0, 0, 124, 125, 5, 101, 0, 0, 125,
		126, 5, 110, 0, 0, 126, 32, 1, 0, 0, 0, 127, 128, 5, 101, 0, 0, 128, 129,
		5, 108, 0, 0, 129, 130, 5, 115, 0, 0, 130, 131, 5, 101, 0, 0, 131, 132,
		5, 105, 0, 0, 132, 133, 5, 102, 0, 0, 133, 34, 1, 0, 0, 0, 134, 135, 5,
		101, 0, 0, 135, 136, 5, 108, 0, 0, 136, 137, 5, 115, 0, 0, 137, 138, 5,
		101, 0, 0, 138, 36, 1, 0, 0, 0, 139, 140, 5, 102, 0, 0, 140, 141, 5, 111,
		0, 0, 141, 142, 5, 114, 0, 0, 142, 38, 1, 0, 0, 0, 143, 144, 5, 100, 0,
		0, 144, 145, 5, 111, 0, 0, 145, 40, 1, 0, 0, 0, 146, 147, 5, 119, 0, 0,
		147, 148, 5, 104, 0, 0, 148, 149, 5, 105, 0, 0, 149, 150, 5, 108, 0, 0,
		150, 151, 5, 101, 0, 0, 151, 42, 1, 0, 0, 0, 152, 153, 5, 123, 0, 0, 153,
		44, 1, 0, 0, 0, 154, 155, 5, 125, 0, 0, 155, 46, 1, 0, 0, 0, 156, 157,
		5, 91, 0, 0, 157, 48, 1, 0, 0, 0, 158, 159, 5, 93, 0, 0, 159, 50, 1, 0,
		0, 0, 160, 161, 5, 45, 0, 0, 161, 162, 5, 45, 0, 0, 162, 163, 5, 91, 0,
		0, 163, 164, 5, 91, 0, 0, 164, 52, 1, 0, 0, 0, 165, 166, 5, 45, 0, 0, 166,
		167, 5, 45, 0, 0, 167, 168, 5, 93, 0, 0, 168, 169, 5, 93, 0, 0, 169, 54,
		1, 0, 0, 0, 170, 171, 5, 93, 0, 0, 171, 172, 5, 93, 0, 0, 172, 56, 1, 0,
		0, 0, 173, 175, 7, 0, 0, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0,
		176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 184, 1, 0, 0, 0, 178,
		180, 5, 46, 0, 0, 179, 181, 7, 0, 0, 0, 180, 179, 1, 0, 0, 0, 181, 182,
		1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 185, 1, 0,
		0, 0, 184, 178, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 58, 1, 0, 0, 0,
		186, 191, 5, 34, 0, 0, 187, 190, 3, 61, 30, 0, 188, 190, 8, 1, 0, 0, 189,
		187, 1, 0, 0, 0, 189, 188, 1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189,
		1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 194, 1, 0, 0, 0, 193, 191, 1, 0,
		0, 0, 194, 195, 5, 34, 0, 0, 195, 60, 1, 0, 0, 0, 196, 197, 5, 92, 0, 0,
		197, 198, 7, 2, 0, 0, 198, 62, 1, 0, 0, 0, 199, 200, 7, 3, 0, 0, 200, 64,
		1, 0, 0, 0, 201, 202, 7, 0, 0, 0, 202, 66, 1, 0, 0, 0, 203, 205, 7, 4,
		0, 0, 204, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0,
		206, 207, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 6, 33, 0, 0, 209,
		68, 1, 0, 0, 0, 7, 0, 176, 182, 184, 189, 191, 206, 1, 6, 0, 0,
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

// lua_grammar_antlr4LexerInit initializes any static state used to implement lua_grammar_antlr4Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// Newlua_grammar_antlr4Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Lua_grammar_antlr4LexerInit() {
	staticData := &Lua_grammar_antlr4LexerLexerStaticData
	staticData.once.Do(lua_grammar_antlr4lexerLexerInit)
}

// Newlua_grammar_antlr4Lexer produces a new lexer instance for the optional input antlr.CharStream.
func Newlua_grammar_antlr4Lexer(input antlr.CharStream) *lua_grammar_antlr4Lexer {
	Lua_grammar_antlr4LexerInit()
	l := new(lua_grammar_antlr4Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &Lua_grammar_antlr4LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "lua_grammar_antlr4.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// lua_grammar_antlr4Lexer tokens.
const (
	lua_grammar_antlr4LexerT__0   = 1
	lua_grammar_antlr4LexerT__1   = 2
	lua_grammar_antlr4LexerT__2   = 3
	lua_grammar_antlr4LexerT__3   = 4
	lua_grammar_antlr4LexerT__4   = 5
	lua_grammar_antlr4LexerT__5   = 6
	lua_grammar_antlr4LexerT__6   = 7
	lua_grammar_antlr4LexerT__7   = 8
	lua_grammar_antlr4LexerT__8   = 9
	lua_grammar_antlr4LexerT__9   = 10
	lua_grammar_antlr4LexerT__10  = 11
	lua_grammar_antlr4LexerT__11  = 12
	lua_grammar_antlr4LexerT__12  = 13
	lua_grammar_antlr4LexerT__13  = 14
	lua_grammar_antlr4LexerT__14  = 15
	lua_grammar_antlr4LexerT__15  = 16
	lua_grammar_antlr4LexerT__16  = 17
	lua_grammar_antlr4LexerT__17  = 18
	lua_grammar_antlr4LexerT__18  = 19
	lua_grammar_antlr4LexerT__19  = 20
	lua_grammar_antlr4LexerT__20  = 21
	lua_grammar_antlr4LexerT__21  = 22
	lua_grammar_antlr4LexerT__22  = 23
	lua_grammar_antlr4LexerT__23  = 24
	lua_grammar_antlr4LexerT__24  = 25
	lua_grammar_antlr4LexerT__25  = 26
	lua_grammar_antlr4LexerT__26  = 27
	lua_grammar_antlr4LexerT__27  = 28
	lua_grammar_antlr4LexerNUMBER = 29
	lua_grammar_antlr4LexerSTRING = 30
	lua_grammar_antlr4LexerLETTER = 31
	lua_grammar_antlr4LexerDIGIT  = 32
	lua_grammar_antlr4LexerWS     = 33
)
