package conf

import (
	"encoding/json"
	"io"

	"github.com/BurntSushi/toml"
)

type Encoder interface {
	NewEncoder(w io.Writer) Encoder
	Encode(v interface{}) error
}

type Decoder interface {
	NewDecoder(r io.Reader) Decoder
	Decode(v interface{}) error
}

type JSON struct {
	*json.Decoder
	*json.Encoder
}

func (j *JSON) NewEncoder(w io.Writer) Encoder {
	return &JSON{Encoder: json.NewEncoder(w)}
}

func (j *JSON) Encode(v interface{}) error {
	return j.Encoder.Encode(v)
}

func (j *JSON) NewDecoder(r io.Reader) Decoder {
	decoder := &JSON{Decoder: json.NewDecoder(r)}
	decoder.Decoder.UseNumber()
	return decoder
}

func (j *JSON) Decode(v interface{}) error {
	return j.Decoder.Decode(v)
}

type TOML struct {
	*toml.Encoder
	io.Reader
}

func (j *TOML) NewEncoder(w io.Writer) Encoder {
	return &TOML{Encoder: toml.NewEncoder(w)}
}

func (j *TOML) Encode(v interface{}) error {
	return j.Encoder.Encode(v)
}

func (j *TOML) NewDecoder(r io.Reader) Decoder {
	return &TOML{Reader: r}
}

func (j *TOML) Decode(v interface{}) error {
	_, err := toml.DecodeReader(j.Reader, v)
	return err
}
