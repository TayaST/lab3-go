package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T) {
	ht := NewHashTable(5)
	
	_, ok := ht.Get("key1")
	assert.False(t, ok)

	ht.Insert("key1", "val1")
	ht.Insert("key2", "val2")
	assert.Equal(t, 2, ht.size)

	val, ok := ht.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "val1", val)

	// Обновление значения
	ht.Insert("key1", "new_val")
	val, _ = ht.Get("key1")
	assert.Equal(t, "new_val", val)

	ht.Erase("key1")
	_, ok = ht.Get("key1")
	assert.False(t, ok)

	// Сериализация
	ht.Deserialize("a=1, b=2")
	assert.Equal(t, 2, ht.size)
	val, _ = ht.Get("b")
	assert.Equal(t, "2", val)
}