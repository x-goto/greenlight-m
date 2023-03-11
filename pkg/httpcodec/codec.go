package httpcodec

import (
	"io"
)

type HTTPCodec interface {
	Encode(data any) ([]byte, error)
	Decode(r io.ReadCloser, dst any) error
	ContentType() string
}
