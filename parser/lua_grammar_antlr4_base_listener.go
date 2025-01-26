// Code generated from lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // lua_grammar_antlr4

import "github.com/antlr4-go/antlr/v4"

// Baselua_grammar_antlr4Listener is a complete listener for a parse tree produced by lua_grammar_antlr4Parser.
type Baselua_grammar_antlr4Listener struct{}

var _ lua_grammar_antlr4Listener = &Baselua_grammar_antlr4Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Baselua_grammar_antlr4Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Baselua_grammar_antlr4Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Baselua_grammar_antlr4Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Baselua_grammar_antlr4Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *Baselua_grammar_antlr4Listener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *Baselua_grammar_antlr4Listener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitStatement(ctx *StatementContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *Baselua_grammar_antlr4Listener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *Baselua_grammar_antlr4Listener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *Baselua_grammar_antlr4Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *Baselua_grammar_antlr4Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterLiteral is called when production literal is entered.
func (s *Baselua_grammar_antlr4Listener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *Baselua_grammar_antlr4Listener) ExitLiteral(ctx *LiteralContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *Baselua_grammar_antlr4Listener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *Baselua_grammar_antlr4Listener) ExitFunction_call(ctx *Function_callContext) {}

// EnterFunction_declaration is called when production function_declaration is entered.
func (s *Baselua_grammar_antlr4Listener) EnterFunction_declaration(ctx *Function_declarationContext) {
}

// ExitFunction_declaration is called when production function_declaration is exited.
func (s *Baselua_grammar_antlr4Listener) ExitFunction_declaration(ctx *Function_declarationContext) {}

// EnterBlock is called when production block is entered.
func (s *Baselua_grammar_antlr4Listener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *Baselua_grammar_antlr4Listener) ExitBlock(ctx *BlockContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitIf_statement(ctx *If_statementContext) {}

// EnterFor_statement is called when production for_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterFor_statement(ctx *For_statementContext) {}

// ExitFor_statement is called when production for_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitFor_statement(ctx *For_statementContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterTable is called when production table is entered.
func (s *Baselua_grammar_antlr4Listener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *Baselua_grammar_antlr4Listener) ExitTable(ctx *TableContext) {}

// EnterField is called when production field is entered.
func (s *Baselua_grammar_antlr4Listener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *Baselua_grammar_antlr4Listener) ExitField(ctx *FieldContext) {}

// EnterTable_access is called when production table_access is entered.
func (s *Baselua_grammar_antlr4Listener) EnterTable_access(ctx *Table_accessContext) {}

// ExitTable_access is called when production table_access is exited.
func (s *Baselua_grammar_antlr4Listener) ExitTable_access(ctx *Table_accessContext) {}

// EnterMulti_line_comment is called when production multi_line_comment is entered.
func (s *Baselua_grammar_antlr4Listener) EnterMulti_line_comment(ctx *Multi_line_commentContext) {}

// ExitMulti_line_comment is called when production multi_line_comment is exited.
func (s *Baselua_grammar_antlr4Listener) ExitMulti_line_comment(ctx *Multi_line_commentContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *Baselua_grammar_antlr4Listener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *Baselua_grammar_antlr4Listener) ExitIdentifier(ctx *IdentifierContext) {}
