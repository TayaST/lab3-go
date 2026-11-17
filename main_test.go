package main

import (
	"os"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMainExecution(t *testing.T) {
	// Сохраняем старый stdin
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	// Подаем на вход кучу разных команд, чтобы покрыть switch-case
	input := `MPUSH A
MINS 1 B
MSET 0 Z
MDEL 1
FPUSH X
FDEL 0
LPUSH Y
QPUSH Q1
QPOP
QGET
SPUSH S1
SPOP
SGET
HPUT key val
HDEL key
TINSERT 10
TINSERT 5
TGET 10
TCOMPLETE
PRINT M
PRINT F
PRINT L
PRINT Q
PRINT S
PRINT T
PRINT H
PRINT
SAVE test_save.txt
LOAD test_save.txt
UNKNOWN_COMMAND
MINS -1 A
MGET -1
MDEL -1
MSET -1 A
SAVE /invalid/path/file.txt
LOAD /invalid/path/file.txt
FPUSH
LPUSH
QPOP
SPOP
HGET nonexistent
TGET 999
PRINT X
EXIT
`
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString(input)
	w.Close()

	main()

	// Проверяем, что файл сохранился
	content, err := os.ReadFile("test_save.txt")
	assert.NoError(t, err)
	assert.True(t, strings.Contains(string(content), "M_array"))

	// Удаляем временный файл
	os.Remove("test_save.txt")
}