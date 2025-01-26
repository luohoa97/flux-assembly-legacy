// Code generated from lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // lua_grammar_antlr4

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by lua_grammar_antlr4Parser.
type lua_grammar_antlr4Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by lua_grammar_antlr4Parser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#function_call.
	VisitFunction_call(ctx *Function_callContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#function_declaration.
	VisitFunction_declaration(ctx *Function_declarationContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#for_statement.
	VisitFor_statement(ctx *For_statementContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#while_statement.
	VisitWhile_statement(ctx *While_statementContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#table.
	VisitTable(ctx *TableContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#table_access.
	VisitTable_access(ctx *Table_accessContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#multi_line_comment.
	VisitMulti_line_comment(ctx *Multi_line_commentContext) interface{}

	// Visit a parse tree produced by lua_grammar_antlr4Parser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
