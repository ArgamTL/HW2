package main

import "fmt"

type searchMethod interface {
	search(f *find)
}

//binarysearch

type binarysearch struct {
}

func (l *binarysearch) search(f *find) {
	fmt.Println("Searching by binarysearch strategy")
}

// linearsearch
type linearsearch struct {
}

func (l *linearsearch) search(f *find) {
	fmt.Println("Searching by linearsearch strategy")
}

// find
type find struct {
	heystack     []int
	searchMethod searchMethod
	needle       int
	found        bool
}

func initFind(s searchMethod) *find {
	heystack := make([]int, 10)
	return &find{
		heystack:     heystack,
		searchMethod: s,
		needle:       1,
		found:        false,
	}
}

func (f *find) setsearchMethod(s searchMethod) {
	f.searchMethod = s
}

func (f *find) searchstart(datalist []int, needle int) {
	f.searchMethod.search(f)

	f.needle = needle
	copy(f.heystack, datalist)
}

func main() {

	linearsearch := &linearsearch{}
	find := initFind(linearsearch)
	find.searchstart([]int{80, 71, 46, 58, 45, 86, 99, 251, 400}, 71)

	binarysearch := &binarysearch{}
	find.setsearchMethod(binarysearch)
	find.searchstart([]int{80, 71, 46, 58, 45, 86, 99, 251, 400}, 71)
}
