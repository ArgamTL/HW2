Что выведет программа? Объяснить вывод программы.
 Объяснить как работают defer’ы и их порядок вызовов.
=====================================================
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

=============================================================================
Ответ: 2 1

В test() имеем дело с named return value|именнованым возвращаемым значением, 
defer срабатывает после x. 
В anotherTest(), defer работает, когда вычисляется x.




