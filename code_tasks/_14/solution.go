package _14

import "reflect"

type Kind int

const (
	Int Kind = iota
	String
	Bool
	Chan
	Undefined
)

func KindOf(v any) Kind {
	switch v.(type) {
	case int:
		return Int
	case string:
		return String
	case bool:
		return Bool
	}
	if v != nil && reflect.TypeOf(v).Kind() == reflect.Chan {
		return Chan
	}
	return Undefined
}
