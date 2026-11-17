package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// Приватная пара
type pair struct {
	key   string
	value string
}

type HashTable struct {
	buckets [][]pair
	size    int
}

func NewHashTable(bucketsCount int) *HashTable {
	if bucketsCount == 0 { bucketsCount = 1 }
	return &HashTable{buckets: make([][]pair, bucketsCount)}
}

func (h *HashTable) hashFunc(key string) int {
	var hash uint32 = 5381
	for _, c := range key { hash = ((hash << 5) + hash) + uint32(c) }
	return int(hash % uint32(len(h.buckets)))
}

func (h *HashTable) Insert(key, value string) {
	idx := h.hashFunc(key)
	for i, p := range h.buckets[idx] {
		if p.key == key { h.buckets[idx][i].value = value; return }
	}
	h.buckets[idx] = append(h.buckets[idx], pair{key: key, value: value})
	h.size++
}

func (h *HashTable) Get(key string) (string, bool) {
	idx := h.hashFunc(key)
	for _, p := range h.buckets[idx] {
		if p.key == key { return p.value, true }
	}
	return "", false
}

func (h *HashTable) Erase(key string) {
	idx := h.hashFunc(key)
	for i, p := range h.buckets[idx] {
		if p.key == key {
			h.buckets[idx] = append(h.buckets[idx][:i], h.buckets[idx][i+1:]...)
			h.size--
			return
		}
	}
}

func (h *HashTable) String() string {
	var parts []string
	for _, bucket := range h.buckets {
		for _, p := range bucket { parts = append(parts, fmt.Sprintf("%s=%s", p.key, p.value)) }
	}
	return strings.Join(parts, ", ")
}

func (h *HashTable) Deserialize(line string) {
	h.buckets = make([][]pair, len(h.buckets))
	h.size = 0
	for _, pairStr := range strings.Split(line, ",") {
		pairStr = strings.TrimSpace(pairStr)
		if pairStr == "" { continue }
		kv := strings.SplitN(pairStr, "=", 2)
		if len(kv) == 2 { h.Insert(kv[0], kv[1]) }
	}
}

func (h *HashTable) SerializeBinary(w io.Writer) error {
	count := uint64(h.size)
	if err := binary.Write(w, binary.LittleEndian, count); err != nil { return err }
	for _, bucket := range h.buckets {
		for _, p := range bucket {
			klen := uint32(len(p.key))
			binary.Write(w, binary.LittleEndian, klen)
			w.Write([]byte(p.key))
			
			vlen := uint32(len(p.value))
			binary.Write(w, binary.LittleEndian, vlen)
			w.Write([]byte(p.value))
		}
	}
	return nil
}

func (h *HashTable) DeserializeBinary(r io.Reader) error {
	var count uint64
	if err := binary.Read(r, binary.LittleEndian, &count); err != nil { return err }
	h.buckets = make([][]pair, len(h.buckets))
	h.size = 0
	for i := uint64(0); i < count; i++ {
		var klen uint32
		binary.Read(r, binary.LittleEndian, &klen)
		kbuf := make([]byte, klen)
		io.ReadFull(r, kbuf)

		var vlen uint32
		binary.Read(r, binary.LittleEndian, &vlen)
		vbuf := make([]byte, vlen)
		io.ReadFull(r, vbuf)

		h.Insert(string(kbuf), string(vbuf))
	}
	return nil
}