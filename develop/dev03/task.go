/*
### Утилита sort

Отсортировать строки (man sort)

#### Основное:

Поддержать ключи
* -k — указание колонки для сортировки
* -n — сортировать по числовому значению
* -r — сортировать в обратном порядке
* -u — не выводить повторяющиеся строки

#### Дополнительное:

Поддержать ключи
* -M — сортировать по названию месяца
* -b — игнорировать хвостовые пробелы
* -c — проверять отсортированы ли данные
* -h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type KeyValue struct {
	key   int
	value string
}

func keyValueSort(inputstruct []KeyValue, asc bool) []int {
	key := make(map[string]int)
	var sortedslc []string
	for _, kvs := range inputstruct {
		key[kvs.value] = kvs.key
		sortedslc = append(sortedslc, kvs.value)
	}
	sort.Strings(sortedslc)
	var ascdscslc []int
	if asc {
		for i := 0; i < len(sortedslc); i++ {
			ascdscslc = append(ascdscslc, key[sortedslc[i]])
		}
	} else {
		for i := len(sortedslc) - 1; i >= 0; i-- {
			ascdscslc = append(ascdscslc, key[sortedslc[i]])
		}
	}

	return ascdscslc
}

func main() {
	var c int
	var n, r, u bool
	var inputfile io.Reader
	myfile := flag.Arg(0)

	flag.IntVar(&c, "k", 0, "     указание колонки для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	if myfile == "" {
		fmt.Printf("Using myfile.txt as myfile.\n")
		myfile = "myfile.txt"
	}
	of, err := os.Open(myfile)
	defer func(of *os.File) {
		err := of.Close()
		if err != nil {
			fmt.Printf("Error closing: %s", err)
		}
	}(of)

	if err != nil {
		fmt.Printf("Error opening: %s", err)
		os.Exit(1)
	}

	inputfile = of

	var lines [][]string
	buf := bufio.NewScanner(inputfile)
	for buf.Scan() {
		line := buf.Text()
		lines = append(lines, strings.Split(line, " "))
	}

	var kvs []KeyValue
	for i, str := range lines {
		kvs = append(kvs, KeyValue{
			key:   i,
			value: str[c],
		})
	}

	newSorted := keyValueSort(kvs, !r)
	for _, i := range newSorted {
		fmt.Println(lines[i])
	}
}
