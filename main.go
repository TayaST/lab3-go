package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := NewArray()
	fwd := NewForwardList()
	lst := NewList()
	q := NewQueue()
	st := NewStack()
	bst := NewBST()
	ht := NewHashTable(16)

	fmt.Println("======================================================")
	fmt.Println("=== ЛАБОРАТОРНАЯ РАБОТА 3 (Go) ===")
	fmt.Println("======================================================")
	fmt.Println("  M - Массив, F - Односвязный, L - Двусвязный")
	fmt.Println("  Q - Очередь, S - Стек, T - Дерево, H - Хеш-таблица")
	fmt.Println("  PRINT <Структура> - вывести (напр. PRINT M)")
	fmt.Println("  SAVE <file> / LOAD <file> - сохранить/загрузить")
	fmt.Println("  EXIT - выход")
	fmt.Println("======================================================")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\n> ")
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			fmt.Print("> ")
			continue
		}

		parts := strings.Fields(line)
		cmd := parts[0]
		arg1, arg2 := "", ""
		if len(parts) > 1 { arg1 = parts[1] }
		if len(parts) > 2 { arg2 = parts[2] }

		switch cmd {
		case "MPUSH": arr.PushBack(arg1)
		case "MINS": idx, _ := strconv.Atoi(arg1); arr.Insert(idx, arg2)
		case "MGET": idx, _ := strconv.Atoi(arg1); if v, ok := arr.Get(idx); ok { fmt.Println(v) } else { fmt.Println("Ошибка") }
		case "MDEL": idx, _ := strconv.Atoi(arg1); arr.Erase(idx)
		case "MSET": idx, _ := strconv.Atoi(arg1); arr.Set(idx, arg2)
		case "MLEN": fmt.Println("Длина:", arr.Len())

		case "FPUSH": fwd.PushBack(arg1)
		case "FPUSHH": fwd.PushFront(arg1)
		case "FINS": idx, _ := strconv.Atoi(arg1); fwd.InsertAfter(idx, arg2)
		case "FDEL": idx, _ := strconv.Atoi(arg1); fwd.EraseAfter(idx)
		case "FDELH": fwd.PopFront()
		case "FFIND": fmt.Println("Индекс:", fwd.IndexOf(arg1))

		case "LPUSH": lst.PushBack(arg1)
		case "LPUSHH": lst.PushFront(arg1)
		case "LINS": idx, _ := strconv.Atoi(arg1); lst.InsertAfter(idx, arg2)
		case "LDEL": idx, _ := strconv.Atoi(arg1); lst.EraseAfter(idx)
		case "LDELH": lst.PopFront()
		case "LFIND": fmt.Println("Индекс:", lst.IndexOf(arg1))

		case "QPUSH": q.Push(arg1)
		case "QPOP": if v, ok := q.Pop(); ok { fmt.Println("Удален:", v) } else { fmt.Println("Пусто") }
		case "QGET": if v, ok := q.Front(); ok { fmt.Println(v) } else { fmt.Println("Пусто") }

		case "SPUSH": st.Push(arg1)
		case "SPOP": if v, ok := st.Pop(); ok { fmt.Println("Удален:", v) } else { fmt.Println("Пусто") }
		case "SGET": if v, ok := st.Top(); ok { fmt.Println(v) } else { fmt.Println("Пусто") }

		case "TINSERT": val, _ := strconv.Atoi(arg1); bst.Insert(val)
		case "TGET": val, _ := strconv.Atoi(arg1); if bst.Contains(val) { fmt.Println("YES") } else { fmt.Println("NO") }
		case "TCOMPLETE": if bst.IsComplete() { fmt.Println("YES") } else { fmt.Println("NO") }

		case "HPUT": ht.Insert(arg1, arg2)
		case "HGET": if v, ok := ht.Get(arg1); ok { fmt.Println(v) } else { fmt.Println("Не найдено") }
		case "HDEL": ht.Erase(arg1)

		case "PRINT":
			switch arg1 {
			case "M": fmt.Println(arr.String())
			case "F": fmt.Println(fwd.String())
			case "L": fmt.Println(lst.String())
			case "Q": fmt.Println(q.String())
			case "S": fmt.Println(st.String())
			case "T": fmt.Println(bst.InOrder())
			case "H": fmt.Println(ht.String())
			case "":
				fmt.Println("M:", arr.String())
				fmt.Println("F:", fwd.String())
				fmt.Println("L:", lst.String())
				fmt.Println("Q:", q.String())
				fmt.Println("S:", st.String())
				fmt.Println("T:", bst.InOrder())
				fmt.Println("H:", ht.String())
			}

		case "SAVE":
			file, err := os.Create(arg1)
			if err != nil { fmt.Println("Ошибка файла"); break }
			defer file.Close()
			fmt.Fprintf(file, "M_array : %s\n", arr.String())
			fmt.Fprintf(file, "F_forward_list : %s\n", fwd.String())
			fmt.Fprintf(file, "L_list : %s\n", lst.String())
			fmt.Fprintf(file, "Q_queue : %s\n", q.String())
			fmt.Fprintf(file, "S_stack : %s\n", st.String())
			fmt.Fprintf(file, "H_hash_table : %s\n", ht.String())
			fmt.Fprintf(file, "T_mytree : %s\n", bst.String())
			fmt.Println("Сохранено в", arg1)

		case "LOAD":
			file, err := os.Open(arg1)
			if err != nil { fmt.Println("Файл не найден"); break }
			defer file.Close()
			fs := bufio.NewScanner(file)
			for fs.Scan() {
				text := fs.Text()
				kv := strings.SplitN(text, " : ", 2)
				if len(kv) < 2 { continue }
				key, val := kv[0], kv[1]
				switch key {
				case "M_array": arr.Deserialize(val)
				case "F_forward_list": fwd.Deserialize(val)
				case "L_list": lst.Deserialize(val)
				case "Q_queue": q.Deserialize(val)
				case "S_stack": st.Deserialize(val)
				case "H_hash_table": ht.Deserialize(val)
				case "T_mytree": bst.Deserialize(val)
				}
			}
			fmt.Println("Загружено из", arg1)

		case "EXIT", "QUIT":
			return
		default:
			fmt.Println("Неизвестная команда")
		}
		fmt.Print("> ")
	}
}