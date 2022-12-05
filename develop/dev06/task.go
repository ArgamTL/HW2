/*
## Утилита cut

__Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные__

Поддержать флаги:
* -f - "fields" - выбрать поля (колонки)
* -d - "delimiter" - использовать другой разделитель
* -s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type MyFlag struct {
	Fields    int
	Delimiter string
	Separated bool
}

func NewMyFlag() *MyFlag {
	return &MyFlag{
		Fields:    0,
		Delimiter: "",
		Separated: false,
	}
}

func main() {

	myflags := NewMyFlag()
	flag.IntVar(&myflags.Fields, "f", 0, "'fields' - выбрать поля (колонки)")
	flag.StringVar(&myflags.Delimiter, "d", "\t", "'delimiter' - использовать другой разделитель")
	flag.BoolVar(&myflags.Separated, "s", false, "'separated' - только строки с разделителем")
	flag.Parse()
	args := flag.Args()

	if myflags.Fields == 0 {
		log.Fatalln("use -f with value > 0")
	}

	if len(args) == 0 {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalln(err)
			}
			res, _ := Cut(text, myflags)
			fmt.Println(res)
		}
	}

	fileName := args[len(args)-1]
	file, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	splitString := strings.Split(string(file), "\n")

	for _, str := range splitString {
		if res, ok := Cut(str, myflags); ok {
			fmt.Println(res)
		}
	}
}

func Cut(str string, c *MyFlag) (string, bool) {

	if c.Separated && !strings.Contains(str, c.Delimiter) {
		return "", false
	}

	splitStr := strings.Split(str, c.Delimiter)

	if c.Fields <= len(splitStr) {
		return splitStr[c.Fields-1], true
	}

	return "", false
}
