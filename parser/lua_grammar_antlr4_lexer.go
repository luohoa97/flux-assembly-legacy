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
		"", "'='", "'('", "')'", "'>'", "'<'", "'>='", "'=='", "'<='", "'~='",
		"'*'", "'/'", "'+'", "'-'", "','", "'table.insert'", "'{'", "'}'", "'['",
		"']'", "'--[['", "'--]]'", "']]'", "'--'", "'_'", "", "", "'in'", "'print'",
		"'and'", "'break'", "'do'", "'else'", "'elseif'", "'end'", "'false'",
		"'for'", "'goto'", "'if'", "'nil'", "'not'", "'or'", "'repeat'", "'return'",
		"'then'", "'true'", "'until'", "'while'", "'local'", "'function'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "SEPARATOR", "NEWLINE", "KW_IN", "KW_PRINT",
		"KW_AND", "KW_BREAK", "KW_DO", "KW_ELSE", "KW_ELSEIF", "KW_END", "KW_FALSE",
		"KW_FOR", "KW_GOTO", "KW_IF", "KW_NIL", "KW_NOT", "KW_OR", "KW_REPEAT",
		"KW_RETURN", "KW_THEN", "KW_TRUE", "KW_UNTIL", "KW_WHILE", "KW_LOCAL",
		"KW_FUNCTION", "NUMBER", "STRING", "LETTER", "DIGIT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "SEPARATOR",
		"NEWLINE", "KW_IN", "KW_PRINT", "KW_AND", "KW_BREAK", "KW_DO", "KW_ELSE",
		"KW_ELSEIF", "KW_END", "KW_FALSE", "KW_FOR", "KW_GOTO", "KW_IF", "KW_NIL",
		"KW_NOT", "KW_OR", "KW_REPEAT", "KW_RETURN", "KW_THEN", "KW_TRUE", "KW_UNTIL",
		"KW_WHILE", "KW_LOCAL", "KW_FUNCTION", "NUMBER", "STRING", "ESC", "LETTER",
		"DIGIT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 54, 348, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 24, 1, 24, 3, 24, 185, 8, 24, 1, 25, 4, 25, 188, 8, 25, 11, 25, 12,
		25, 189, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1,
		46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 4, 49, 313, 8, 49, 11,
		49, 12, 49, 314, 1, 49, 1, 49, 4, 49, 319, 8, 49, 11, 49, 12, 49, 320,
		3, 49, 323, 8, 49, 1, 50, 1, 50, 1, 50, 5, 50, 328, 8, 50, 10, 50, 12,
		50, 331, 9, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53,
		1, 53, 1, 54, 4, 54, 343, 8, 54, 11, 54, 12, 54, 344, 1, 54, 1, 54, 0,
		0, 55, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37,
		75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46,
		93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 0, 105, 52, 107, 53, 109,
		54, 1, 0, 6, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 2, 0, 34, 34, 92, 92,
		8, 0, 34, 34, 47, 47, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116,
		116, 3, 0, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 354, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47,
		1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0,
		55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0,
		0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0,
		0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0,
		0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1,
		0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93,
		1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0,
		101, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0,
		0, 0, 1, 111, 1, 0, 0, 0, 3, 113, 1, 0, 0, 0, 5, 115, 1, 0, 0, 0, 7, 117,
		1, 0, 0, 0, 9, 119, 1, 0, 0, 0, 11, 121, 1, 0, 0, 0, 13, 124, 1, 0, 0,
		0, 15, 127, 1, 0, 0, 0, 17, 130, 1, 0, 0, 0, 19, 133, 1, 0, 0, 0, 21, 135,
		1, 0, 0, 0, 23, 137, 1, 0, 0, 0, 25, 139, 1, 0, 0, 0, 27, 141, 1, 0, 0,
		0, 29, 143, 1, 0, 0, 0, 31, 156, 1, 0, 0, 0, 33, 158, 1, 0, 0, 0, 35, 160,
		1, 0, 0, 0, 37, 162, 1, 0, 0, 0, 39, 164, 1, 0, 0, 0, 41, 169, 1, 0, 0,
		0, 43, 174, 1, 0, 0, 0, 45, 177, 1, 0, 0, 0, 47, 180, 1, 0, 0, 0, 49, 184,
		1, 0, 0, 0, 51, 187, 1, 0, 0, 0, 53, 193, 1, 0, 0, 0, 55, 196, 1, 0, 0,
		0, 57, 202, 1, 0, 0, 0, 59, 206, 1, 0, 0, 0, 61, 212, 1, 0, 0, 0, 63, 215,
		1, 0, 0, 0, 65, 220, 1, 0, 0, 0, 67, 227, 1, 0, 0, 0, 69, 231, 1, 0, 0,
		0, 71, 237, 1, 0, 0, 0, 73, 241, 1, 0, 0, 0, 75, 246, 1, 0, 0, 0, 77, 249,
		1, 0, 0, 0, 79, 253, 1, 0, 0, 0, 81, 257, 1, 0, 0, 0, 83, 260, 1, 0, 0,
		0, 85, 267, 1, 0, 0, 0, 87, 274, 1, 0, 0, 0, 89, 279, 1, 0, 0, 0, 91, 284,
		1, 0, 0, 0, 93, 290, 1, 0, 0, 0, 95, 296, 1, 0, 0, 0, 97, 302, 1, 0, 0,
		0, 99, 312, 1, 0, 0, 0, 101, 324, 1, 0, 0, 0, 103, 334, 1, 0, 0, 0, 105,
		337, 1, 0, 0, 0, 107, 339, 1, 0, 0, 0, 109, 342, 1, 0, 0, 0, 111, 112,
		5, 61, 0, 0, 112, 2, 1, 0, 0, 0, 113, 114, 5, 40, 0, 0, 114, 4, 1, 0, 0,
		0, 115, 116, 5, 41, 0, 0, 116, 6, 1, 0, 0, 0, 117, 118, 5, 62, 0, 0, 118,
		8, 1, 0, 0, 0, 119, 120, 5, 60, 0, 0, 120, 10, 1, 0, 0, 0, 121, 122, 5,
		62, 0, 0, 122, 123, 5, 61, 0, 0, 123, 12, 1, 0, 0, 0, 124, 125, 5, 61,
		0, 0, 125, 126, 5, 61, 0, 0, 126, 14, 1, 0, 0, 0, 127, 128, 5, 60, 0, 0,
		128, 129, 5, 61, 0, 0, 129, 16, 1, 0, 0, 0, 130, 131, 5, 126, 0, 0, 131,
		132, 5, 61, 0, 0, 132, 18, 1, 0, 0, 0, 133, 134, 5, 42, 0, 0, 134, 20,
		1, 0, 0, 0, 135, 136, 5, 47, 0, 0, 136, 22, 1, 0, 0, 0, 137, 138, 5, 43,
		0, 0, 138, 24, 1, 0, 0, 0, 139, 140, 5, 45, 0, 0, 140, 26, 1, 0, 0, 0,
		141, 142, 5, 44, 0, 0, 142, 28, 1, 0, 0, 0, 143, 144, 5, 116, 0, 0, 144,
		145, 5, 97, 0, 0, 145, 146, 5, 98, 0, 0, 146, 147, 5, 108, 0, 0, 147, 148,
		5, 101, 0, 0, 148, 149, 5, 46, 0, 0, 149, 150, 5, 105, 0, 0, 150, 151,
		5, 110, 0, 0, 151, 152, 5, 115, 0, 0, 152, 153, 5, 101, 0, 0, 153, 154,
		5, 114, 0, 0, 154, 155, 5, 116, 0, 0, 155, 30, 1, 0, 0, 0, 156, 157, 5,
		123, 0, 0, 157, 32, 1, 0, 0, 0, 158, 159, 5, 125, 0, 0, 159, 34, 1, 0,
		0, 0, 160, 161, 5, 91, 0, 0, 161, 36, 1, 0, 0, 0, 162, 163, 5, 93, 0, 0,
		163, 38, 1, 0, 0, 0, 164, 165, 5, 45, 0, 0, 165, 166, 5, 45, 0, 0, 166,
		167, 5, 91, 0, 0, 167, 168, 5, 91, 0, 0, 168, 40, 1, 0, 0, 0, 169, 170,
		5, 45, 0, 0, 170, 171, 5, 45, 0, 0, 171, 172, 5, 93, 0, 0, 172, 173, 5,
		93, 0, 0, 173, 42, 1, 0, 0, 0, 174, 175, 5, 93, 0, 0, 175, 176, 5, 93,
		0, 0, 176, 44, 1, 0, 0, 0, 177, 178, 5, 45, 0, 0, 178, 179, 5, 45, 0, 0,
		179, 46, 1, 0, 0, 0, 180, 181, 5, 95, 0, 0, 181, 48, 1, 0, 0, 0, 182, 185,
		5, 59, 0, 0, 183, 185, 3, 51, 25, 0, 184, 182, 1, 0, 0, 0, 184, 183, 1,
		0, 0, 0, 185, 50, 1, 0, 0, 0, 186, 188, 7, 0, 0, 0, 187, 186, 1, 0, 0,
		0, 188, 189, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 192, 6, 25, 0, 0, 192, 52, 1, 0, 0, 0, 193, 194,
		5, 105, 0, 0, 194, 195, 5, 110, 0, 0, 195, 54, 1, 0, 0, 0, 196, 197, 5,
		112, 0, 0, 197, 198, 5, 114, 0, 0, 198, 199, 5, 105, 0, 0, 199, 200, 5,
		110, 0, 0, 200, 201, 5, 116, 0, 0, 201, 56, 1, 0, 0, 0, 202, 203, 5, 97,
		0, 0, 203, 204, 5, 110, 0, 0, 204, 205, 5, 100, 0, 0, 205, 58, 1, 0, 0,
		0, 206, 207, 5, 98, 0, 0, 207, 208, 5, 114, 0, 0, 208, 209, 5, 101, 0,
		0, 209, 210, 5, 97, 0, 0, 210, 211, 5, 107, 0, 0, 211, 60, 1, 0, 0, 0,
		212, 213, 5, 100, 0, 0, 213, 214, 5, 111, 0, 0, 214, 62, 1, 0, 0, 0, 215,
		216, 5, 101, 0, 0, 216, 217, 5, 108, 0, 0, 217, 218, 5, 115, 0, 0, 218,
		219, 5, 101, 0, 0, 219, 64, 1, 0, 0, 0, 220, 221, 5, 101, 0, 0, 221, 222,
		5, 108, 0, 0, 222, 223, 5, 115, 0, 0, 223, 224, 5, 101, 0, 0, 224, 225,
		5, 105, 0, 0, 225, 226, 5, 102, 0, 0, 226, 66, 1, 0, 0, 0, 227, 228, 5,
		101, 0, 0, 228, 229, 5, 110, 0, 0, 229, 230, 5, 100, 0, 0, 230, 68, 1,
		0, 0, 0, 231, 232, 5, 102, 0, 0, 232, 233, 5, 97, 0, 0, 233, 234, 5, 108,
		0, 0, 234, 235, 5, 115, 0, 0, 235, 236, 5, 101, 0, 0, 236, 70, 1, 0, 0,
		0, 237, 238, 5, 102, 0, 0, 238, 239, 5, 111, 0, 0, 239, 240, 5, 114, 0,
		0, 240, 72, 1, 0, 0, 0, 241, 242, 5, 103, 0, 0, 242, 243, 5, 111, 0, 0,
		243, 244, 5, 116, 0, 0, 244, 245, 5, 111, 0, 0, 245, 74, 1, 0, 0, 0, 246,
		247, 5, 105, 0, 0, 247, 248, 5, 102, 0, 0, 248, 76, 1, 0, 0, 0, 249, 250,
		5, 110, 0, 0, 250, 251, 5, 105, 0, 0, 251, 252, 5, 108, 0, 0, 252, 78,
		1, 0, 0, 0, 253, 254, 5, 110, 0, 0, 254, 255, 5, 111, 0, 0, 255, 256, 5,
		116, 0, 0, 256, 80, 1, 0, 0, 0, 257, 258, 5, 111, 0, 0, 258, 259, 5, 114,
		0, 0, 259, 82, 1, 0, 0, 0, 260, 261, 5, 114, 0, 0, 261, 262, 5, 101, 0,
		0, 262, 263, 5, 112, 0, 0, 263, 264, 5, 101, 0, 0, 264, 265, 5, 97, 0,
		0, 265, 266, 5, 116, 0, 0, 266, 84, 1, 0, 0, 0, 267, 268, 5, 114, 0, 0,
		268, 269, 5, 101, 0, 0, 269, 270, 5, 116, 0, 0, 270, 271, 5, 117, 0, 0,
		271, 272, 5, 114, 0, 0, 272, 273, 5, 110, 0, 0, 273, 86, 1, 0, 0, 0, 274,
		275, 5, 116, 0, 0, 275, 276, 5, 104, 0, 0, 276, 277, 5, 101, 0, 0, 277,
		278, 5, 110, 0, 0, 278, 88, 1, 0, 0, 0, 279, 280, 5, 116, 0, 0, 280, 281,
		5, 114, 0, 0, 281, 282, 5, 117, 0, 0, 282, 283, 5, 101, 0, 0, 283, 90,
		1, 0, 0, 0, 284, 285, 5, 117, 0, 0, 285, 286, 5, 110, 0, 0, 286, 287, 5,
		116, 0, 0, 287, 288, 5, 105, 0, 0, 288, 289, 5, 108, 0, 0, 289, 92, 1,
		0, 0, 0, 290, 291, 5, 119, 0, 0, 291, 292, 5, 104, 0, 0, 292, 293, 5, 105,
		0, 0, 293, 294, 5, 108, 0, 0, 294, 295, 5, 101, 0, 0, 295, 94, 1, 0, 0,
		0, 296, 297, 5, 108, 0, 0, 297, 298, 5, 111, 0, 0, 298, 299, 5, 99, 0,
		0, 299, 300, 5, 97, 0, 0, 300, 301, 5, 108, 0, 0, 301, 96, 1, 0, 0, 0,
		302, 303, 5, 102, 0, 0, 303, 304, 5, 117, 0, 0, 304, 305, 5, 110, 0, 0,
		305, 306, 5, 99, 0, 0, 306, 307, 5, 116, 0, 0, 307, 308, 5, 105, 0, 0,
		308, 309, 5, 111, 0, 0, 309, 310, 5, 110, 0, 0, 310, 98, 1, 0, 0, 0, 311,
		313, 7, 1, 0, 0, 312, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 312,
		1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 322, 1, 0, 0, 0, 316, 318, 5, 46,
		0, 0, 317, 319, 7, 1, 0, 0, 318, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0,
		320, 318, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 323, 1, 0, 0, 0, 322,
		316, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 100, 1, 0, 0, 0, 324, 329,
		5, 34, 0, 0, 325, 328, 3, 103, 51, 0, 326, 328, 8, 2, 0, 0, 327, 325, 1,
		0, 0, 0, 327, 326, 1, 0, 0, 0, 328, 331, 1, 0, 0, 0, 329, 327, 1, 0, 0,
		0, 329, 330, 1, 0, 0, 0, 330, 332, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 332,
		333, 5, 34, 0, 0, 333, 102, 1, 0, 0, 0, 334, 335, 5, 92, 0, 0, 335, 336,
		7, 3, 0, 0, 336, 104, 1, 0, 0, 0, 337, 338, 7, 4, 0, 0, 338, 106, 1, 0,
		0, 0, 339, 340, 7, 1, 0, 0, 340, 108, 1, 0, 0, 0, 341, 343, 7, 5, 0, 0,
		342, 341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 344,
		345, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 347, 6, 54, 0, 0, 347, 110,
		1, 0, 0, 0, 9, 0, 184, 189, 314, 320, 322, 327, 329, 344, 1, 6, 0, 0,
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
	lua_grammar_antlr4LexerT__0        = 1
	lua_grammar_antlr4LexerT__1        = 2
	lua_grammar_antlr4LexerT__2        = 3
	lua_grammar_antlr4LexerT__3        = 4
	lua_grammar_antlr4LexerT__4        = 5
	lua_grammar_antlr4LexerT__5        = 6
	lua_grammar_antlr4LexerT__6        = 7
	lua_grammar_antlr4LexerT__7        = 8
	lua_grammar_antlr4LexerT__8        = 9
	lua_grammar_antlr4LexerT__9        = 10
	lua_grammar_antlr4LexerT__10       = 11
	lua_grammar_antlr4LexerT__11       = 12
	lua_grammar_antlr4LexerT__12       = 13
	lua_grammar_antlr4LexerT__13       = 14
	lua_grammar_antlr4LexerT__14       = 15
	lua_grammar_antlr4LexerT__15       = 16
	lua_grammar_antlr4LexerT__16       = 17
	lua_grammar_antlr4LexerT__17       = 18
	lua_grammar_antlr4LexerT__18       = 19
	lua_grammar_antlr4LexerT__19       = 20
	lua_grammar_antlr4LexerT__20       = 21
	lua_grammar_antlr4LexerT__21       = 22
	lua_grammar_antlr4LexerT__22       = 23
	lua_grammar_antlr4LexerT__23       = 24
	lua_grammar_antlr4LexerSEPARATOR   = 25
	lua_grammar_antlr4LexerNEWLINE     = 26
	lua_grammar_antlr4LexerKW_IN       = 27
	lua_grammar_antlr4LexerKW_PRINT    = 28
	lua_grammar_antlr4LexerKW_AND      = 29
	lua_grammar_antlr4LexerKW_BREAK    = 30
	lua_grammar_antlr4LexerKW_DO       = 31
	lua_grammar_antlr4LexerKW_ELSE     = 32
	lua_grammar_antlr4LexerKW_ELSEIF   = 33
	lua_grammar_antlr4LexerKW_END      = 34
	lua_grammar_antlr4LexerKW_FALSE    = 35
	lua_grammar_antlr4LexerKW_FOR      = 36
	lua_grammar_antlr4LexerKW_GOTO     = 37
	lua_grammar_antlr4LexerKW_IF       = 38
	lua_grammar_antlr4LexerKW_NIL      = 39
	lua_grammar_antlr4LexerKW_NOT      = 40
	lua_grammar_antlr4LexerKW_OR       = 41
	lua_grammar_antlr4LexerKW_REPEAT   = 42
	lua_grammar_antlr4LexerKW_RETURN   = 43
	lua_grammar_antlr4LexerKW_THEN     = 44
	lua_grammar_antlr4LexerKW_TRUE     = 45
	lua_grammar_antlr4LexerKW_UNTIL    = 46
	lua_grammar_antlr4LexerKW_WHILE    = 47
	lua_grammar_antlr4LexerKW_LOCAL    = 48
	lua_grammar_antlr4LexerKW_FUNCTION = 49
	lua_grammar_antlr4LexerNUMBER      = 50
	lua_grammar_antlr4LexerSTRING      = 51
	lua_grammar_antlr4LexerLETTER      = 52
	lua_grammar_antlr4LexerDIGIT       = 53
	lua_grammar_antlr4LexerWS          = 54
)
