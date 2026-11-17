package main

import (
	"encoding/binary"
	"io"
	"strings"
)

type Queue struct{ data []string }

func NewQueue() *Queue { return &Queue{data: make([]string, 0)} }

func (q *Queue) Push(v string) { q.data = append(q.data, v) }

func (q *Queue) Pop() (string, bool) {
	if len(q.data) == 0 { return "", false }
	v := q.data[0]
	q.data = q.data[1:]
	return v, true
}

func (q *Queue) Front() (string, bool) {
	if len(q.data) == 0 { return "", false }
	return q.data[0], true
}

func (q *Queue) String() string { return strings.Join(q.data, " ") }

func (q *Queue) Deserialize(line string) {
	q.data = []string{}
	for _, p := range strings.Fields(line) { q.data = append(q.data, p) }
}

func (q *Queue) SerializeBinary(w io.Writer) error {
	count := uint64(len(q.data))
	if err := binary.Write(w, binary.LittleEndian, count); err != nil { return err }
	for _, v := range q.data {
		length := uint32(len(v))
		if err := binary.Write(w, binary.LittleEndian, length); err != nil { return err }
		if _, err := w.Write([]byte(v)); err != nil { return err }
	}
	return nil
}

func (q *Queue) DeserializeBinary(r io.Reader) error {
	var count uint64
	if err := binary.Read(r, binary.LittleEndian, &count); err != nil { return err }
	q.data = make([]string, 0, count)
	for i := uint64(0); i < count; i++ {
		var length uint32
		if err := binary.Read(r, binary.LittleEndian, &length); err != nil { return err }
		buf := make([]byte, length)
		if _, err := io.ReadFull(r, buf); err != nil { return err }
		q.data = append(q.data, string(buf))
	}
	return nil
}