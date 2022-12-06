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
		buf:   make([]byte, 14),
		index: 14,
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
		for i := 0; i < len(m.buf)-1; i++ {
			for j := i + 1; j < len(m.buf); j++ {
				if m.buf[i] == m.buf[j] {
					goto NextByte
				}
			}
		}

		fmt.Println("Found", string(m.buf))
		return m.index, nil

	NextByte:
		data := make([]byte, 1)
		_, err := m.r.Read(data)
		if err != nil {
			return 0, err
		}

		m.buf = append(m.buf[1:], data...)
		m.index++
	}
}
