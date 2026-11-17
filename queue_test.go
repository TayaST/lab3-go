package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	_, ok := q.Front()
	assert.False(t, ok) // Пустая очередь
	_, ok = q.Pop()
	assert.False(t, ok)

	q.Push("A")
	q.Push("B")
	assert.Equal(t, "A B", q.String())

	val, ok := q.Front()
	assert.True(t, ok)
	assert.Equal(t, "A", val)

	q.Pop()
	val, _ = q.Front()
	assert.Equal(t, "B", val)

	// Сериализация
	q.Deserialize("1 2 3")
	assert.Equal(t, "1 2 3", q.String())
}