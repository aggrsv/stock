package conf

import (
	"io/ioutil"
	"os"
)

// Config stores in file.

func ReadFile(path string, v interface{}, decoder Decoder) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return decoder.NewDecoder(file).Decode(v)
}

func RestoreFile(path string, data interface{}, encoder Encoder) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	return encoder.NewEncoder(file).Encode(data)
}

func ReadText(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	d, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(d)
}
