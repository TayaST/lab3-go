package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestForwardList(t *testing.T) {
	fwd := NewForwardList()
	fwd.PushBack("B")
	fwd.PushFront("A")
	fwd.InsertAfter(0, "X") // A, X, B
	assert.Equal(t, 3, fwd.size)
	assert.Equal(t, "A X B", fwd.String())

	fwd.PopFront() // X, B
	assert.Equal(t, 2, fwd.size)
	
	fwd.EraseAfter(0) // X
	assert.Equal(t, 1, fwd.size)

	assert.Equal(t, 0, fwd.IndexOf("X"))
	assert.Equal(t, -1, fwd.IndexOf("NotExist"))

	// Сериализация
	fwd.Deserialize("Hello World")
	assert.Equal(t, 2, fwd.size)
	assert.Equal(t, "Hello World", fwd.String())
}