package unicode

import (
	"fmt"
	"ling/util"
	"math"
)

// Format Control Code Points. https://tc39.es/ecma262/#table-format-control-code-point-usage
const (
	ZeroWidthNonJoiner = "\u200C"
	ZWNJ               = ZeroWidthNonJoiner
	ZeroWidthJoiner    = "\u200D"
	ZWJ                = ZeroWidthJoiner
)

// IsWhitespace Code Points. https://tc39.es/ecma262/#table-white-space-code-points
const (
	CharacterTabulation   = "\u0009"
	Tab                   = CharacterTabulation
	LineTabulation        = "\u000B"
	VT                    = LineTabulation
	FormFeed              = "\u000C"
	FF                    = FormFeed
	ZeroWidthNoBreakSpace = "\uFEFF" // Also a format control code point*
	ZWNBSP                = ZeroWidthNoBreakSpace
	UnicodeSpaceSeparator = "\t|\n|\v|\f|\r| |\u0085|\u00A0"
	USP                   = UnicodeSpaceSeparator
)

// IsWhitespace :
// Checks if a rune is considered ECMA whitespace.
// https://tc39.es/ecma262/#prod-WhiteSpace
func IsWhitespace(r rune) bool {
	pattern := fmt.Sprintf("(%s|%s|%s|%s|%s)", Tab, VT, FF, ZWNBSP, USP)
	return util.Match(pattern, string(r))
}

// Line Terminators Code Points. https://tc39.es/ecma262/#sec-line-terminators
const (
	Linefeed           = "\u000A"
	LF                 = Linefeed
	CarriageReturn     = "\u000D"
	CR                 = CarriageReturn
	LineSeparator      = "\u2028"
	LS                 = LineSeparator
	ParagraphSeparator = "\u2029"
	PS                 = ParagraphSeparator
)

// IsLineTerminator :
// Checks if a rune is considered ECMA IsLineTerminator
// https://tc39.es/ecma262/#prod-LineTerminator
func IsLineTerminator(r rune) bool {
	rString := string(r)
	return rString == LF || rString == CR || rString == LS || rString == PS
}

// IsLineTerminatorSequence :
// Checks if a rune and its lookahead is considered ECMA IsLineTerminatorSequence.
// *<CR> <LF> should be considered a single token.
// https://tc39.es/ecma262/#prod-LineTerminatorSequence
func IsLineTerminatorSequence(r rune, lookahead rune) bool {
	rString := string(r)
	if rString == LF || rString == CR || rString == LS || rString == PS {
		return IsLineTerminator(lookahead)
	}

	return false
}

// IsSourceCharacter :
// Checks if a rune is considered ECMA IsSourceCharacter.
// https://tc39.es/ecma262/#prod-SourceCharacter
func IsSourceCharacter(r rune) bool {
	return r >= 0x000 && r <= 0x10FFFF
}

// IsLeadingSurrogate :
// Checks if a rune is considered an utf16 leading surrogate.
// https://tc39.es/ecma262/#leading-surrogate
func IsLeadingSurrogate(r rune) bool {
	return r >= 0xD800 && r <= 0xDBFF
}

// IsTrailingSurrogate :
// Checks if a rune is considered an utf16 trailing surrogate.
// https://tc39.es/ecma262/#leading-surrogate
func IsTrailingSurrogate(r rune) bool {
	return r >= 0xDC00 && r <= 0xDFFF
}

// UTF16EncodeCodePoint :
// https://tc39.es/ecma262/#sec-utf16encodecodepoint
func UTF16EncodeCodePoint(cp rune) string {
	// verify
	if !IsSourceCharacter(cp) {
		return ""
	}

	if cp <= 0xFFFF {
		return string(cp)
	}

	cu1 := rune(math.Floor(float64((cp-0x10000)/0x400))) + 0xD800
	cu2 := ((cp - 0x10000) % 0x400) + 0xDC00
	return string(cu1) + string(cu2)
}

// CodePointsToString :
// https://tc39.es/ecma262/#sec-codepointstostring
func CodePointsToString(text string) string {
	// we don't verify if text is a valid utf16 sequence because UTF16EncodeCodePoint takes care of it.
	result := ""
	for _, cp := range text {
		result += UTF16EncodeCodePoint(cp)
	}
	return result
}

// UTF16SurrogatePairToCodePoint :
// https://tc39.es/ecma262/#sec-utf16decodesurrogatepair
func UTF16SurrogatePairToCodePoint(lead rune, trail rune) rune {
	if !IsLeadingSurrogate(lead) || !IsTrailingSurrogate(trail) {
		return -1
	}

	return (lead-0xD800)*0x400 + (trail - 0xDC00) + 0x10000
}

// CodePointAt :
// https://tc39.es/ecma262/#sec-codepointat
//func CodePointAt(str string, pos int) *_type.UnicodeRecord {
//	size := len(str)
//	if pos < 0 || pos >= size {
//		return nil
//	}
//	first := str[pos]
//	cp := rune(first)
//	if !IsLeadingSurrogate(cp) || !IsTrailingSurrogate(cp) {
//		return &_type.UnicodeRecord{CodePoint: cp, CodeUnitCount: 1, IsUnpairedSurrogate: false}
//	}
//
//	if IsTrailingSurrogate(cp) || pos+1 == size {
//		return &_type.UnicodeRecord{CodePoint: cp, CodeUnitCount: 1, IsUnpairedSurrogate: true}
//	}
//
//	second := str[pos+1]
//	cp2 := rune(second)
//	if !IsTrailingSurrogate(cp2) {
//		return &_type.UnicodeRecord{CodePoint: cp, CodeUnitCount: 1, IsUnpairedSurrogate: true}
//	}
//	cp = UTF16SurrogatePairToCodePoint(cp, cp2)
//	return &_type.UnicodeRecord{CodePoint: cp, CodeUnitCount: 2, IsUnpairedSurrogate: false}
//}

// StringToCodePoints :
// https://tc39.es/ecma262/#sec-stringtocodepoints
//func StringToCodePoints(str string) *[]rune {
//	codePoints := []rune{}
//	size := len(str)
//	position := 0
//	var cp *_type.UnicodeRecord
//	for position < size {
//		cp = CodePointAt(str, position)
//		codePoints = append(codePoints, cp.CodePoint)
//		position = position + int(cp.CodeUnitCount)
//	}
//	return &codePoints
//}

// ParseText :
// https://tc39.es/ecma262/#sec-parsetext
func ParseText(sourceText *[]rune) {

}
