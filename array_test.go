package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDynamicArray(t *testing.T) {
	arr := NewArray()
	assert.Equal(t, 0, arr.Len())

	// Тест PushBack
	arr.PushBack("A")
	arr.PushBack("B")
	assert.Equal(t, 2, arr.Len())

	// Тест Insert
	arr.Insert(1, "X")
	assert.Equal(t, 3, arr.Len())
	val, ok := arr.Get(1)
	assert.True(t, ok)
	assert.Equal(t, "X", val)

	// Тест Set
	arr.Set(0, "Z")
	val, _ = arr.Get(0)
	assert.Equal(t, "Z", val)

	// Тест Erase
	arr.Erase(1)
	assert.Equal(t, 2, arr.Len())
	assert.Equal(t, "Z B", arr.String())

	// Тест выхода за границы (Get)
	_, ok = arr.Get(10)
	assert.False(t, ok)

	// Тест выхода за границы (Set)
	arr.Set(10, "Error") // Не должно изменить массив
	assert.Equal(t, 2, arr.Len())

	// Тест выхода за границы (Erase)
	arr.Erase(10) // Не должно удалить ничего
	assert.Equal(t, 2, arr.Len())

	// Тест выхода за границы (Insert)
	arr.Insert(10, "Error") // Не должно вставить
	assert.Equal(t, 2, arr.Len())

	// Сериализация
	arr.Deserialize("A B C")
	assert.Equal(t, 3, arr.Len())
	assert.Equal(t, "A B C", arr.String())

		// Покрываем отрицательные индексы (ветка false)
	arr.Insert(-1, "Error") // Не должно вставить
	arr.Get(-1)             // Не должно вернуть
	arr.Set(-1, "Error")    // Не должно изменить
	arr.Erase(-1)           // Не должно удалить
	
	assert.Equal(t, 2, arr.Len()) // Размер не изменился
}