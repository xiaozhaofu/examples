package io_interfaces

import (
	"fmt"
	"io"
	"os"
)

func Copy(in io.ReadSeeker, out io.Writer) error {
	w := io.MultiWriter(out, os.Stdout)
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	in.Seek(0, 0)

	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}
	fmt.Println()
	return nil
}

func PipeExample() error {
	r, w := io.Pipe()

	go func() {
		w.Write([]byte("test\n"))
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {

	}
}
