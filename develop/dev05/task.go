/*
## Утилита grep

__Реализовать утилиту фильтрации (man grep)__

Поддержать флаги:
* -A - "after" печатать +N строк после совпадения
* -B - "before" печатать +N строк до совпадения
* -C - "context" (A+B) печатать ±N строк вокруг совпадения
* -c - "count" (количество строк)
* -i - "ignore-case" (игнорировать регистр)
* -v - "invert" (вместо совпадения, исключать)
* -F - "fixed", точное совпадение со строкой, не паттерн
* -n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type FlagSupported struct {
	After      int
	Before     int
	Context    int
	Count      int
	Match      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
	Word       string
}

func NewFlagSupported() *FlagSupported {
	return &FlagSupported{
		After:      0,
		Before:     0,
		Context:    0,
		Match:      false,
		IgnoreCase: false,
		Invert:     false,
		Fixed:      false,
		LineNum:    false,
		Word:       "",
		Count:      0,
	}
}

func (c *FlagSupported) PhraseToLower() {
	c.Word = strings.ToLower(c.Word)
}

func (c *FlagSupported) AdjustbyContext() {

	if c.Context > c.After {
		c.After = c.Context
	}
	if c.Context > c.Before {
		c.Before = c.Context
	}
}

func (c *FlagSupported) AddMatch() {
	c.Count++
}

type Node struct {
	Key int
	Val string
}

type GrepResult struct {
	Result []Node
}

func NewGrep() *GrepResult {
	return &GrepResult{
		Result: []Node{},
	}
}

func (g *GrepResult) SortAsc() {
	sort.Slice(g.Result, func(i, j int) bool {
		return g.Result[i].Key < g.Result[j].Key
	})
}

func (g *GrepResult) PrintRes(indx bool) {
	g.SortAsc()
	if indx {
		for _, v := range g.Result {
			fmt.Printf("[%d. %s]\n\n", v.Key, v.Val)
		}
	} else {
		for _, v := range g.Result {
			fmt.Printf("[%s]\n\n", v.Val)
		}
	}
}

func main() {
	newflag := NewFlagSupported()
	flag.IntVar(&newflag.After, "A", 0, "'after' печатать +N строк после совпадения")
	flag.IntVar(&newflag.Before, "B", 0, "'before' печатать +N строк до совпадения")
	flag.IntVar(&newflag.Context, "C", 0, "'context' (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&newflag.Match, "c", false, "'count' (количество строк)")
	flag.BoolVar(&newflag.IgnoreCase, "i", false, "'ignore-case' (игнорировать регистр)")
	flag.BoolVar(&newflag.Invert, "v", false, "'invert' (вместо совпадения, исключать)")
	flag.BoolVar(&newflag.Fixed, "F", false, "'fixed', точное совпадение со строкой")
	flag.BoolVar(&newflag.LineNum, "n", false, "'line num', печатать номер строки")
	flag.Parse()
	newflag.AdjustbyContext()
	args := flag.Args()

	if len(args) < 2 {
		log.Fatalln("Usage: -[flags] [pattern or string] [file]")
	}

	wordPhraseslice := args[:len(args)-1]
	newflag.Word = strings.Join(wordPhraseslice, " ")

	file, err := os.ReadFile(args[len(args)-1])
	if err != nil {
		log.Fatalln(err)
	}

	splitString := strings.Split(string(file), "\n")
	result := grep(splitString, newflag)
	printRes(newflag, result)
}

func grep(text []string, c *FlagSupported) []*GrepResult {
	var result []*GrepResult
	var condition bool

	for index, str := range text {

		if c.IgnoreCase {
			str = strings.ToLower(str)
			c.PhraseToLower()
		}

		if c.Fixed {
			condition = c.Word == str
		} else {
			condition = strings.Contains(str, c.Word)
		}

		if c.Invert {
			condition = !condition
		}
		match := NewGrep()

		if condition {
			c.AddMatch()
			var upRange, downRange = 0, len(text) - 1
			if d := index - c.Before; d > upRange {
				upRange = d
			}
			if d := index + c.After; d < downRange {
				downRange = d
			}
			for i := upRange; i <= downRange; i++ {
				match.Result = append(match.Result, Node{
					Key: i + 1,
					Val: text[i],
				})
			}
			result = append(result, match)
		}

	}
	return result
}

func printRes(c *FlagSupported, res []*GrepResult) {
	if c.Match {
		fmt.Printf("FOUND: %d\n", c.Count)
	}
	for _, match := range res {
		match.PrintRes(c.LineNum)
	}
	fmt.Println("END")
}
