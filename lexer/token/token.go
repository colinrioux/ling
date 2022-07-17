package token

import "fmt"

type Token struct {
	Type  Type
	Value interface{}
}

func (tok Token) String() string {
	return fmt.Sprintf("Token(%s, %v)", tok.Type, tok.Value)
}

func NewToken(_type Type, value interface{}) *Token {
	return &Token{Type: _type, Value: value}
}

func ValueToInt(tok *Token) int {
	res, _ := tok.Value.(int)
	return res
}

type Type uint16

const (
	ILLEGAL Type = iota
	EOF
	SINGLE_LINE_COMMENT // //
	LMULTI_LINE_COMMENT // /*
	RMULTI_LINE_COMMENT // */
	NUMBER
	ADDITION       // +
	SUBTRACTION    // -
	MULTIPLICATION // *
	DIVISION       // /
	LPAREN         // (
	RPAREN         // )
	IDENTIFIER
	KEYWORD
	ASSIGN                 // =
	SEMICOLON              // ;
	LBRACE                 // {
	RBRACE                 // }
	MULTIPLY_ASSIGN        // *=
	DIVISION_ASSIGN        // /=
	MODULO_ASSIGN          // %=
	ADDITION_ASSIGN        // +=
	SUBTRACTION_ASSIGN     // -=
	LSHIFT_ASSIGN          // <<=
	RSHIFT_ASSIGN          // >>=
	URSHIFT_ASSIGN         // <<<=
	BITWISE_AND_ASSIGN     // &=
	BITWISE_XOR_ASSIGN     // ^=
	BITWISE_OR_ASSIGN      // |=
	EXPONENT_ASSIGN        // **=
	MODULO                 // %
	BITWISE_AND            // &
	BITWISE_XOR            // ^
	BITWISE_OR             // |
	BITWISE_NOT            // ~
	COMMA                  // ,
	DECREMENT              // --
	INCREMENT              // ++
	EQUALITY               // ==
	EXPONENT               // **
	GREATER_THAN           // >
	GREATER_THAN_OR_EQUAL  // >=
	LESS_THAN              // <
	LESS_THAN_OR_EQUAL     // <=
	INEQUALITY             // !=
	LSHIFT                 // <<
	RSHIFT                 // >>
	URSHIFT                // >>>
	LOGICAL_AND            // &&
	LOGICAL_AND_ASSIGN     // &&=
	LOGICAL_OR             // ||
	LOGICAL_OR_ASSIGN      // ||=
	LOGICAL_NOT            // !
	NULLISH                // ??
	LOGICAL_NULLISH_ASSIGN // ??=
	OPTIONAL_CHAINING      // ?.
	SPREAD                 // ...
	STRICT_EQUALITY        // ===
	STRICT_INEQUALITY      // !==
	LBRACKET               // [
	RBRACKET               // ]
	DOT                    // .
	DOUBLE_QUOTE           // "
	SINGLE_QUOTE           // '
	BACKTICK               // `
	POUND                  // #
	DOLLAR                 // $
)

var asString = [...]string{
	"UNKNOWN",
	"EOF",
	"SINGLE_LINE_COMMENT",
	"LMULTI_LINE_COMMENT",
	"RMULTI_LINE_COMMENT",
	"NUMBER",
	"ADDITION",
	"SUBTRACTION",
	"MULTIPLICATION",
	"DIVISION",
	"LPAREN",
	"RPAREN",
	"IDENTIFIER",
	"KEYWORD",
	"ASSIGN",
	"SEMICOLON",
	"LBRACE",
	"RBRACE",
	"MULTIPLY_ASSIGN",
	"DIVISION_ASSIGN",
	"MODULO_ASSIGN",
	"ADDITION_ASSIGN",
	"SUBTRACTION_ASSIGN",
	"LSHIFT_ASSIGN",
	"RSHIFT_ASSIGN",
	"URSHIFT_ASSIGN",
	"BITWISE_AND_ASSIGN",
	"BITWISE_XOR_ASSIGN",
	"BITWISE_OR_ASSIGN",
	"EXPONENT_ASSIGN",
	"MODULO",
	"BITWISE_AND",
	"BITWISE_XOR",
	"BITWISE_OR",
	"BITWISE_NOT",
	"COMMA",
	"DECREMENT",
	"INCREMENT",
	"EQUALITY",
	"EXPONENT",
	"GREATER_THAN",
	"GREATER_THAN_OR_EQ",
	"LESS_THAN",
	"LESS_THAN_OR_EQUAL",
	"INEQUALITY",
	"LSHIFT",
	"RSHIFT",
	"URSHIFT",
	"LOGICAL_AND",
	"LOGICAL_AND_ASSIGN",
	"LOGICAL_OR",
	"LOGICAL_OR_ASSIGN",
	"LOGICAL_NOT",
	"NULLISH",
	"LOGICAL_NULLISH_ASSIGN",
	"OPTIONAL_CHAINING",
	"SPREAD",
	"STRICT_EQUALITY",
	"STRICT_INEQUALITY",
	"LBRACKET",
	"RBRACKET",
	"DOT",
	"DOUBLE_QUOTE",
	"SINGLE_QUOTE",
	"BACKTICK",
	"POUND",
	"DOLLAR",
}

func (tok Type) String() string {
	return asString[tok]
}
