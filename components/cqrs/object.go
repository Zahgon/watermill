package cqrs

import (
	"reflect"
)

func isPointer(v interface{}) error { _ = "STUB: not implemented"; return nil }

type NonPointerError struct {
	Type reflect.Type
}

func (e NonPointerError) Error() string { _ = "STUB: not implemented"; return "" }
