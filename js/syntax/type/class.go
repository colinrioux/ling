package _type

// ClassFieldDefinition :
// A type to represent class fields.
// https://tc39.es/ecma262/#sec-classfielddefinition-record-specification-type
type ClassFieldDefinition struct {
	Name        string
	Initializer *ECMAObject // TODO function object
}

// NewClassFieldDefinitionRecord :
// Creates a new ClassFieldDefinition.
func NewClassFieldDefinitionRecord(name string, initializer *ECMAObject) *ClassFieldDefinition {
	return &ClassFieldDefinition{
		Name:        name,
		Initializer: initializer,
	}
}

// ClassStaticBlockDefinition :
// A type used to encapsulate code for a class static execution block.
// https://tc39.es/ecma262/#sec-classstaticblockdefinition-record-specification-type
type ClassStaticBlockDefinition struct {
	BodyFunction *ECMAObject // TODO function object
}

// NewClassStaticBlockDefinition :
// Create a new ClassStaticBlockDefinition.
func NewClassStaticBlockDefinition(bodyFunction *ECMAObject) *ClassStaticBlockDefinition {
	return &ClassStaticBlockDefinition{
		BodyFunction: bodyFunction,
	}
}
