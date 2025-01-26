// Code generated from lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // lua_grammar_antlr4

import "github.com/antlr4-go/antlr/v4"

type Baselua_grammar_antlr4Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *Baselua_grammar_antlr4Visitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitFunction_call(ctx *Function_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitFunction_declaration(ctx *Function_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitFor_statement(ctx *For_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitWhile_statement(ctx *While_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitTable(ctx *TableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitTable_access(ctx *Table_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitMulti_line_comment(ctx *Multi_line_commentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
