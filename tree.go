package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Приватный узел
type tNode struct {
	value int
	left  *tNode
	right *tNode
}

type BST struct {
	root *tNode
	size int
}

func NewBST() *BST { return &BST{} }

func (t *BST) Insert(value int) { t.root = insertRec(t.root, value, &t.size) }

func insertRec(node *tNode, value int, size *int) *tNode {
	if node == nil { *size++; return &tNode{value: value} }
	if value < node.value { node.left = insertRec(node.left, value, size) } else if value > node.value { node.right = insertRec(node.right, value, size) }
	return node
}

func (t *BST) Contains(value int) bool {
	cur := t.root
	for cur != nil {
		if value == cur.value { return true }
		if value < cur.value { cur = cur.left } else { cur = cur.right }
	}
	return false
}

func (t *BST) IsComplete() bool { return checkComplete(t.root, 0, t.size) }

func checkComplete(node *tNode, index, total int) bool {
	if node == nil { return true }
	if index >= total { return false }
	return checkComplete(node.left, 2*index+1, total) && checkComplete(node.right, 2*index+2, total)
}

func (t *BST) InOrder() string {
	var sb strings.Builder
	inOrderRec(t.root, &sb)
	return strings.TrimSpace(sb.String())
}

func inOrderRec(node *tNode, sb *strings.Builder) {
	if node == nil { return }
	inOrderRec(node.left, sb)
	sb.WriteString(fmt.Sprintf("%d ", node.value))
	inOrderRec(node.right, sb)
}

func (t *BST) String() string {
	var parts []string
	preOrderRec(t.root, &parts)
	return strings.Join(parts, ", ")
}

func preOrderRec(node *tNode, parts *[]string) {
	if node == nil { return }
	*parts = append(*parts, strconv.Itoa(node.value))
	preOrderRec(node.left, parts)
	preOrderRec(node.right, parts)
}

func (t *BST) Deserialize(line string) {
	t.root, t.size = nil, 0
	for _, p := range strings.Split(line, ",") {
		p = strings.TrimSpace(p)
		if p != "" {
			val, _ := strconv.Atoi(p)
			t.Insert(val)
		}
	}
}

// вспомогательная функция для рекурсивной записи
func serializeTreeBinary(node *tNode, w io.Writer) error {
	if node == nil {
		return nil
	}
	if err := binary.Write(w, binary.LittleEndian, int64(node.value)); err != nil {
		return err
	}
	if err := serializeTreeBinary(node.left, w); err != nil {
		return err
	}
	return serializeTreeBinary(node.right, w)
}

// SerializeBinary записывает дерево в бинарный формат
func (t *BST) SerializeBinary(w io.Writer) error {
	count := uint64(t.size)
	if err := binary.Write(w, binary.LittleEndian, count); err != nil {
		return err
	}
	return serializeTreeBinary(t.root, w)
}

// DeserializeBinary читает дерево из бинарного формата
func (t *BST) DeserializeBinary(r io.Reader) error {
	var count uint64
	if err := binary.Read(r, binary.LittleEndian, &count); err != nil {
		return err
	}
	t.root, t.size = nil, 0
	for i := uint64(0); i < count; i++ {
		var val int64
		if err := binary.Read(r, binary.LittleEndian, &val); err != nil {
			return err
		}
		t.Insert(int(val))
	}
	return nil
}