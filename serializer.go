package serializer

import (
	"encoding/json"
	"io"

	"github.com/vmihailenco/msgpack"
)

type Serializer interface {
	Decode(r io.Reader, v interface{}) error
	Encode(w io.Writer, v interface{}) error
}

type JSONSerializer struct {
	Prefix, Indent string
}

func (JSONSerializer) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func (j JSONSerializer) Encode(w io.Writer, v interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent(j.Prefix, j.Indent)

	return encoder.Encode(v)
}

type MsgPSerializer struct{}

func (MsgPSerializer) Decode(r io.Reader, v interface{}) error {
	return msgpack.NewDecoder(r).Decode(v)
}

func (MsgPSerializer) Encode(w io.Writer, v interface{}) error {
	return msgpack.NewEncoder(w).Encode(v)
}
