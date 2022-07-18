package comment

import (
	"ling/syntax/unicode"
)

// SingleLineCommentChar :
// https://tc39.es/ecma262/#prod-SingleLineCommentChar
func SingleLineCommentChar(r rune) bool {
	return unicode.IsSourceCharacter(r) && !unicode.IsLineTerminator(r)
}

func SingleLineCommentChars(rs ...rune) bool {

}

func SingleLineComment(rs ...rune) bool {

}
