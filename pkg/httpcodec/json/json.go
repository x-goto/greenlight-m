package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type jsonCodec struct {
}

func New() *jsonCodec {
	return &jsonCodec{}
}

func (j *jsonCodec) Encode(data any) ([]byte, error) {
	js, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	js = append(js, '\n')

	return js, nil
}

func (j *jsonCodec) ContentType() string {
	return "application/json"
}

func (j *jsonCodec) Decode(r io.ReadCloser, dst any) error {
	err := json.NewDecoder(r).Decode(dst)
	if err != nil {
		var SyntaxError *json.SyntaxError
		var UnmarshalTypeError *json.UnmarshalTypeError
		var InvalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &SyntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", SyntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly formed JSON")
		case errors.As(err, &UnmarshalTypeError):
			if UnmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", UnmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type(at character %d)", UnmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			//in case if r.Close() happens
			return errors.New("body must not be empty")
		case errors.As(err, &InvalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	return nil
}
