package requestreply

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type BackendPubsubMarshaler[Result any] interface {
	MarshalReply(params BackendOnCommandProcessedParams[Result]) (*message.Message, error)
	UnmarshalReply(msg *message.Message) (reply Reply[Result], err error)
}

const (
	ErrorMetadataKey    = "_watermill_requestreply_error"
	HasErrorMetadataKey = "_watermill_requestreply_has_error"
)

type BackendPubsubJSONMarshaler[Result any] struct{}

func (m BackendPubsubJSONMarshaler[Result]) MarshalReply(
	params BackendOnCommandProcessedParams[Result],
) (*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m BackendPubsubJSONMarshaler[Result]) UnmarshalReply(msg *message.Message) (Reply[Result], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
