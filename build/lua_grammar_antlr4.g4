grammar lua_grammar_antlr4;

program: (statement | multi_line_comment)*;

statement
    : assignment
    | expression
    | if_statement
    | for_statement
    | while_statement
    | function_declaration
    ;

assignment
    : ('local')? identifier '=' expression
    ;

expression
    : literal
    | identifier
    | function_call
    | table
    | table_access
    | '(' expression ')'
    | expression arith_operator=('*'|'/') expression
    | expression arith_operator=('+'|'-') expression
    ;

literal
    : NUMBER
    | STRING
    | 'true'
    | 'false'
    | 'nil'
    ;

function_call
    : identifier '(' (expression (',' expression)*)? ')'
    ;

function_declaration
    : 'function' identifier '(' (identifier (',' identifier)*)? ')' block 'end'
    ;

block
    : (statement)*
    ;

if_statement
    : 'if' expression 'then' block ('elseif' expression 'then' block)* ('else' block)? 'end'
    ;

for_statement
    : 'for' identifier '=' expression ',' expression (',' expression)? 'do' block 'end'
    ;

while_statement
    : 'while' expression 'do' block 'end'
    ;

table
    : '{' (field (',' field)*)? '}'
    ;

field
    : identifier '=' expression
    | expression
    ;

table_access
    : identifier '[' expression ']'
    ;

multi_line_comment
    : '--[[' ~'--]]'* ']]'
    ;

identifier
    : LETTER (LETTER|DIGIT)*;

NUMBER
    : [0-9]+ ('.' [0-9]+)?;

STRING
    : '"' (ESC | ~["\\])* '"'
    ;

fragment ESC
    : '\\' ["\\/bfnrt]
    ;

LETTER
    : [a-zA-Z_];

DIGIT
    : [0-9];

WS
    : [ \t\r\n]+ -> skip;
