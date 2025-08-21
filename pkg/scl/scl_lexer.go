package scl

import (
	"fmt"
	"strings"
	"unicode"
)

type lexer struct {
	input string
	pos   int
}

const (
	Iota      int  = 0
	Dot       byte = '.'
	Colon     byte = ':'
	Semicolon byte = ';'
)

func (l *lexer) Lex(lval *yySymType) int {
	for l.pos < len(l.input) {
		ch := rune(l.input[l.pos])
		l.pos++

		// skip spaces
		if unicode.IsSpace(ch) {
			continue
		}

		// identifiers & keywords
		if unicode.IsLetter(ch) {
			return l.handleLetter(lval)
		}

		// numbers
		if unicode.IsDigit(ch) {
			start := l.pos - 1
			for l.pos < len(l.input) && unicode.IsDigit(rune(l.input[l.pos])) {
				l.pos++
			}
			lval.str = l.input[start:l.pos]
			return NUMBER
		}

		// strings "..."
		if ch == '"' {
			start := l.pos
			for l.pos < len(l.input) && rune(l.input[l.pos]) != '"' {
				l.pos++
			}
			str := l.input[start:l.pos]
			l.pos++ // skip closing "
			lval.str = str
			return STRING
		}

		// symbols
		switch ch {
		case ':':
			return COLON
		case ';':
			return SEMICOLON
		case '.':
			return DOT
		}

		panic(fmt.Sprintf("unexpected character: %q", ch))
	}
	return 0
}

func (l *lexer) Error(s string) {
	fmt.Println("Parse error:", s)
}

func (l *lexer) handleLetter(lval *yySymType) int {
	start := l.pos - 1
	for l.pos < len(l.input) &&
		(unicode.IsLetter(rune(l.input[l.pos])) || unicode.IsDigit(rune(l.input[l.pos]))) {
		l.pos++
	}
	word := l.input[start:l.pos]
	upper := strings.ToUpper(word)
	switch upper {
	case "SET":
		return SET
	case "GET":
		return GET
	case "DELETE":
		return DELETE
	case "TRUNCATE":
		return TRUNCATE
	case "DROP":
		return DROP
	case "UPDATE":
		return UPDATE
	default:
		lval.str = word

		// TODO:
		// This works pretty okeish, but I think it can improve
		// It doesn't have to be super dynamic at this time, but I rather have a very strong
		// lexer that makes sense instead of only forcefully parse specific words that might be there

		if l.input[l.pos] == Dot {
			return COLLECTION
		}

		if l.input[l.pos] == Colon {
			return KEY
		}

		if l.input[l.pos] == Semicolon {
			if strings.Contains(l.input, "DROP") || strings.Contains(l.input, "TRUNCATE") {
				return COLLECTION
			} else {
				if strings.Contains(l.input, ".") {
					return KEY
				}
			}

			return KEY
		}

		return yyErrCode
	}
}
