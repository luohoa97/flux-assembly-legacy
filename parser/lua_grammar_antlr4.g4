grammar lua_grammar_antlr4;

program: (statement | comment)*;

statement
    : assignment
    | expression
    | control_statement
    | function_declaration
    | function_call
    ;

control_statement
    : if_statement
    | for_statement
    | while_statement
    ;

statement_terminator: SEPARATOR;

assignment: (KW_LOCAL)? identifier '=' expression statement_terminator;

expression
    : literal
    | identifier
    | '(' expression ')'
    | expression operators expression
    | function_call
    | table
    | table_access
    ;

operators
    : comparison_operator
    | arith_operator
    | logical_operator
    ;

comparison_operator: '>' | '<' | '>=' | '==' | '<=' | '~=';
arith_operator: '*' | '/' | '+' | '-';
logical_operator: KW_AND | KW_OR;

comment: single_line_comment | multi_line_comment;

literal
    : NUMBER
    | STRING
    | KW_TRUE
    | KW_FALSE
    | KW_NIL
    ;

function_call: identifier '(' (expression (',' expression)*)? ')'
    | table_insert
    | KW_PRINT '(' expression ')'
    ; 

table_insert: 'table.insert' '(' identifier ',' expression ')';

function_declaration: KW_FUNCTION identifier '(' (identifier (',' identifier)*)? ')' block KW_END;

block: statement+;

if_statement: KW_IF expression KW_THEN block (KW_ELSEIF expression KW_THEN block)* (KW_ELSE block)? KW_END;
for_statement
    : KW_FOR identifier KW_IN expression KW_DO block KW_END
    | KW_FOR identifier '=' expression ',' expression (',' expression)? KW_DO block KW_END
    ;
while_statement: KW_WHILE expression KW_DO block KW_END;

table: '{' (field (',' field)*)? '}';

field: identifier '=' expression | expression;

table_access: identifier '[' expression ']';

multi_line_comment: '--[[' ~'--]]'* ']]';

single_line_comment: '--';

SEPARATOR: ';' | NEWLINE;

NEWLINE: [\r\n]+ -> skip;

print_statement: KW_PRINT '(' expression ')';

KW_IN: 'in';
KW_PRINT: 'print';
KW_AND: 'and';
KW_BREAK: 'break';
KW_DO: 'do';
KW_ELSE: 'else';
KW_ELSEIF: 'elseif';
KW_END: 'end';
KW_FALSE: 'false';
KW_FOR: 'for';
KW_GOTO: 'goto';
KW_IF: 'if';
KW_NIL: 'nil';
KW_NOT: 'not';
KW_OR: 'or';
KW_REPEAT: 'repeat';
KW_RETURN: 'return';
KW_THEN: 'then';
KW_TRUE: 'true';
KW_UNTIL: 'until';
KW_WHILE: 'while';
KW_LOCAL: 'local';
KW_FUNCTION: 'function';
identifier: LETTER (LETTER | DIGIT | '_')*;

NUMBER: [0-9]+ ('.' [0-9]+)?;

STRING: '"' (ESC | ~["\\])* '"';

fragment ESC: '\\' ["\\/bfnrt];

LETTER: [a-zA-Z_];

DIGIT: [0-9];

WS: [ \t\r\n]+ -> skip;
