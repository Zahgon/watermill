package cqrs

func FullyQualifiedStructName(v interface{}) string { _ = "STUB: not implemented"; return "" }

func StructName(v interface{}) string { _ = "STUB: not implemented"; return "" }

type namedStruct interface {
	Name() string
}

func NamedStruct(fallback func(v interface{}) string) func(v interface{}) string {
	_ = "STUB: not implemented"
	return nil
}
