package main

import (
	"bytes"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestArrayBinary(t *testing.T) {
	arr := NewArray(); arr.PushBack("Hello"); arr.PushBack("World")
	var buf bytes.Buffer
	err := arr.SerializeBinary(&buf)
	assert.NoError(t, err)
	var arr2 DynamicArray
	err = arr2.DeserializeBinary(&buf)
	assert.NoError(t, err)
	assert.Equal(t, 2, arr2.Len())
	assert.Equal(t, "Hello", arr2.data[0])
}

func TestForwardListBinary(t *testing.T) {
	fwd := NewForwardList(); fwd.PushBack("A"); fwd.PushBack("B")
	var buf bytes.Buffer
	err := fwd.SerializeBinary(&buf)
	assert.NoError(t, err)
	var fwd2 ForwardList
	err = fwd2.DeserializeBinary(&buf)
	assert.NoError(t, err)
	assert.Equal(t, 2, fwd2.size)
	assert.Equal(t, "A B", fwd2.String())
}

func TestListBinary(t *testing.T) {
	lst := NewList(); lst.PushBack("A"); lst.PushBack("B")
	var buf bytes.Buffer
	err := lst.SerializeBinary(&buf)
	assert.NoError(t, err)
	var lst2 List
	err = lst2.DeserializeBinary(&buf)
	assert.NoError(t, err)
	assert.Equal(t, 2, lst2.size)
}

func TestStackBinary(t *testing.T) {
	st := NewStack(); st.Push("A"); st.Push("B")
	var buf bytes.Buffer
	err := st.SerializeBinary(&buf)
	assert.NoError(t, err)
	var st2 Stack
	err = st2.DeserializeBinary(&buf)
	assert.NoError(t, err)
	assert.Equal(t, "A B", st2.String())
}

func TestQueueBinary(t *testing.T) {
	q := NewQueue(); q.Push("A"); q.Push("B")
	var buf bytes.Buffer
	err := q.SerializeBinary(&buf)
	assert.NoError(t, err)
	var q2 Queue
	err = q2.DeserializeBinary(&buf)
	assert.NoError(t, err)
	assert.Equal(t, "A B", q2.String())
}

func TestHashTableBinary(t *testing.T) {
	ht := NewHashTable(5); ht.Insert("key1", "val1"); ht.Insert("key2", "val2")
	var buf bytes.Buffer
	err := ht.SerializeBinary(&buf)
	assert.NoError(t, err)
	var ht2 HashTable
	ht2.buckets = make([][]pair, 5) // Важно инициализировать бакеты!
	err = ht2.DeserializeBinary(&buf)
	assert.NoError(t, err)
	assert.Equal(t, 2, ht2.size)
	val, ok := ht2.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "val1", val)
}

func TestTreeBinary(t *testing.T) {
	bst := NewBST(); bst.Insert(10); bst.Insert(5); bst.Insert(15)
	var buf bytes.Buffer
	err := bst.SerializeBinary(&buf)
	assert.NoError(t, err)
	var bst2 BST
	err = bst2.DeserializeBinary(&buf)
	assert.NoError(t, err)
	assert.Equal(t, 3, bst2.size)
	assert.True(t, bst2.Contains(10))
}