package _type

// RequireObjectCoercible :
// Returns a normal completion with the argument if it can be converted to an object.
// https://tc39.es/ecma262/#sec-requireobjectcoercible
// TODO
func RequireObjectCoercible(argument *ECMAPrimitive) *CompletionRecord {
	if argument.Type != UndefinedType_ && argument.Type != NullType_ {
		return NewCompletionRecord(NormalCompletion, argument, "")
	}
	// TODO throw a TypeError for null/undefined
	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// IsArray :
// Returns a normal completion with a boolean if the argument can be identified as an array.
// https://tc39.es/ecma262/#sec-isarray
// TODO
func IsArray(argument *ECMAPrimitive) *CompletionRecord {
	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// IsCallable :
// Returns a normal completion with a boolean if the argument is a callable function
// with a call internal method.
// https://tc39.es/ecma262/#sec-iscallable
// TODO
func IsCallable(argument *ECMAPrimitive) *CompletionRecord {
	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// IsConstructor :
// Returns a normal completion with a boolean if the argument is a function object with
// a constructor function.
// https://tc39.es/ecma262/#sec-isconstructor
// TODO
func IsConstructor(argument *ECMAPrimitive) *CompletionRecord {
	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// IsExtensible :
// Determine whether it is permitted to add additional properties to this object.
// https://tc39.es/ecma262/#table-essential-internal-methods
// https://tc39.es/ecma262/#sec-isextensible-o
func IsExtensible(target *ECMAObject) *CompletionRecord {
	// INVARIANTS---
	return NewCompletionRecord(NormalCompletion, target.InternalSlots.Extensible, "")
	// ---INVARIANTS
}

// IsIntegralNumber :
// Returns a normal completion with a boolean if the argument is a finite integral value.
// https://tc39.es/ecma262/#sec-isintegralnumber
// TODO
func IsIntegralNumber(argument *ECMAPrimitive) *CompletionRecord {
	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// IsPropertyKey :
// Returns a normal completion with a boolean if the argument may be used as a property key.
// https://tc39.es/ecma262/#sec-ispropertykey
// TODO
func IsPropertyKey(argument *ECMAPrimitive) *CompletionRecord {
	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// IsRegExp :
// Returns a normal completion with a boolean if the argument is a regular expression.
// https://tc39.es/ecma262/#sec-isregexp
// TODO
func IsRegExp(argument *ECMAPrimitive) *CompletionRecord {
	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// IsStringWellFormedUnicode :
// Checks if str is a well formed UTF-16 sequence.
// https://tc39.es/ecma262/#sec-isstringwellformedunicode
// TODO
func IsStringWellFormedUnicode(str string) (bool, error) {
	return false, nil
}

// SameValue :
// Check if two values are equal to each other.
// https://tc39.es/ecma262/#sec-samevalue
// TODO
func SameValue(x *ECMAPrimitive, y *ECMAPrimitive) (bool, error) {
	if x.Type != y.Type {
		return false, nil
	}
	if x.Type == NumberType_ {
		// TODO
		return false, nil
	}

	if x.Type == BigIntType_ {
		// TODO
		return false, nil
	}
	return SameValueNonNumeric(x, y)
}

// SameValueZero :
// Check if two values are equal to each other, except it treats +0 and -0 as the same.
// https://tc39.es/ecma262/#sec-samevaluezero
// TODO
func SameValueZero(x *ECMAPrimitive, y *ECMAPrimitive) (bool, error) {
	if x.Type != y.Type {
		return false, nil
	}
	if x.Type == NumberType_ {
		// TODO
		return false, nil
	}

	if x.Type == BigIntType_ {
		// TODO
		return false, nil
	}
	return SameValueNonNumeric(x, y)
}

// SameValueNonNumeric :
// Check if two non-numeric values are equal to each other.
// https://tc39.es/ecma262/#sec-samevaluenonnumeric
// TODO
func SameValueNonNumeric(x *ECMAPrimitive, y *ECMAPrimitive) (bool, error) {
	// handle undefined & null types
	if x.Type == UndefinedType_ || x.Type == NullType_ {
		return true, nil
	}

	// handle string type
	if x.Type == StringType_ {
		// TODO : does go string allow ecma like strings
		return x.Value.(string) == y.Value.(string), nil
	}

	// handle boolean type
	if x.Type == BooleanType_ {
		return x.Value.(bool) == y.Value.(bool), nil
	}

	// handle symbol type
	if x.Type == SymbolType_ {
		// TODO
		return false, nil
	}

	// otherwise: handle object type
	// TODO
	return false, nil
}

// IsLessThan :
// Returns a normal completion with a boolean or undefined. This function is the basis for x < y.
// https://tc39.es/ecma262/#sec-islessthan
// TODO
func IsLessThan(x *ECMAPrimitive, y *ECMAPrimitive, leftFirst bool) *CompletionRecord {
	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// IsLooselyEqual :
// Returns a normal completion with a boolean. This function is the basis for x == y.
// https://tc39.es/ecma262/#sec-islooselyequal
// TODO
func IsLooselyEqual(x *ECMAPrimitive, y *ECMAPrimitive) *CompletionRecord {
	return NewCompletionRecord(ThrowCompletion, nil, "")
}

// IsStrictlyEqual :
// Returns a normal completion with a boolean. This function is the basis for x === y.
// https://tc39.es/ecma262/#sec-isstrictlyequal
// TODO
func IsStrictlyEqual(x *ECMAPrimitive, y *ECMAPrimitive) *CompletionRecord {
	return NewCompletionRecord(ThrowCompletion, nil, "")
}
