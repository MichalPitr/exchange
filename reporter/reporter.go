package reporter

import (
	"bufio"
	"fmt"
	"os"
)

type Reporter struct {
	file   *os.File
	writer *bufio.Writer
}

func New(filename string) (*Reporter, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	writer := bufio.NewWriter(file)
	return &Reporter{file: file, writer: writer}, nil
}

func (r *Reporter) Println(s string) error {
	_, err := fmt.Fprintln(r.writer, s)
	return err
}

func (r *Reporter) Flush() error {
	return r.writer.Flush()
}

func (r *Reporter) Close() error {
	if err := r.Flush(); err != nil {
		return err
	}
	return r.file.Close()
}
