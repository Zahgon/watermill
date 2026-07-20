package forwarder

type Marshaler interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

type DefaultMarshaler struct{}

func (DefaultMarshaler) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (DefaultMarshaler) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }
