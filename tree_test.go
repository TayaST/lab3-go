package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBST(t *testing.T) {
	bst := NewBST()
	assert.True(t, bst.IsComplete()) // Пустое дерево полное

	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	assert.Equal(t, 3, bst.size)
	assert.True(t, bst.Contains(5))
	assert.False(t, bst.Contains(20))
	assert.True(t, bst.IsComplete())

	bst.Insert(12) // Нарушает полноту
	assert.False(t, bst.IsComplete())

	// Обходы
	assert.Equal(t, "5 10 12 15", bst.InOrder())
	assert.Equal(t, "10, 5, 15, 12", bst.String()) // PreOrder

	// Сериализация
	bst.Deserialize("10, 5, 15")
	assert.Equal(t, 3, bst.size)
	assert.True(t, bst.Contains(15))
}