package bytessting

import (
	"bytes"
	"io"
)

func Buffer(rawString string) *bytes.Buffer {
	return bytes.NewBufferString(rawString)
}

func toString(r io.Reader) (string, error) {
	b, err := io.ReadAll(r)

	if err != nil {
		return "", err
	}
	return string(b), nil
}
