package comment

import "duck/ling/js/syntax/unicode"

// SingleLineCommentChar :
// https://tc39.es/ecma262/#prod-SingleLineCommentChar
func SingleLineCommentChar(r rune) bool {
	return unicode.SourceCharacter(r) && !unicode.LineTerminator(r)
}

func SingleLineCommentChars(rs ...rune) bool {

}

func SingleLineComment(rs ...rune) bool {

}