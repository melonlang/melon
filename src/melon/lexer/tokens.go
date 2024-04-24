package lexer

import (
	"fmt"
)

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER
	STRING
	IDENTIFIER

	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_PAREN
	CLOSE_PAREN
	OPEN_BRACE
	CLOSE_BRACE

	ASSIGN
	EQUALS
	NOT
	NOT_EQUAL

	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL

	OR
	AND

	DOT
	DOT_DOT
	COLON
	QUESTION
	COMMA
	APPEND
	PREPEND
	PLUS_EQUALS
	MINUS_EQUALS
	MULTIPLY_EQUALS
	DIVIDE_EQUALS
	MODULUS_EQUALS
	POWER_EQUALS

	PLUS
	MINUS
	STAR
	SLASH
	PERCENT
	POWER
	DASH

	// reserved keywords
	LET
	VAR
	NEW
	IMPORT
	FROM
	FUNC
	IF
	ELSE
	ELSEIF
	WHILE
	DO
	FOR
	IN
)

type Token struct {
	Value string
	Kind  TokenKind
}

var keywords = map[string]TokenKind{
	"let":    LET,
	"var":    VAR,
	"new":    NEW,
	"import": IMPORT,
	"from":   FROM,
	"func":   FUNC,
	"if":     IF,
	"else":   ELSE,
	"elseif": ELSEIF,
	"while":  WHILE,
	"do":     DO,
	"for":    FOR,
	"in":     IN,
}

func (token Token) isOneOfMany(expectedTokes ...TokenKind) bool {
	for _, expected := range expectedTokes {
		if expected == token.Kind {
			return true
		}
	}
	return false
}

func (token Token) Debug() {
	if token.isOneOfMany(IDENTIFIER, NUMBER, STRING) {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(token.Kind))
	}
}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		value,
		kind,
	}
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "EOF"
	case NUMBER:
		return "NUMBER"
	case STRING:
		return "STRING"
	case IDENTIFIER:
		return "IDENTIFIER"
	case OPEN_BRACKET:
		return "OPEN_BRACKET"
	case CLOSE_BRACKET:
		return "CLOSE_BRACKET"
	case OPEN_PAREN:
		return "OPEN_PAREN"
	case CLOSE_PAREN:
		return "CLOSE_PAREN"
	case OPEN_BRACE:
		return "OPEN_BRACE"
	case CLOSE_BRACE:
		return "CLOSE_BRACE"
	case ASSIGN:
		return "ASSIGN"
	case EQUALS:
		return "EQUALS"
	case NOT:
		return "NOT"
	case NOT_EQUAL:
		return "NOT_EQUAL"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case OR:
		return "OR"
	case AND:
		return "AND"
	case DOT:
		return "DOT"
	case DOT_DOT:
		return "DOT_DOT"
	case COLON:
		return "COLON"
	case QUESTION:
		return "QUESTION"
	case COMMA:
		return "COMMA"
	case APPEND:
		return "APPEND"
	case PREPEND:
		return "PREPEND"
	case PLUS_EQUALS:
		return "PLUS_EQUALS"
	case MINUS_EQUALS:
		return "MINUS_EQUALS"
	case MULTIPLY_EQUALS:
		return "MULTIPLY_EQUALS"
	case DIVIDE_EQUALS:
		return "DIVIDE_EQUALS"
	case MODULUS_EQUALS:
		return "MODULUS_EQUALS"
	case POWER_EQUALS:
		return "POWER_EQUALS"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case STAR:
		return "STAR"
	case SLASH:
		return "SLASH"
	case PERCENT:
		return "PERCENT"
	case POWER:
		return "POWER"
	case DASH:
		return "DASH"
	case LET:
		return "LET"
	case VAR:
		return "VAR"
	case NEW:
		return "NEW"
	case IMPORT:
		return "IMPORT"
	case FROM:
		return "FROM"
	case FUNC:
		return "FUNC"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case ELSEIF:
		return "ELSEIF"
	case WHILE:
		return "WHILE"
	case DO:
		return "DO"
	case FOR:
		return "FOR"
	case IN:
		return "IN"
	default:
		return fmt.Sprintf("unknown token kind: %d", kind)
	}
}
