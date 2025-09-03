package scl

import (
	"fmt"
	"strings"
	"unicode"
)

type lexer struct {
	input     string
	pos       int
	lastToken int
	expect    expectState
}

// Expectation state helps us classify identifiers as COLLECTION or KEY
type expectState int

const (
	expAny        expectState = iota
	expCollection             // after verbs like SET/GET/UPDATE/DELETE/DROP/TRUNCATE
	expKey                    // after seeing DOT following a collection
	expValue                  // after COLON or after EQ in TTI=...
)

// helper bytes for quick lookahead comparisons
const (
	// Dot divides a collection from its keys as a SQL-like statement would
	// For example SELECT * FROM db.table here matches GET collection.key;
	chDot = '.'

	// Colon is a key-value separator
	// This works both for JSON values and SCL statements, as in
	// - { "key": "value" }
	// - SET collection.key: "value"
	chColon = ':'

	// Semicolon is used to terminate statements and it's mandatory
	chSemicolon = ';'

	// EQ sign is used for the TTI (time to invalidate) of each key
	chEq = '='

	// Asterisk is used for wildcard matching and selection/deletion of all keys
	// GET users.*; returns all the key:value pairs in the users collection
	// DROP users; reflects the behavior of GET users.*; + Loop[id] (DELETE users.{id})
	chAsterisk = '*'
)

func (l *lexer) Lex(lval *yySymType) int {
	for {
		l.skipSpaces()

		if l.eof() {
			return 0
		}

		r := l.peek()

		// Identifiers / keywords (allow letters and underscore to start)
		if unicode.IsLetter(r) || r == '_' {
			word := l.readIdent()
			upper := strings.ToUpper(word)

			// Reserved keywords for statements and configuration (like TTI) but also T/F values
			switch upper {
			case "SET":
				l.lastToken = SET
				l.expect = expCollection
				return SET
			case "GET":
				l.lastToken = GET
				l.expect = expCollection
				return GET
			case "UPDATE":
				l.lastToken = UPDATE
				l.expect = expCollection
				return UPDATE
			case "DELETE":
				l.lastToken = DELETE
				l.expect = expCollection
				return DELETE
			case "DROP":
				l.lastToken = DROP
				l.expect = expCollection
				return DROP
			case "TRUNCATE":
				l.lastToken = TRUNCATE
				l.expect = expCollection
				return TRUNCATE
			case "TRUE":
				l.lastToken = TRUE
				return TRUE
			case "FALSE":
				l.lastToken = FALSE
				return FALSE
			case "NULL":
				l.lastToken = NULL
				return NULL
			case "TTI":
				// e.g. UPDATE coll.key TTI=123
				l.lastToken = TTI
				return TTI
			}

			// Not a reserved word → classify by context/state
			lval.str = word

			switch l.expect {
			case expCollection:
				// After a verb we expect a collection name (like SET collection.key)
				l.lastToken = COLLECTION
				// After the collection, usually we will see DOT (which goes to a KEY),
				// but we keep the state loose; DOT will set expKey when it comes.
				// We can also expect a semicolon for actions like DROP/TRUNCATE (i.e. DROP collection;)
				l.expect = expAny
				return COLLECTION
			case expKey:
				// We just saw a DOT. This identifier is a key in a statement
				l.lastToken = KEY
				l.expect = expAny
				return KEY
			default:
				// if next char is a DOT we'll treat as collection,
				// if it was after a DOT, then it's a key,
				// if next is ':' -> key,
				// if next is ';' -> collection (e.g. DROP users;)
				switch l.peekNextNonSpace() {
				case chDot:
					l.lastToken = COLLECTION
					return COLLECTION
				case chColon:
					l.lastToken = KEY
					return KEY
				case chSemicolon, 0:
					l.lastToken = COLLECTION
					return COLLECTION
				default:
					// Fallback: treat as KEY (common in DELETE/GET without colon after the key)
					// TODO: analyse specific cases where this behavior does not apply
					l.lastToken = KEY
					return KEY
				}
			}
		}

		// Numbers (basic JSON-style: optional leading '-', digits, optional decimal)
		// TODO: analyse if carrot-case applies to numbers or we should allow snake_case (camelCase already works)
		if r == '-' || unicode.IsDigit(r) {
			num := l.readNumber()
			lval.str = num
			if l.expect == expKey {
				l.lastToken = KEY
				l.expect = expAny
				return KEY
			}
			l.lastToken = NUMBER
			return NUMBER
		}

		// Strings: support basic escapes \" \\ \n \t \r
		if r == '"' {
			str, err := l.readString()
			if err != nil {
				panic(err)
			}
			lval.str = str
			l.lastToken = STRING
			return STRING
		}

		// Single-character punctuation / operators
		switch r {
		case '{':
			l.advance()
			l.lastToken = LBRACE
			return LBRACE
		case '}':
			l.advance()
			l.lastToken = RBRACE
			return RBRACE
		case '[':
			l.advance()
			l.lastToken = LBRACK
			return LBRACK
		case ']':
			l.advance()
			l.lastToken = RBRACK
			return RBRACK
		case ',':
			l.advance()
			l.lastToken = COMMA
			return COMMA
		case chColon:
			l.advance()
			l.lastToken = COLON
			l.expect = expValue
			return COLON
		case chEq:
			l.advance()
			l.lastToken = EQ
			// If we just saw TTI, the next token(s) form a value
			if l.lastToken == TTI {
				l.expect = expValue
			} else {
				l.expect = expValue
			}
			return EQ
		case chDot:
			l.advance()
			l.lastToken = DOT
			l.expect = expKey
			return DOT
		case chAsterisk:
			l.advance()
			l.lastToken = ASTERISK
			return ASTERISK
		case chSemicolon:
			l.advance()
			l.lastToken = SEMICOLON
			l.expect = expAny
			return SEMICOLON
		default:
			panic(fmt.Sprintf("unexpected character: %q at pos %d", r, l.pos))
		}
	}
}

func (l *lexer) Error(s string) {
	fmt.Println("Parse error:", s)
}

/* ---------------- helpers ---------------- */

func (l *lexer) eof() bool {
	return l.pos >= len(l.input)
}

func (l *lexer) peek() rune {
	if l.eof() {
		return 0
	}
	return rune(l.input[l.pos])
}

func (l *lexer) advance() {
	if !l.eof() {
		l.pos++
	}
}

func (l *lexer) skipSpaces() {
	for !l.eof() {
		r := rune(l.input[l.pos])
		if !unicode.IsSpace(r) {
			return
		}
		l.pos++
	}
}

func (l *lexer) peekNextNonSpace() byte {
	i := l.pos
	for i < len(l.input) {
		c := l.input[i]
		if !unicode.IsSpace(rune(c)) {
			return c
		}
		i++
	}
	return 0
}

func (l *lexer) readIdent() string {
	start := l.pos
	// first rune already checked in caller
	l.advance()
	for !l.eof() {
		r := rune(l.input[l.pos])
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			l.advance()
			continue
		}
		break
	}
	return l.input[start:l.pos]
}

func (l *lexer) readNumber() string {
	start := l.pos
	// optional leading '-'
	// TODO: Do I really need this? (references comment above)
	if l.peek() == '-' {
		l.advance()
	}
	digits := false
	for !l.eof() && unicode.IsDigit(l.peek()) {
		l.advance()
		digits = true
	}
	// optional fractional part
	if !l.eof() && l.peek() == '.' {
		l.advance()
		for !l.eof() && unicode.IsDigit(l.peek()) {
			l.advance()
			digits = true
		}
	}
	if !digits {
		panic(fmt.Errorf("invalid number at pos %d", start))
	}
	return l.input[start:l.pos]
}

func (l *lexer) readString() (string, error) {
	// consume opening quote
	l.advance()
	var out []rune
	for !l.eof() {
		r := l.peek()
		l.advance()
		switch r {
		case '"':
			return string(out), nil
		case '\\':
			if l.eof() {
				return "", fmt.Errorf("unterminated escape")
			}
			er := l.peek()
			l.advance()
			switch er {
			case '"':
				out = append(out, '"')
			case '\\':
				out = append(out, '\\')
			case 'n':
				out = append(out, '\n')
			case 't':
				out = append(out, '\t')
			case 'r':
				out = append(out, '\r')
			default:
				// keep unknown escape literally
				out = append(out, '\\', er)
			}
		default:
			out = append(out, r)
		}
	}
	return "", fmt.Errorf("unterminated string")
}
