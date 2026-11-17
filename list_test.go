package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	lst := NewList()
	lst.PushBack("B")
	lst.PushFront("A")
	lst.InsertAfter(0, "X") // A, X, B
	assert.Equal(t, 3, lst.size)
	assert.Equal(t, "A X B", lst.String())

	lst.PopFront() // X, B
	assert.Equal(t, 2, lst.size)
	
	lst.EraseAfter(0) // X
	assert.Equal(t, 1, lst.size)

	assert.Equal(t, 0, lst.IndexOf("X"))
	assert.Equal(t, -1, lst.IndexOf("NotExist"))

	// Сериализация
	lst.Deserialize("Test Data")
	assert.Equal(t, 2, lst.size)
	assert.Equal(t, "Test Data", lst.String())
}