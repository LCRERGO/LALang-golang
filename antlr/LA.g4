/*
 * This file is the lexer specification of language L.A.
 * Most of the tokens have been writen down in portuguese
 * for the discipline of compiler construction
*/
lexer grammar LA;

PALAVRA_CHAVE: 'algoritmo' | 'declare' | 'literal' | 'inteiro' | 'fim_algoritmo' |
    'leia' | 'escreva' | 'real' | 'e' | 'ou' | 'nao' | 'logico' | 'se' | 'senao' |
    'fim_se' | 'entao' | 'caso' | 'seja' | '..' | 'fim_caso' | 'para' | 'faca' |
    'ate' | 'fim_para' | 'enquanto' | 'fim_enquanto' | 'registro' | 'fim_registro' |
    'tipo' | 'procedimento' | 'var' | 'fim_procedimento' | 'retorne' | 'fim_funcao' |
    'funcao' | 'constante' | 'falso' | 'verdadeiro';  

NUM_INT: [0-9]+;
NUM_REAL: [0-9]+ '.' [0-9]+;
IDENT: ([a-z] | [A-Z])([a-z] | [A-Z] | [0-9] | '_')*;
CADEIA: '"' (~('\n' | '\r' | '"'))* '"';

/* Skippable Tokens */
COMENTARIO: '{' (~('}' | '\n' | '\r'))* '}' -> skip;
WS: (' ' | '\t' | '\r' | '\n') -> skip;

/* Operators */
OP_ARIT: '+' | '-' | '*' | '/' | '%';
OP_REL: '>' | '>='| '<' | '<=' | '<>' | '=';
OP_PON: '^' | '&' | '.' | '[' | ']';
DELIM: ':';
ATRIB: '<-';
VIRGULA: ',';

/* Precedence Rules */
ABREPARENTESES: '(';
FECHAPARENTESES: ')';
