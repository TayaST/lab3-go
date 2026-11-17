package main

import (
	"encoding/binary"
	"io"
	"strings"
)

// Структура с приватным полем data
type DynamicArray struct {
	data []string
}

// Конструктор (аналог new в C++)
func NewArray() *DynamicArray {
	return &DynamicArray{data: make([]string, 0)}
}

func (a *DynamicArray) PushBack(v string) { a.data = append(a.data, v) }

func (a *DynamicArray) Insert(index int, v string) {
	if index >= 0 && index <= len(a.data) {
		a.data = append(a.data, "")
		copy(a.data[index+1:], a.data[index:])
		a.data[index] = v
	}
}

func (a *DynamicArray) Get(index int) (string, bool) {
	if index >= 0 && index < len(a.data) { return a.data[index], true }
	return "", false
}

func (a *DynamicArray) Set(index int, v string) {
	if index >= 0 && index < len(a.data) { a.data[index] = v }
}

func (a *DynamicArray) Erase(index int) {
	if index >= 0 && index < len(a.data) {
		a.data = append(a.data[:index], a.data[index+1:]...)
	}
}

func (a *DynamicArray) Len() int { return len(a.data) }

func (a *DynamicArray) String() string { return strings.Join(a.data, " ") }

func (a *DynamicArray) Deserialize(line string) {
	a.data = []string{}
	for _, p := range strings.Fields(line) { a.data = append(a.data, p) }
}

// SerializeBinary записывает массив в бинарный формат
func (a *DynamicArray) SerializeBinary(w io.Writer) error {
	count := uint64(len(a.data))
	if err := binary.Write(w, binary.LittleEndian, count); err != nil {
		return err
	}
	for _, s := range a.data {
		length := uint32(len(s))
		if err := binary.Write(w, binary.LittleEndian, length); err != nil {
			return err
		}
		if _, err := w.Write([]byte(s)); err != nil {
			return err
		}
	}
	return nil
}

// DeserializeBinary читает массив из бинарного формата
func (a *DynamicArray) DeserializeBinary(r io.Reader) error {
	var count uint64
	if err := binary.Read(r, binary.LittleEndian, &count); err != nil {
		return err
	}
	a.data = make([]string, 0, count)
	for i := uint64(0); i < count; i++ {
		var length uint32
		if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
			return err
		}
		buf := make([]byte, length)
		if _, err := io.ReadFull(r, buf); err != nil {
			return err
		}
		a.data = append(a.data, string(buf))
	}
	return nil
}