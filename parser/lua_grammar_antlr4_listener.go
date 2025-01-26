// Code generated from lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // lua_grammar_antlr4

import "github.com/antlr4-go/antlr/v4"

// lua_grammar_antlr4Listener is a complete listener for a parse tree produced by lua_grammar_antlr4Parser.
type lua_grammar_antlr4Listener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterFor_statement is called when entering the for_statement production.
	EnterFor_statement(c *For_statementContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterTable_access is called when entering the table_access production.
	EnterTable_access(c *Table_accessContext)

	// EnterMulti_line_comment is called when entering the multi_line_comment production.
	EnterMulti_line_comment(c *Multi_line_commentContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitFor_statement is called when exiting the for_statement production.
	ExitFor_statement(c *For_statementContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitTable_access is called when exiting the table_access production.
	ExitTable_access(c *Table_accessContext)

	// ExitMulti_line_comment is called when exiting the multi_line_comment production.
	ExitMulti_line_comment(c *Multi_line_commentContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
