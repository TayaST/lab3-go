package main

import (
	"encoding/binary"
	"io"
	"strings"
)

// Приватный узел
type fNode struct {
	value string
	next  *fNode
}

type ForwardList struct {
	head *fNode
	tail *fNode
	size int
}

func NewForwardList() *ForwardList { return &ForwardList{} }

func (l *ForwardList) PushFront(v string) {
	n := &fNode{value: v, next: l.head}
	l.head = n
	if l.tail == nil { l.tail = n }
	l.size++
}

func (l *ForwardList) PushBack(v string) {
	n := &fNode{value: v}
	if l.tail == nil { l.head = n; l.tail = n } else { l.tail.next = n; l.tail = n }
	l.size++
}

func (l *ForwardList) InsertAfter(index int, v string) {
	if index < 0 || index >= l.size { return }
	c := l.head
	for i := 0; i < index; i++ { c = c.next }
	c.next = &fNode{value: v, next: c.next}
	if c == l.tail { l.tail = c.next }
	l.size++
}

func (l *ForwardList) PopFront() {
	if l.head == nil { return }
	l.head = l.head.next
	if l.head == nil { l.tail = nil }
	l.size--
}

func (l *ForwardList) EraseAfter(index int) {
	if index < 0 || index+1 >= l.size { return }
	c := l.head
	for i := 0; i < index; i++ { c = c.next }
	victim := c.next
	c.next = victim.next
	if victim == l.tail { l.tail = c }
	l.size--
}

func (l *ForwardList) IndexOf(v string) int {
	i := 0
	for c := l.head; c != nil; c = c.next {
		if c.value == v { return i }
		i++
	}
	return -1
}

func (l *ForwardList) String() string {
	var sb strings.Builder
	for c := l.head; c != nil; c = c.next {
		sb.WriteString(c.value)
		if c.next != nil { sb.WriteString(" ") }
	}
	return sb.String()
}

func (l *ForwardList) Deserialize(line string) {
	l.head, l.tail, l.size = nil, nil, 0
	for _, p := range strings.Fields(line) { l.PushBack(p) }
}

func (l *ForwardList) SerializeBinary(w io.Writer) error {
	count := uint64(l.size)
	if err := binary.Write(w, binary.LittleEndian, count); err != nil { return err }
	for c := l.head; c != nil; c = c.next {
		length := uint32(len(c.value))
		if err := binary.Write(w, binary.LittleEndian, length); err != nil { return err }
		if _, err := w.Write([]byte(c.value)); err != nil { return err }
	}
	return nil
}

func (l *ForwardList) DeserializeBinary(r io.Reader) error {
	var count uint64
	if err := binary.Read(r, binary.LittleEndian, &count); err != nil { return err }
	l.head, l.tail, l.size = nil, nil, 0
	for i := uint64(0); i < count; i++ {
		var length uint32
		if err := binary.Read(r, binary.LittleEndian, &length); err != nil { return err }
		buf := make([]byte, length)
		if _, err := io.ReadFull(r, buf); err != nil { return err }
		l.PushBack(string(buf))
	}
	return nil
}