package json

import (
	"io"

	"h12.io/mongodb-json/json"
)

func NewDecoder(r io.Reader) *json.Decoder {
	d := json.NewDecoder(r)
	d.Extend(&jsonExt)
	return d
}
