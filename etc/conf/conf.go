package conf

var (
	JSONCodec = &JSON{}
	TOMLCodec = &TOML{}
)

func ReadJSON(path string, v interface{}) error {
	return ReadFile(path, v, JSONCodec)
}

func WriteJSON(path string, v interface{}) error {
	return RestoreFile(path, v, JSONCodec)
}

func ReadTOML(path string, v interface{}) error {
	return ReadFile(path, v, TOMLCodec)
}

func WriteTOML(path string, v interface{}) error {
	return RestoreFile(path, v, TOMLCodec)
}
