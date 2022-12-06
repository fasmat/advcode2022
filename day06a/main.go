package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	m, err := NewMarkerReader(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(m.NextMarker())
}

type MarkerReader struct {
	buf   []byte
	index int

	r io.Reader
}

func NewMarkerReader(r io.Reader) (*MarkerReader, error) {
	m := &MarkerReader{
		buf:   make([]byte, 4),
		index: 4,
		r:     r,
	}

	_, err := m.r.Read(m.buf)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MarkerReader) NextMarker() (int, error) {
	for {
		if m.buf[0] != m.buf[1] && m.buf[0] != m.buf[2] && m.buf[0] != m.buf[3] &&
			m.buf[1] != m.buf[2] && m.buf[1] != m.buf[3] &&
			m.buf[2] != m.buf[3] {
			fmt.Println("Found", string(m.buf))
			return m.index, nil
		}

		data := make([]byte, 1)
		_, err := m.r.Read(data)
		if err != nil {
			return 0, err
		}

		m.buf = append(m.buf[1:], data...)
		m.index++
	}
}
