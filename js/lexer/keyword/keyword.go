package keyword

type ReservedKeyword string

const (
	AWAIT      ReservedKeyword = "await"
	BREAK      ReservedKeyword = "break"
	CASE       ReservedKeyword = "case"
	CATCH      ReservedKeyword = "catch"
	CLASS      ReservedKeyword = "class"
	CONST      ReservedKeyword = "const"
	CONTINUE   ReservedKeyword = "continue"
	DEBUGGER   ReservedKeyword = "debugger"
	DEFAULT    ReservedKeyword = "default"
	DELETE     ReservedKeyword = "delete"
	DO         ReservedKeyword = "do"
	ELSE       ReservedKeyword = "else"
	ENUM       ReservedKeyword = "enum"
	EXPORT     ReservedKeyword = "export"
	EXTENDS    ReservedKeyword = "extends"
	FALSE      ReservedKeyword = "false"
	FINALLY    ReservedKeyword = "finally"
	FOR        ReservedKeyword = "for"
	FUNCTION   ReservedKeyword = "function"
	IF         ReservedKeyword = "if"
	IMPORT     ReservedKeyword = "import"
	IN         ReservedKeyword = "in"
	INSTANCEOF ReservedKeyword = "instanceof"
	NEW        ReservedKeyword = "new"
	NULL       ReservedKeyword = "null"
	RETURN     ReservedKeyword = "return"
	SUPER      ReservedKeyword = "super"
	SWITCH     ReservedKeyword = "switch"
	THIS       ReservedKeyword = "this"
	THROW      ReservedKeyword = "throw"
	TRUE       ReservedKeyword = "true"
	TRY        ReservedKeyword = "try"
	TYPEOF     ReservedKeyword = "typeof"
	VAR        ReservedKeyword = "var"
	VOID       ReservedKeyword = "void"
	WHILE      ReservedKeyword = "while"
	WITH       ReservedKeyword = "with"
	YIELD      ReservedKeyword = "yield"
	RKNOTFOUND ReservedKeyword = ""
)

func GetReservedKeyword(keyword string) ReservedKeyword {
	switch keyword {
	case "await":
		return AWAIT
	case "break":
		return BREAK
	case "case":
		return CASE
	case "catch":
		return CATCH
	case "class":
		return CLASS
	case "const":
		return CONST
	case "continue":
		return CONTINUE
	case "debugger":
		return DEBUGGER
	case "default":
		return DEFAULT
	case "delete":
		return DELETE
	case "do":
		return DO
	case "else":
		return ELSE
	case "enum":
		return ENUM
	case "export":
		return EXPORT
	case "extends":
		return EXTENDS
	case "false":
		return FALSE
	case "finally":
		return FINALLY
	case "for":
		return FOR
	case "function":
		return FUNCTION
	case "if":
		return IF
	case "import":
		return IMPORT
	case "in":
		return IN
	case "instanceof":
		return INSTANCEOF
	case "new":
		return NEW
	case "null":
		return NULL
	case "return":
		return RETURN
	case "super":
		return SUPER
	case "switch":
		return SWITCH
	case "this":
		return THIS
	case "throw":
		return THROW
	case "true":
		return TRUE
	case "try":
		return TRY
	case "typeof":
		return TYPEOF
	case "var":
		return VAR
	case "void":
		return VOID
	case "while":
		return WHILE
	case "with":
		return WITH
	case "yield":
		return YIELD
	}
	return RKNOTFOUND
}

type StrictReservedKeyword string

const (
	LET         StrictReservedKeyword = "let"
	STATIC      StrictReservedKeyword = "static"
	IMPLEMENTS  StrictReservedKeyword = "implements"
	INTERFACE   StrictReservedKeyword = "interface"
	PACKAGE     StrictReservedKeyword = "package"
	PRIVATE     StrictReservedKeyword = "private"
	PROTECTED   StrictReservedKeyword = "protected"
	PUBLIC      StrictReservedKeyword = "public"
	SRKNOTFOUND StrictReservedKeyword = ""
)

func GetStrictReservedKeyword(keyword string) StrictReservedKeyword {
	switch keyword {
	case "let":
		return LET
	case "static":
		return STATIC
	case "implements":
		return IMPLEMENTS
	case "interface":
		return INTERFACE
	case "package":
		return PACKAGE
	case "private":
		return PRIVATE
	case "protected":
		return PROTECTED
	case "public":
		return PUBLIC
	}
	return SRKNOTFOUND
}
