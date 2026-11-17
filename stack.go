package main

import (
	"encoding/binary"
	"io"
	"strings"
)

type Stack struct{ data []string }

func NewStack() *Stack { return &Stack{data: make([]string, 0)} }

func (s *Stack) Push(v string) { s.data = append(s.data, v) }

func (s *Stack) Pop() (string, bool) {
	if len(s.data) == 0 { return "", false }
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}

func (s *Stack) Top() (string, bool) {
	if len(s.data) == 0 { return "", false }
	return s.data[len(s.data)-1], true
}

func (s *Stack) String() string { return strings.Join(s.data, " ") }

func (s *Stack) Deserialize(line string) {
	s.data = []string{}
	for _, p := range strings.Fields(line) { s.data = append(s.data, p) }
}

func (s *Stack) SerializeBinary(w io.Writer) error {
	count := uint64(len(s.data))
	if err := binary.Write(w, binary.LittleEndian, count); err != nil { return err }
	for _, v := range s.data {
		length := uint32(len(v))
		if err := binary.Write(w, binary.LittleEndian, length); err != nil { return err }
		if _, err := w.Write([]byte(v)); err != nil { return err }
	}
	return nil
}

func (s *Stack) DeserializeBinary(r io.Reader) error {
	var count uint64
	if err := binary.Read(r, binary.LittleEndian, &count); err != nil { return err }
	s.data = make([]string, 0, count)
	for i := uint64(0); i < count; i++ {
		var length uint32
		if err := binary.Read(r, binary.LittleEndian, &length); err != nil { return err }
		buf := make([]byte, length)
		if _, err := io.ReadFull(r, buf); err != nil { return err }
		s.data = append(s.data, string(buf))
	}
	return nil
}