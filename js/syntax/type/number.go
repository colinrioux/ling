package _type

import "math"

// ECMANumber
// Primitive ECMA number.
type ECMANumber struct {
	*ECMAPrimitive
}

// NewPrimitiveNumber :
// Create a new primitive number.
func NewPrimitiveNumber(identifier string, value interface{}) *ECMANumber {
	return &ECMANumber{
		ECMAPrimitive: NewPrimitive3(NumberType_, identifier, value),
	}
}

// ToBool :
// Convert this primitive into a boolean value.
// https://developer.mozilla.org/en-US/docs/Glossary/Falsy
func (p *ECMANumber) ToBool() (bool, error) {
	v := p.ECMAPrimitive.GetValue().(float64)
	return v != 0 && !math.IsNaN(v), nil
}

//
//import (
//	"duck/ling/js/syntax/property"
//	"fmt"
//	"math"
//)
//
//type ECMANumber struct {
//	_type      ECMAType
//	value      float64
//	attributes property.SimpleAttribute
//}
//
//// GetType :
//// Get the ECMAType of this ECMANumber.
//func (n *ECMANumber) GetType() ECMAType {
//	return n._type
//}
//
//// GetValue :
//// Get the value of this ECMANumber.
//func (n *ECMANumber) GetValue() float64 {
//	return n.value
//}
//
//// GetAttributes :
//// Get the attributes of this ECMANumber.
//func (n *ECMANumber) GetAttributes() property.SimpleAttribute {
//	return n.attributes
//}
//
//// Is :
//// Checks if the value of this ECMANumber is equal to another ECMANumber.
//func (n *ECMANumber) Is(other *ECMANumber) bool {
//	return n.value == other.GetValue()
//}
//
//// IsNaN :
//// Checks if this ECMANumber is equal to NaN.
//// https://tc39.es/ecma262/multipage/numbers-and-dates.html#sec-number.isnan
//func (n *ECMANumber) IsNaN() bool {
//	return n.value == NumberNaN.GetValue()
//}
//
//// IsPositiveInfinity :
//// Checks if this ECMANumber is considered +Inf.
//func (n *ECMANumber) IsPositiveInfinity() bool {
//	return n.value == NumberPositiveInfinity.GetValue() || n.value > NumberMaxValue.GetValue()
//}
//
//// IsNegativeInfinity :
//// Checks if this ECMANumber is considered -Inf.
//func (n *ECMANumber) IsNegativeInfinity() bool {
//	return n.value == NumberNegativeInfinity.GetValue() || n.value < NumberMinValue.GetValue()
//}
//
//// IsFinite :
//// Checks if this ECMANumber is finite.
//// https://tc39.es/ecma262/multipage/numbers-and-dates.html#sec-number.isfinite
//func (n *ECMANumber) IsFinite() bool {
//	return !n.IsNaN() && !n.IsPositiveInfinity() && !n.IsNegativeInfinity()
//}
//
//// IsIntegralNumber :
//// Checks if this ECMANumber is an integer.
//// https://tc39.es/ecma262/multipage/abstract-operations.html#sec-isintegralnumber
//func (n *ECMANumber) IsIntegralNumber() bool {
//	if n.IsNaN() || n.IsPositiveInfinity() || n.IsNegativeInfinity() {
//		return false
//	}
//	return math.Floor(math.Abs(n.value)) == math.Abs(n.value)
//}
//
//// IsSafeInteger :
//// Checks if this ECMANumber is considered a safe integer.
//// https://tc39.es/ecma262/multipage/numbers-and-dates.html#sec-number.issafeinteger
//func (n *ECMANumber) IsSafeInteger() bool {
//	if n.IsIntegralNumber() {
//		return math.Abs(n.value) <= (2 ^ 53 - 1)
//	}
//	return false
//}
//
//func (n *ECMANumber) String() string {
//	return fmt.Sprintf("Number(%v)%v", n.value, n.attributes.String())
//}
//
//// NewECMANumber1 :
//// Constructs a new ECMANumber with default attribute.
//func NewECMANumber1(value float64) *ECMANumber {
//	return &ECMANumber{
//		_type:      NumberType,
//		value:      value,
//		attributes: property.SimpleAttribute{},
//	}
//}
//
//// NewECMANumber2 :
//// Constructs a new ECMANumber with customized attribute.
//func NewECMANumber2(value float64, attributes property.SimpleAttribute) *ECMANumber {
//	return &ECMANumber{
//		_type:      NumberType,
//		value:      value,
//		attributes: attributes,
//	}
//}
//
//// NumberEpsilon :
//// Representation for epsilon in ECMA.
//var NumberEpsilon = NewECMANumber2(
//	2.2204460492503130808472633361816e-16,
//	property.SimpleAttribute{
//		Writable:     false,
//		Enumerable:   false,
//		Configurable: false,
//	},
//)
//
//// NumberMaxSafeInteger :
//// Largest integral number in ECMA.
//var NumberMaxSafeInteger = NewECMANumber2(
//	9007199254740991,
//	property.SimpleAttribute{
//		Writable:     false,
//		Enumerable:   false,
//		Configurable: false,
//	},
//)
//
//// NumberMaxValue :
//// Largest number in ECMA.
//var NumberMaxValue = NewECMANumber2(
//	1.7976931348623157e308,
//	property.SimpleAttribute{
//		Writable:     false,
//		Enumerable:   false,
//		Configurable: false,
//	},
//)
//
//// NumberMinSafeInteger :
//// Smallest integral number in ECMA.
//var NumberMinSafeInteger = NewECMANumber2(
//	-9007199254740991,
//	property.SimpleAttribute{
//		Writable:     false,
//		Enumerable:   false,
//		Configurable: false,
//	},
//)
//
//// NumberMinValue :
//// Smallest number in ECMA.
//var NumberMinValue = NewECMANumber2(
//	5e-324,
//	property.SimpleAttribute{
//		Writable:     false,
//		Enumerable:   false,
//		Configurable: false,
//	},
//)
//
//// NumberNegativeInfinity :
//// -Inf in ECMA.
//var NumberNegativeInfinity = NewECMANumber2(
//	math.Inf(-1),
//	property.SimpleAttribute{
//		Writable:     false,
//		Enumerable:   false,
//		Configurable: false,
//	},
//)
//
//// NumberPositiveInfinity :
//// +Inf in ECMA.
//var NumberPositiveInfinity = NewECMANumber2(
//	math.Inf(+1),
//	property.SimpleAttribute{
//		Writable:     false,
//		Enumerable:   false,
//		Configurable: false,
//	},
//)
//
//// NumberNaN :
//// NaN in ECMA.
//var NumberNaN = NewECMANumber2(
//	math.NaN(),
//	property.SimpleAttribute{
//		Writable:     false,
//		Enumerable:   false,
//		Configurable: false,
//	},
//)
