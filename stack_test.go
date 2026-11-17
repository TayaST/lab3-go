package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	st := NewStack()
	_, ok := st.Top()
	assert.False(t, ok)
	_, ok = st.Pop()
	assert.False(t, ok)

	st.Push("A")
	st.Push("B")
	assert.Equal(t, "A B", st.String())

	val, ok := st.Top()
	assert.True(t, ok)
	assert.Equal(t, "B", val)

	st.Pop()
	val, _ = st.Top()
	assert.Equal(t, "A", val)

	// Сериализация
	st.Deserialize("X Y Z")
	assert.Equal(t, "X Y Z", st.String())
}