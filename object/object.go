package object

import "fmt"

type ObjectType string

const (
	INTEGEROBJ   = "INTEGER"
	BOOLEANOBJ   = "BOOLEAN"
	NULLOBJ      = "NULL"
	RETURNVALOBJ = "RETURN_VALUE"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGEROBJ
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEANOBJ
}

type Null struct{}

func (n *Null) Inspect() string {
	return "null"
}

func (n *Null) Type() ObjectType {
	return NULLOBJ
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURNVALOBJ
}
