package main

import (
	"strconv"
	"testing"
)

// 1. Бенчмарк для массива (DynamicArray)
func BenchmarkArrayPushBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := NewArray()
		for j := 0; j < 10000; j++ {
			arr.PushBack("test")
		}
	}
}

// 2. Бенчмарк для односвязного списка (ForwardList)
func BenchmarkForwardListPushBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fwd := NewForwardList()
		for j := 0; j < 10000; j++ {
			fwd.PushBack("test")
		}
	}
}

// 3. Бенчмарк для двусвязного списка (List)
func BenchmarkListPushBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lst := NewList()
		for j := 0; j < 10000; j++ {
			lst.PushBack("test")
		}
	}
}

// 4. Бенчмарк для очереди (Queue)
func BenchmarkQueuePush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := NewQueue()
		for j := 0; j < 10000; j++ {
			q.Push("test")
		}
	}
}

// 5. Бенчмарк для стека (Stack)
func BenchmarkStackPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		st := NewStack()
		for j := 0; j < 10000; j++ {
			st.Push("test")
		}
	}
}

// 6. Бенчмарк для хеш-таблицы (HashTable)
func BenchmarkHashTableInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ht := NewHashTable(1000)
		for j := 0; j < 1000; j++ {
			ht.Insert("key_"+strconv.Itoa(j), "value")
		}
	}
}

// 7. Бенчмарк для дерева (BST)
func BenchmarkTreeInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bst := NewBST()
		for j := 0; j < 1000; j++ {
			bst.Insert(j)
		}
	}
}

// 8. Бенчмарк для поиска в дереве (BST)
func BenchmarkTreeContains(b *testing.B) {
	bst := NewBST()
	for j := 0; j < 1000; j++ {
		bst.Insert(j)
	}
	
	b.ResetTimer() // Сбрасываем таймер, чтобы не учитывать время на создание дерева
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			bst.Contains(j)
		}
	}
}