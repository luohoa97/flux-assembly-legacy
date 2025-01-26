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
		"program", "statement", "control_statement", "statement_terminator",
		"assignment", "expression", "operators", "comparison_operator", "arith_operator",
		"logical_operator", "comment", "literal", "function_call", "table_insert",
		"function_declaration", "block", "if_statement", "for_statement", "while_statement",
		"table", "field", "table_access", "multi_line_comment", "single_line_comment",
		"print_statement", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 269, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		1, 0, 5, 0, 55, 8, 0, 10, 0, 12, 0, 58, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 65, 8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 70, 8, 2, 1, 3, 1, 3, 1, 4,
		3, 4, 75, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 92, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		5, 5, 98, 8, 5, 10, 5, 12, 5, 101, 9, 5, 1, 6, 1, 6, 1, 6, 3, 6, 106, 8,
		6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 3, 10, 116, 8, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 125, 8, 12, 10,
		12, 12, 12, 128, 9, 12, 3, 12, 130, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 140, 8, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14,
		155, 8, 14, 10, 14, 12, 14, 158, 9, 14, 3, 14, 160, 8, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 15, 4, 15, 167, 8, 15, 11, 15, 12, 15, 168, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 180, 8, 16,
		10, 16, 12, 16, 183, 9, 16, 1, 16, 1, 16, 3, 16, 187, 8, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 207, 8, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 3, 17, 213, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 225, 8, 19, 10, 19, 12, 19, 228,
		9, 19, 3, 19, 230, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 3, 20, 239, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		5, 22, 248, 8, 22, 10, 22, 12, 22, 251, 9, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 5, 25, 264, 8, 25,
		10, 25, 12, 25, 267, 9, 25, 1, 25, 0, 1, 10, 26, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 0, 6, 1, 0, 4, 9, 1, 0, 10, 13, 2, 0, 29, 29, 41, 41, 4, 0, 35, 35,
		39, 39, 45, 45, 50, 51, 1, 0, 21, 21, 2, 0, 24, 24, 52, 53, 276, 0, 56,
		1, 0, 0, 0, 2, 64, 1, 0, 0, 0, 4, 69, 1, 0, 0, 0, 6, 71, 1, 0, 0, 0, 8,
		74, 1, 0, 0, 0, 10, 91, 1, 0, 0, 0, 12, 105, 1, 0, 0, 0, 14, 107, 1, 0,
		0, 0, 16, 109, 1, 0, 0, 0, 18, 111, 1, 0, 0, 0, 20, 115, 1, 0, 0, 0, 22,
		117, 1, 0, 0, 0, 24, 139, 1, 0, 0, 0, 26, 141, 1, 0, 0, 0, 28, 148, 1,
		0, 0, 0, 30, 166, 1, 0, 0, 0, 32, 170, 1, 0, 0, 0, 34, 212, 1, 0, 0, 0,
		36, 214, 1, 0, 0, 0, 38, 220, 1, 0, 0, 0, 40, 238, 1, 0, 0, 0, 42, 240,
		1, 0, 0, 0, 44, 245, 1, 0, 0, 0, 46, 254, 1, 0, 0, 0, 48, 256, 1, 0, 0,
		0, 50, 261, 1, 0, 0, 0, 52, 55, 3, 2, 1, 0, 53, 55, 3, 20, 10, 0, 54, 52,
		1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0,
		56, 57, 1, 0, 0, 0, 57, 1, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59, 65, 3, 8,
		4, 0, 60, 65, 3, 10, 5, 0, 61, 65, 3, 4, 2, 0, 62, 65, 3, 28, 14, 0, 63,
		65, 3, 24, 12, 0, 64, 59, 1, 0, 0, 0, 64, 60, 1, 0, 0, 0, 64, 61, 1, 0,
		0, 0, 64, 62, 1, 0, 0, 0, 64, 63, 1, 0, 0, 0, 65, 3, 1, 0, 0, 0, 66, 70,
		3, 32, 16, 0, 67, 70, 3, 34, 17, 0, 68, 70, 3, 36, 18, 0, 69, 66, 1, 0,
		0, 0, 69, 67, 1, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 5, 1, 0, 0, 0, 71, 72,
		5, 25, 0, 0, 72, 7, 1, 0, 0, 0, 73, 75, 5, 48, 0, 0, 74, 73, 1, 0, 0, 0,
		74, 75, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 3, 50, 25, 0, 77, 78, 5,
		1, 0, 0, 78, 79, 3, 10, 5, 0, 79, 80, 3, 6, 3, 0, 80, 9, 1, 0, 0, 0, 81,
		82, 6, 5, -1, 0, 82, 92, 3, 22, 11, 0, 83, 92, 3, 50, 25, 0, 84, 85, 5,
		2, 0, 0, 85, 86, 3, 10, 5, 0, 86, 87, 5, 3, 0, 0, 87, 92, 1, 0, 0, 0, 88,
		92, 3, 24, 12, 0, 89, 92, 3, 38, 19, 0, 90, 92, 3, 42, 21, 0, 91, 81, 1,
		0, 0, 0, 91, 83, 1, 0, 0, 0, 91, 84, 1, 0, 0, 0, 91, 88, 1, 0, 0, 0, 91,
		89, 1, 0, 0, 0, 91, 90, 1, 0, 0, 0, 92, 99, 1, 0, 0, 0, 93, 94, 10, 4,
		0, 0, 94, 95, 3, 12, 6, 0, 95, 96, 3, 10, 5, 5, 96, 98, 1, 0, 0, 0, 97,
		93, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0,
		0, 0, 100, 11, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 106, 3, 14, 7, 0,
		103, 106, 3, 16, 8, 0, 104, 106, 3, 18, 9, 0, 105, 102, 1, 0, 0, 0, 105,
		103, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106, 13, 1, 0, 0, 0, 107, 108, 7,
		0, 0, 0, 108, 15, 1, 0, 0, 0, 109, 110, 7, 1, 0, 0, 110, 17, 1, 0, 0, 0,
		111, 112, 7, 2, 0, 0, 112, 19, 1, 0, 0, 0, 113, 116, 3, 46, 23, 0, 114,
		116, 3, 44, 22, 0, 115, 113, 1, 0, 0, 0, 115, 114, 1, 0, 0, 0, 116, 21,
		1, 0, 0, 0, 117, 118, 7, 3, 0, 0, 118, 23, 1, 0, 0, 0, 119, 120, 3, 50,
		25, 0, 120, 129, 5, 2, 0, 0, 121, 126, 3, 10, 5, 0, 122, 123, 5, 14, 0,
		0, 123, 125, 3, 10, 5, 0, 124, 122, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126,
		124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126,
		1, 0, 0, 0, 129, 121, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 1, 0,
		0, 0, 131, 132, 5, 3, 0, 0, 132, 140, 1, 0, 0, 0, 133, 140, 3, 26, 13,
		0, 134, 135, 5, 28, 0, 0, 135, 136, 5, 2, 0, 0, 136, 137, 3, 10, 5, 0,
		137, 138, 5, 3, 0, 0, 138, 140, 1, 0, 0, 0, 139, 119, 1, 0, 0, 0, 139,
		133, 1, 0, 0, 0, 139, 134, 1, 0, 0, 0, 140, 25, 1, 0, 0, 0, 141, 142, 5,
		15, 0, 0, 142, 143, 5, 2, 0, 0, 143, 144, 3, 50, 25, 0, 144, 145, 5, 14,
		0, 0, 145, 146, 3, 10, 5, 0, 146, 147, 5, 3, 0, 0, 147, 27, 1, 0, 0, 0,
		148, 149, 5, 49, 0, 0, 149, 150, 3, 50, 25, 0, 150, 159, 5, 2, 0, 0, 151,
		156, 3, 50, 25, 0, 152, 153, 5, 14, 0, 0, 153, 155, 3, 50, 25, 0, 154,
		152, 1, 0, 0, 0, 155, 158, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157,
		1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159, 151, 1, 0,
		0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 5, 3, 0, 0,
		162, 163, 3, 30, 15, 0, 163, 164, 5, 34, 0, 0, 164, 29, 1, 0, 0, 0, 165,
		167, 3, 2, 1, 0, 166, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 166,
		1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 31, 1, 0, 0, 0, 170, 171, 5, 38,
		0, 0, 171, 172, 3, 10, 5, 0, 172, 173, 5, 44, 0, 0, 173, 181, 3, 30, 15,
		0, 174, 175, 5, 33, 0, 0, 175, 176, 3, 10, 5, 0, 176, 177, 5, 44, 0, 0,
		177, 178, 3, 30, 15, 0, 178, 180, 1, 0, 0, 0, 179, 174, 1, 0, 0, 0, 180,
		183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 186,
		1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 185, 5, 32, 0, 0, 185, 187, 3, 30,
		15, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0,
		188, 189, 5, 34, 0, 0, 189, 33, 1, 0, 0, 0, 190, 191, 5, 36, 0, 0, 191,
		192, 3, 50, 25, 0, 192, 193, 5, 27, 0, 0, 193, 194, 3, 10, 5, 0, 194, 195,
		5, 31, 0, 0, 195, 196, 3, 30, 15, 0, 196, 197, 5, 34, 0, 0, 197, 213, 1,
		0, 0, 0, 198, 199, 5, 36, 0, 0, 199, 200, 3, 50, 25, 0, 200, 201, 5, 1,
		0, 0, 201, 202, 3, 10, 5, 0, 202, 203, 5, 14, 0, 0, 203, 206, 3, 10, 5,
		0, 204, 205, 5, 14, 0, 0, 205, 207, 3, 10, 5, 0, 206, 204, 1, 0, 0, 0,
		206, 207, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 5, 31, 0, 0, 209,
		210, 3, 30, 15, 0, 210, 211, 5, 34, 0, 0, 211, 213, 1, 0, 0, 0, 212, 190,
		1, 0, 0, 0, 212, 198, 1, 0, 0, 0, 213, 35, 1, 0, 0, 0, 214, 215, 5, 47,
		0, 0, 215, 216, 3, 10, 5, 0, 216, 217, 5, 31, 0, 0, 217, 218, 3, 30, 15,
		0, 218, 219, 5, 34, 0, 0, 219, 37, 1, 0, 0, 0, 220, 229, 5, 16, 0, 0, 221,
		226, 3, 40, 20, 0, 222, 223, 5, 14, 0, 0, 223, 225, 3, 40, 20, 0, 224,
		222, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227,
		1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229, 221, 1, 0,
		0, 0, 229, 230, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 232, 5, 17, 0, 0,
		232, 39, 1, 0, 0, 0, 233, 234, 3, 50, 25, 0, 234, 235, 5, 1, 0, 0, 235,
		236, 3, 10, 5, 0, 236, 239, 1, 0, 0, 0, 237, 239, 3, 10, 5, 0, 238, 233,
		1, 0, 0, 0, 238, 237, 1, 0, 0, 0, 239, 41, 1, 0, 0, 0, 240, 241, 3, 50,
		25, 0, 241, 242, 5, 18, 0, 0, 242, 243, 3, 10, 5, 0, 243, 244, 5, 19, 0,
		0, 244, 43, 1, 0, 0, 0, 245, 249, 5, 20, 0, 0, 246, 248, 8, 4, 0, 0, 247,
		246, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249, 250,
		1, 0, 0, 0, 250, 252, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 252, 253, 5, 22,
		0, 0, 253, 45, 1, 0, 0, 0, 254, 255, 5, 23, 0, 0, 255, 47, 1, 0, 0, 0,
		256, 257, 5, 28, 0, 0, 257, 258, 5, 2, 0, 0, 258, 259, 3, 10, 5, 0, 259,
		260, 5, 3, 0, 0, 260, 49, 1, 0, 0, 0, 261, 265, 5, 52, 0, 0, 262, 264,
		7, 5, 0, 0, 263, 262, 1, 0, 0, 0, 264, 267, 1, 0, 0, 0, 265, 263, 1, 0,
		0, 0, 265, 266, 1, 0, 0, 0, 266, 51, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0,
		24, 54, 56, 64, 69, 74, 91, 99, 105, 115, 126, 129, 139, 156, 159, 168,
		181, 186, 206, 212, 226, 229, 238, 249, 265,
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
	lua_grammar_antlr4ParserEOF         = antlr.TokenEOF
	lua_grammar_antlr4ParserT__0        = 1
	lua_grammar_antlr4ParserT__1        = 2
	lua_grammar_antlr4ParserT__2        = 3
	lua_grammar_antlr4ParserT__3        = 4
	lua_grammar_antlr4ParserT__4        = 5
	lua_grammar_antlr4ParserT__5        = 6
	lua_grammar_antlr4ParserT__6        = 7
	lua_grammar_antlr4ParserT__7        = 8
	lua_grammar_antlr4ParserT__8        = 9
	lua_grammar_antlr4ParserT__9        = 10
	lua_grammar_antlr4ParserT__10       = 11
	lua_grammar_antlr4ParserT__11       = 12
	lua_grammar_antlr4ParserT__12       = 13
	lua_grammar_antlr4ParserT__13       = 14
	lua_grammar_antlr4ParserT__14       = 15
	lua_grammar_antlr4ParserT__15       = 16
	lua_grammar_antlr4ParserT__16       = 17
	lua_grammar_antlr4ParserT__17       = 18
	lua_grammar_antlr4ParserT__18       = 19
	lua_grammar_antlr4ParserT__19       = 20
	lua_grammar_antlr4ParserT__20       = 21
	lua_grammar_antlr4ParserT__21       = 22
	lua_grammar_antlr4ParserT__22       = 23
	lua_grammar_antlr4ParserT__23       = 24
	lua_grammar_antlr4ParserSEPARATOR   = 25
	lua_grammar_antlr4ParserNEWLINE     = 26
	lua_grammar_antlr4ParserKW_IN       = 27
	lua_grammar_antlr4ParserKW_PRINT    = 28
	lua_grammar_antlr4ParserKW_AND      = 29
	lua_grammar_antlr4ParserKW_BREAK    = 30
	lua_grammar_antlr4ParserKW_DO       = 31
	lua_grammar_antlr4ParserKW_ELSE     = 32
	lua_grammar_antlr4ParserKW_ELSEIF   = 33
	lua_grammar_antlr4ParserKW_END      = 34
	lua_grammar_antlr4ParserKW_FALSE    = 35
	lua_grammar_antlr4ParserKW_FOR      = 36
	lua_grammar_antlr4ParserKW_GOTO     = 37
	lua_grammar_antlr4ParserKW_IF       = 38
	lua_grammar_antlr4ParserKW_NIL      = 39
	lua_grammar_antlr4ParserKW_NOT      = 40
	lua_grammar_antlr4ParserKW_OR       = 41
	lua_grammar_antlr4ParserKW_REPEAT   = 42
	lua_grammar_antlr4ParserKW_RETURN   = 43
	lua_grammar_antlr4ParserKW_THEN     = 44
	lua_grammar_antlr4ParserKW_TRUE     = 45
	lua_grammar_antlr4ParserKW_UNTIL    = 46
	lua_grammar_antlr4ParserKW_WHILE    = 47
	lua_grammar_antlr4ParserKW_LOCAL    = 48
	lua_grammar_antlr4ParserKW_FUNCTION = 49
	lua_grammar_antlr4ParserNUMBER      = 50
	lua_grammar_antlr4ParserSTRING      = 51
	lua_grammar_antlr4ParserLETTER      = 52
	lua_grammar_antlr4ParserDIGIT       = 53
	lua_grammar_antlr4ParserWS          = 54
)

// lua_grammar_antlr4Parser rules.
const (
	lua_grammar_antlr4ParserRULE_program              = 0
	lua_grammar_antlr4ParserRULE_statement            = 1
	lua_grammar_antlr4ParserRULE_control_statement    = 2
	lua_grammar_antlr4ParserRULE_statement_terminator = 3
	lua_grammar_antlr4ParserRULE_assignment           = 4
	lua_grammar_antlr4ParserRULE_expression           = 5
	lua_grammar_antlr4ParserRULE_operators            = 6
	lua_grammar_antlr4ParserRULE_comparison_operator  = 7
	lua_grammar_antlr4ParserRULE_arith_operator       = 8
	lua_grammar_antlr4ParserRULE_logical_operator     = 9
	lua_grammar_antlr4ParserRULE_comment              = 10
	lua_grammar_antlr4ParserRULE_literal              = 11
	lua_grammar_antlr4ParserRULE_function_call        = 12
	lua_grammar_antlr4ParserRULE_table_insert         = 13
	lua_grammar_antlr4ParserRULE_function_declaration = 14
	lua_grammar_antlr4ParserRULE_block                = 15
	lua_grammar_antlr4ParserRULE_if_statement         = 16
	lua_grammar_antlr4ParserRULE_for_statement        = 17
	lua_grammar_antlr4ParserRULE_while_statement      = 18
	lua_grammar_antlr4ParserRULE_table                = 19
	lua_grammar_antlr4ParserRULE_field                = 20
	lua_grammar_antlr4ParserRULE_table_access         = 21
	lua_grammar_antlr4ParserRULE_multi_line_comment   = 22
	lua_grammar_antlr4ParserRULE_single_line_comment  = 23
	lua_grammar_antlr4ParserRULE_print_statement      = 24
	lua_grammar_antlr4ParserRULE_identifier           = 25
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllComment() []ICommentContext
	Comment(i int) ICommentContext

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

func (s *ProgramContext) AllComment() []ICommentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommentContext); ok {
			len++
		}
	}

	tst := make([]ICommentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommentContext); ok {
			tst[i] = t.(ICommentContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Comment(i int) ICommentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
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

	return t.(ICommentContext)
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
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8902574129381380) != 0 {
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case lua_grammar_antlr4ParserT__1, lua_grammar_antlr4ParserT__14, lua_grammar_antlr4ParserT__15, lua_grammar_antlr4ParserKW_PRINT, lua_grammar_antlr4ParserKW_FALSE, lua_grammar_antlr4ParserKW_FOR, lua_grammar_antlr4ParserKW_IF, lua_grammar_antlr4ParserKW_NIL, lua_grammar_antlr4ParserKW_TRUE, lua_grammar_antlr4ParserKW_WHILE, lua_grammar_antlr4ParserKW_LOCAL, lua_grammar_antlr4ParserKW_FUNCTION, lua_grammar_antlr4ParserNUMBER, lua_grammar_antlr4ParserSTRING, lua_grammar_antlr4ParserLETTER:
			{
				p.SetState(52)
				p.Statement()
			}

		case lua_grammar_antlr4ParserT__19, lua_grammar_antlr4ParserT__22:
			{
				p.SetState(53)
				p.Comment()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(58)
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
	Control_statement() IControl_statementContext
	Function_declaration() IFunction_declarationContext
	Function_call() IFunction_callContext

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

func (s *StatementContext) Control_statement() IControl_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControl_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControl_statementContext)
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

func (s *StatementContext) Function_call() IFunction_callContext {
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
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.Control_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(62)
			p.Function_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.Function_call()
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

// IControl_statementContext is an interface to support dynamic dispatch.
type IControl_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_statement() IIf_statementContext
	For_statement() IFor_statementContext
	While_statement() IWhile_statementContext

	// IsControl_statementContext differentiates from other interfaces.
	IsControl_statementContext()
}

type Control_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControl_statementContext() *Control_statementContext {
	var p = new(Control_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_control_statement
	return p
}

func InitEmptyControl_statementContext(p *Control_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_control_statement
}

func (*Control_statementContext) IsControl_statementContext() {}

func NewControl_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_statementContext {
	var p = new(Control_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_control_statement

	return p
}

func (s *Control_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_statementContext) If_statement() IIf_statementContext {
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

func (s *Control_statementContext) For_statement() IFor_statementContext {
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

func (s *Control_statementContext) While_statement() IWhile_statementContext {
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

func (s *Control_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterControl_statement(s)
	}
}

func (s *Control_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitControl_statement(s)
	}
}

func (s *Control_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitControl_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Control_statement() (localctx IControl_statementContext) {
	localctx = NewControl_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, lua_grammar_antlr4ParserRULE_control_statement)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserKW_IF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.If_statement()
		}

	case lua_grammar_antlr4ParserKW_FOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.For_statement()
		}

	case lua_grammar_antlr4ParserKW_WHILE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.While_statement()
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

// IStatement_terminatorContext is an interface to support dynamic dispatch.
type IStatement_terminatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEPARATOR() antlr.TerminalNode

	// IsStatement_terminatorContext differentiates from other interfaces.
	IsStatement_terminatorContext()
}

type Statement_terminatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatement_terminatorContext() *Statement_terminatorContext {
	var p = new(Statement_terminatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_statement_terminator
	return p
}

func InitEmptyStatement_terminatorContext(p *Statement_terminatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_statement_terminator
}

func (*Statement_terminatorContext) IsStatement_terminatorContext() {}

func NewStatement_terminatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statement_terminatorContext {
	var p = new(Statement_terminatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_statement_terminator

	return p
}

func (s *Statement_terminatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Statement_terminatorContext) SEPARATOR() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserSEPARATOR, 0)
}

func (s *Statement_terminatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_terminatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statement_terminatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterStatement_terminator(s)
	}
}

func (s *Statement_terminatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitStatement_terminator(s)
	}
}

func (s *Statement_terminatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitStatement_terminator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Statement_terminator() (localctx IStatement_terminatorContext) {
	localctx = NewStatement_terminatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, lua_grammar_antlr4ParserRULE_statement_terminator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(lua_grammar_antlr4ParserSEPARATOR)
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Expression() IExpressionContext
	Statement_terminator() IStatement_terminatorContext
	KW_LOCAL() antlr.TerminalNode

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

func (s *AssignmentContext) Statement_terminator() IStatement_terminatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_terminatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_terminatorContext)
}

func (s *AssignmentContext) KW_LOCAL() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_LOCAL, 0)
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
	p.EnterRule(localctx, 8, lua_grammar_antlr4ParserRULE_assignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserKW_LOCAL {
		{
			p.SetState(73)
			p.Match(lua_grammar_antlr4ParserKW_LOCAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(76)
		p.Identifier()
	}
	{
		p.SetState(77)
		p.Match(lua_grammar_antlr4ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.expression(0)
	}
	{
		p.SetState(79)
		p.Statement_terminator()
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

	// Getter signatures
	Literal() ILiteralContext
	Identifier() IIdentifierContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	Function_call() IFunction_callContext
	Table() ITableContext
	Table_access() ITable_accessContext
	Operators() IOperatorsContext

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

func (s *ExpressionContext) Operators() IOperatorsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorsContext)
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, lua_grammar_antlr4ParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(82)
			p.Literal()
		}

	case 2:
		{
			p.SetState(83)
			p.Identifier()
		}

	case 3:
		{
			p.SetState(84)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.expression(0)
		}
		{
			p.SetState(86)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(88)
			p.Function_call()
		}

	case 5:
		{
			p.SetState(89)
			p.Table()
		}

	case 6:
		{
			p.SetState(90)
			p.Table_access()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(99)
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
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_expression)
			p.SetState(93)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				goto errorExit
			}
			{
				p.SetState(94)
				p.Operators()
			}
			{
				p.SetState(95)
				p.expression(5)
			}

		}
		p.SetState(101)
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

// IOperatorsContext is an interface to support dynamic dispatch.
type IOperatorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Comparison_operator() IComparison_operatorContext
	Arith_operator() IArith_operatorContext
	Logical_operator() ILogical_operatorContext

	// IsOperatorsContext differentiates from other interfaces.
	IsOperatorsContext()
}

type OperatorsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorsContext() *OperatorsContext {
	var p = new(OperatorsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_operators
	return p
}

func InitEmptyOperatorsContext(p *OperatorsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_operators
}

func (*OperatorsContext) IsOperatorsContext() {}

func NewOperatorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorsContext {
	var p = new(OperatorsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_operators

	return p
}

func (s *OperatorsContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorsContext) Comparison_operator() IComparison_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparison_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparison_operatorContext)
}

func (s *OperatorsContext) Arith_operator() IArith_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArith_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArith_operatorContext)
}

func (s *OperatorsContext) Logical_operator() ILogical_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_operatorContext)
}

func (s *OperatorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterOperators(s)
	}
}

func (s *OperatorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitOperators(s)
	}
}

func (s *OperatorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitOperators(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Operators() (localctx IOperatorsContext) {
	localctx = NewOperatorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, lua_grammar_antlr4ParserRULE_operators)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__3, lua_grammar_antlr4ParserT__4, lua_grammar_antlr4ParserT__5, lua_grammar_antlr4ParserT__6, lua_grammar_antlr4ParserT__7, lua_grammar_antlr4ParserT__8:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Comparison_operator()
		}

	case lua_grammar_antlr4ParserT__9, lua_grammar_antlr4ParserT__10, lua_grammar_antlr4ParserT__11, lua_grammar_antlr4ParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Arith_operator()
		}

	case lua_grammar_antlr4ParserKW_AND, lua_grammar_antlr4ParserKW_OR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.Logical_operator()
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

// IComparison_operatorContext is an interface to support dynamic dispatch.
type IComparison_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsComparison_operatorContext differentiates from other interfaces.
	IsComparison_operatorContext()
}

type Comparison_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparison_operatorContext() *Comparison_operatorContext {
	var p = new(Comparison_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_comparison_operator
	return p
}

func InitEmptyComparison_operatorContext(p *Comparison_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_comparison_operator
}

func (*Comparison_operatorContext) IsComparison_operatorContext() {}

func NewComparison_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparison_operatorContext {
	var p = new(Comparison_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_comparison_operator

	return p
}

func (s *Comparison_operatorContext) GetParser() antlr.Parser { return s.parser }
func (s *Comparison_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparison_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comparison_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterComparison_operator(s)
	}
}

func (s *Comparison_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitComparison_operator(s)
	}
}

func (s *Comparison_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitComparison_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Comparison_operator() (localctx IComparison_operatorContext) {
	localctx = NewComparison_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, lua_grammar_antlr4ParserRULE_comparison_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1008) != 0) {
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

// IArith_operatorContext is an interface to support dynamic dispatch.
type IArith_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArith_operatorContext differentiates from other interfaces.
	IsArith_operatorContext()
}

type Arith_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArith_operatorContext() *Arith_operatorContext {
	var p = new(Arith_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_arith_operator
	return p
}

func InitEmptyArith_operatorContext(p *Arith_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_arith_operator
}

func (*Arith_operatorContext) IsArith_operatorContext() {}

func NewArith_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arith_operatorContext {
	var p = new(Arith_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_arith_operator

	return p
}

func (s *Arith_operatorContext) GetParser() antlr.Parser { return s.parser }
func (s *Arith_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arith_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arith_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterArith_operator(s)
	}
}

func (s *Arith_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitArith_operator(s)
	}
}

func (s *Arith_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitArith_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Arith_operator() (localctx IArith_operatorContext) {
	localctx = NewArith_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, lua_grammar_antlr4ParserRULE_arith_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15360) != 0) {
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

// ILogical_operatorContext is an interface to support dynamic dispatch.
type ILogical_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_AND() antlr.TerminalNode
	KW_OR() antlr.TerminalNode

	// IsLogical_operatorContext differentiates from other interfaces.
	IsLogical_operatorContext()
}

type Logical_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogical_operatorContext() *Logical_operatorContext {
	var p = new(Logical_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_logical_operator
	return p
}

func InitEmptyLogical_operatorContext(p *Logical_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_logical_operator
}

func (*Logical_operatorContext) IsLogical_operatorContext() {}

func NewLogical_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_operatorContext {
	var p = new(Logical_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_logical_operator

	return p
}

func (s *Logical_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_operatorContext) KW_AND() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_AND, 0)
}

func (s *Logical_operatorContext) KW_OR() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_OR, 0)
}

func (s *Logical_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterLogical_operator(s)
	}
}

func (s *Logical_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitLogical_operator(s)
	}
}

func (s *Logical_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitLogical_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Logical_operator() (localctx ILogical_operatorContext) {
	localctx = NewLogical_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, lua_grammar_antlr4ParserRULE_logical_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		_la = p.GetTokenStream().LA(1)

		if !(_la == lua_grammar_antlr4ParserKW_AND || _la == lua_grammar_antlr4ParserKW_OR) {
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

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Single_line_comment() ISingle_line_commentContext
	Multi_line_comment() IMulti_line_commentContext

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) Single_line_comment() ISingle_line_commentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingle_line_commentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingle_line_commentContext)
}

func (s *CommentContext) Multi_line_comment() IMulti_line_commentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulti_line_commentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulti_line_commentContext)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitComment(s)
	}
}

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, lua_grammar_antlr4ParserRULE_comment)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.Single_line_comment()
		}

	case lua_grammar_antlr4ParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Multi_line_comment()
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	KW_TRUE() antlr.TerminalNode
	KW_FALSE() antlr.TerminalNode
	KW_NIL() antlr.TerminalNode

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

func (s *LiteralContext) KW_TRUE() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_TRUE, 0)
}

func (s *LiteralContext) KW_FALSE() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_FALSE, 0)
}

func (s *LiteralContext) KW_NIL() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_NIL, 0)
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
	p.EnterRule(localctx, 22, lua_grammar_antlr4ParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3413468208168960) != 0) {
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
	Table_insert() ITable_insertContext
	KW_PRINT() antlr.TerminalNode

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

func (s *Function_callContext) Table_insert() ITable_insertContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_insertContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_insertContext)
}

func (s *Function_callContext) KW_PRINT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_PRINT, 0)
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
	p.EnterRule(localctx, 24, lua_grammar_antlr4ParserRULE_function_call)
	var _la int

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserLETTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Identifier()
		}
		{
			p.SetState(120)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7917068104073220) != 0 {
			{
				p.SetState(121)
				p.expression(0)
			}
			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == lua_grammar_antlr4ParserT__13 {
				{
					p.SetState(122)
					p.Match(lua_grammar_antlr4ParserT__13)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(123)
					p.expression(0)
				}

				p.SetState(128)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(131)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case lua_grammar_antlr4ParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Table_insert()
		}

	case lua_grammar_antlr4ParserKW_PRINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.Match(lua_grammar_antlr4ParserKW_PRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.expression(0)
		}
		{
			p.SetState(137)
			p.Match(lua_grammar_antlr4ParserT__2)
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

// ITable_insertContext is an interface to support dynamic dispatch.
type ITable_insertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Expression() IExpressionContext

	// IsTable_insertContext differentiates from other interfaces.
	IsTable_insertContext()
}

type Table_insertContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_insertContext() *Table_insertContext {
	var p = new(Table_insertContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table_insert
	return p
}

func InitEmptyTable_insertContext(p *Table_insertContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table_insert
}

func (*Table_insertContext) IsTable_insertContext() {}

func NewTable_insertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_insertContext {
	var p = new(Table_insertContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table_insert

	return p
}

func (s *Table_insertContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_insertContext) Identifier() IIdentifierContext {
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

func (s *Table_insertContext) Expression() IExpressionContext {
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

func (s *Table_insertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_insertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_insertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterTable_insert(s)
	}
}

func (s *Table_insertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitTable_insert(s)
	}
}

func (s *Table_insertContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitTable_insert(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Table_insert() (localctx ITable_insertContext) {
	localctx = NewTable_insertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, lua_grammar_antlr4ParserRULE_table_insert)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(lua_grammar_antlr4ParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Identifier()
	}
	{
		p.SetState(144)
		p.Match(lua_grammar_antlr4ParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.expression(0)
	}
	{
		p.SetState(146)
		p.Match(lua_grammar_antlr4ParserT__2)
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
	KW_FUNCTION() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	Block() IBlockContext
	KW_END() antlr.TerminalNode

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

func (s *Function_declarationContext) KW_FUNCTION() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_FUNCTION, 0)
}

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

func (s *Function_declarationContext) KW_END() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_END, 0)
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
	p.EnterRule(localctx, 28, lua_grammar_antlr4ParserRULE_function_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(lua_grammar_antlr4ParserKW_FUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Identifier()
	}
	{
		p.SetState(150)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserLETTER {
		{
			p.SetState(151)
			p.Identifier()
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__13 {
			{
				p.SetState(152)
				p.Match(lua_grammar_antlr4ParserT__13)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(153)
				p.Identifier()
			}

			p.SetState(158)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(161)
		p.Match(lua_grammar_antlr4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Block()
	}
	{
		p.SetState(163)
		p.Match(lua_grammar_antlr4ParserKW_END)
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
	p.EnterRule(localctx, 30, lua_grammar_antlr4ParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8902574119944196) != 0) {
		{
			p.SetState(165)
			p.Statement()
		}

		p.SetState(168)
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
	KW_IF() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllKW_THEN() []antlr.TerminalNode
	KW_THEN(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	KW_END() antlr.TerminalNode
	AllKW_ELSEIF() []antlr.TerminalNode
	KW_ELSEIF(i int) antlr.TerminalNode
	KW_ELSE() antlr.TerminalNode

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

func (s *If_statementContext) KW_IF() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_IF, 0)
}

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

func (s *If_statementContext) AllKW_THEN() []antlr.TerminalNode {
	return s.GetTokens(lua_grammar_antlr4ParserKW_THEN)
}

func (s *If_statementContext) KW_THEN(i int) antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_THEN, i)
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

func (s *If_statementContext) KW_END() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_END, 0)
}

func (s *If_statementContext) AllKW_ELSEIF() []antlr.TerminalNode {
	return s.GetTokens(lua_grammar_antlr4ParserKW_ELSEIF)
}

func (s *If_statementContext) KW_ELSEIF(i int) antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_ELSEIF, i)
}

func (s *If_statementContext) KW_ELSE() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_ELSE, 0)
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
	p.EnterRule(localctx, 32, lua_grammar_antlr4ParserRULE_if_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(lua_grammar_antlr4ParserKW_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.expression(0)
	}
	{
		p.SetState(172)
		p.Match(lua_grammar_antlr4ParserKW_THEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Block()
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == lua_grammar_antlr4ParserKW_ELSEIF {
		{
			p.SetState(174)
			p.Match(lua_grammar_antlr4ParserKW_ELSEIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.expression(0)
		}
		{
			p.SetState(176)
			p.Match(lua_grammar_antlr4ParserKW_THEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Block()
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserKW_ELSE {
		{
			p.SetState(184)
			p.Match(lua_grammar_antlr4ParserKW_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Block()
		}

	}
	{
		p.SetState(188)
		p.Match(lua_grammar_antlr4ParserKW_END)
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
	KW_FOR() antlr.TerminalNode
	Identifier() IIdentifierContext
	KW_IN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	KW_DO() antlr.TerminalNode
	Block() IBlockContext
	KW_END() antlr.TerminalNode

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

func (s *For_statementContext) KW_FOR() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_FOR, 0)
}

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

func (s *For_statementContext) KW_IN() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_IN, 0)
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

func (s *For_statementContext) KW_DO() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_DO, 0)
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

func (s *For_statementContext) KW_END() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_END, 0)
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
	p.EnterRule(localctx, 34, lua_grammar_antlr4ParserRULE_for_statement)
	var _la int

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.Match(lua_grammar_antlr4ParserKW_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Identifier()
		}
		{
			p.SetState(192)
			p.Match(lua_grammar_antlr4ParserKW_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.expression(0)
		}
		{
			p.SetState(194)
			p.Match(lua_grammar_antlr4ParserKW_DO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Block()
		}
		{
			p.SetState(196)
			p.Match(lua_grammar_antlr4ParserKW_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.Match(lua_grammar_antlr4ParserKW_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Identifier()
		}
		{
			p.SetState(200)
			p.Match(lua_grammar_antlr4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.expression(0)
		}
		{
			p.SetState(202)
			p.Match(lua_grammar_antlr4ParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.expression(0)
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == lua_grammar_antlr4ParserT__13 {
			{
				p.SetState(204)
				p.Match(lua_grammar_antlr4ParserT__13)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(205)
				p.expression(0)
			}

		}
		{
			p.SetState(208)
			p.Match(lua_grammar_antlr4ParserKW_DO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Block()
		}
		{
			p.SetState(210)
			p.Match(lua_grammar_antlr4ParserKW_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IWhile_statementContext is an interface to support dynamic dispatch.
type IWhile_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_WHILE() antlr.TerminalNode
	Expression() IExpressionContext
	KW_DO() antlr.TerminalNode
	Block() IBlockContext
	KW_END() antlr.TerminalNode

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

func (s *While_statementContext) KW_WHILE() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_WHILE, 0)
}

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

func (s *While_statementContext) KW_DO() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_DO, 0)
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

func (s *While_statementContext) KW_END() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_END, 0)
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
	p.EnterRule(localctx, 36, lua_grammar_antlr4ParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(lua_grammar_antlr4ParserKW_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.expression(0)
	}
	{
		p.SetState(216)
		p.Match(lua_grammar_antlr4ParserKW_DO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Block()
	}
	{
		p.SetState(218)
		p.Match(lua_grammar_antlr4ParserKW_END)
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
	p.EnterRule(localctx, 38, lua_grammar_antlr4ParserRULE_table)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(lua_grammar_antlr4ParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7917068104073220) != 0 {
		{
			p.SetState(221)
			p.Field()
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__13 {
			{
				p.SetState(222)
				p.Match(lua_grammar_antlr4ParserT__13)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(223)
				p.Field()
			}

			p.SetState(228)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(231)
		p.Match(lua_grammar_antlr4ParserT__16)
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
	p.EnterRule(localctx, 40, lua_grammar_antlr4ParserRULE_field)
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(233)
			p.Identifier()
		}
		{
			p.SetState(234)
			p.Match(lua_grammar_antlr4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(237)
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
	p.EnterRule(localctx, 42, lua_grammar_antlr4ParserRULE_table_access)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Identifier()
	}
	{
		p.SetState(241)
		p.Match(lua_grammar_antlr4ParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.expression(0)
	}
	{
		p.SetState(243)
		p.Match(lua_grammar_antlr4ParserT__18)
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
	p.EnterRule(localctx, 44, lua_grammar_antlr4ParserRULE_multi_line_comment)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(lua_grammar_antlr4ParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(246)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || _la == lua_grammar_antlr4ParserT__20 {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Match(lua_grammar_antlr4ParserT__21)
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

// ISingle_line_commentContext is an interface to support dynamic dispatch.
type ISingle_line_commentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingle_line_commentContext differentiates from other interfaces.
	IsSingle_line_commentContext()
}

type Single_line_commentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_line_commentContext() *Single_line_commentContext {
	var p = new(Single_line_commentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_single_line_comment
	return p
}

func InitEmptySingle_line_commentContext(p *Single_line_commentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_single_line_comment
}

func (*Single_line_commentContext) IsSingle_line_commentContext() {}

func NewSingle_line_commentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_line_commentContext {
	var p = new(Single_line_commentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_single_line_comment

	return p
}

func (s *Single_line_commentContext) GetParser() antlr.Parser { return s.parser }
func (s *Single_line_commentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_line_commentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_line_commentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterSingle_line_comment(s)
	}
}

func (s *Single_line_commentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitSingle_line_comment(s)
	}
}

func (s *Single_line_commentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitSingle_line_comment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Single_line_comment() (localctx ISingle_line_commentContext) {
	localctx = NewSingle_line_commentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, lua_grammar_antlr4ParserRULE_single_line_comment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
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

// IPrint_statementContext is an interface to support dynamic dispatch.
type IPrint_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_PRINT() antlr.TerminalNode
	Expression() IExpressionContext

	// IsPrint_statementContext differentiates from other interfaces.
	IsPrint_statementContext()
}

type Print_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrint_statementContext() *Print_statementContext {
	var p = new(Print_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_print_statement
	return p
}

func InitEmptyPrint_statementContext(p *Print_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_print_statement
}

func (*Print_statementContext) IsPrint_statementContext() {}

func NewPrint_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print_statementContext {
	var p = new(Print_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_print_statement

	return p
}

func (s *Print_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Print_statementContext) KW_PRINT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_PRINT, 0)
}

func (s *Print_statementContext) Expression() IExpressionContext {
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

func (s *Print_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Print_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterPrint_statement(s)
	}
}

func (s *Print_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitPrint_statement(s)
	}
}

func (s *Print_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitPrint_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Print_statement() (localctx IPrint_statementContext) {
	localctx = NewPrint_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, lua_grammar_antlr4ParserRULE_print_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(lua_grammar_antlr4ParserKW_PRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.expression(0)
	}
	{
		p.SetState(259)
		p.Match(lua_grammar_antlr4ParserT__2)
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
	p.EnterRule(localctx, 50, lua_grammar_antlr4ParserRULE_identifier)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(lua_grammar_antlr4ParserLETTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(262)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13510798898888704) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
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
	case 5:
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
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
