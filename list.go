package main

import (
	"encoding/binary"
	"io"
	"strings"
)

// Приватный узел
type lNode struct {
	value string
	prev  *lNode
	next  *lNode
}

type List struct {
	head *lNode
	tail *lNode
	size int
}

func NewList() *List { return &List{} }

func (l *List) PushFront(v string) {
	n := &lNode{value: v, next: l.head}
	if l.head != nil { l.head.prev = n }
	l.head = n
	if l.tail == nil { l.tail = n }
	l.size++
}

func (l *List) PushBack(v string) {
	n := &lNode{value: v, prev: l.tail}
	if l.tail != nil { l.tail.next = n }
	l.tail = n
	if l.head == nil { l.head = n }
	l.size++
}

func (l *List) InsertAfter(index int, v string) {
	if index < 0 || index >= l.size { return }
	c := l.atNode(index)
	n := &lNode{value: v, prev: c, next: c.next}
	if c.next != nil { c.next.prev = n } else { l.tail = n }
	c.next = n
	l.size++
}

func (l *List) PopFront() {
	if l.head == nil { return }
	l.head = l.head.next
	if l.head != nil { l.head.prev = nil } else { l.tail = nil }
	l.size--
}

func (l *List) EraseAfter(index int) {
	if index < 0 || index+1 >= l.size { return }
	victim := l.atNode(index + 1)
	victim.prev.next = victim.next
	if victim.next != nil { victim.next.prev = victim.prev } else { l.tail = victim.prev }
	l.size--
}

func (l *List) atNode(index int) *lNode {
	c := l.head
	for i := 0; i < index; i++ { c = c.next }
	return c
}

func (l *List) IndexOf(v string) int {
	i := 0
	for c := l.head; c != nil; c = c.next {
		if c.value == v { return i }
		i++
	}
	return -1
}

func (l *List) String() string {
	var sb strings.Builder
	for c := l.head; c != nil; c = c.next {
		sb.WriteString(c.value)
		if c.next != nil { sb.WriteString(" ") }
	}
	return sb.String()
}

func (l *List) Deserialize(line string) {
	l.head, l.tail, l.size = nil, nil, 0
	for _, p := range strings.Fields(line) { l.PushBack(p) }
}

func (l *List) SerializeBinary(w io.Writer) error {
	count := uint64(l.size)
	if err := binary.Write(w, binary.LittleEndian, count); err != nil { return err }
	for c := l.head; c != nil; c = c.next {
		length := uint32(len(c.value))
		if err := binary.Write(w, binary.LittleEndian, length); err != nil { return err }
		if _, err := w.Write([]byte(c.value)); err != nil { return err }
	}
	return nil
}

func (l *List) DeserializeBinary(r io.Reader) error {
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