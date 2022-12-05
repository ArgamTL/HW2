/*
## Поиск анаграмм по словарю

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:

* 'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
* 'листок', 'слиток' и 'столик' - другому.

Входные данные для функции:
* ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.

Выходные данные:
Ссылка на мапу множеств анаграмм.
* Ключ - первое встретившееся в словаре слово из множества
* Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.

Множества из одного элемента __не должны__ попасть в результат.
Все слова должны быть приведены к __нижнему регистру__.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package dev04

import (
	"sort"
	"strings"
)

func RunSort(input *[]string) map[string][]string {
	outputmap := make(map[string][]string)
	inputmap := make(map[string]string)
	for _, str := range *input {
		if len(str) < 2 {
			continue
		}

		str = strings.ToLower(str)
		byteLower := []rune(str)
		sortByte := []rune(str)

		sort.Slice(sortByte, func(i, j int) bool {
			return sortByte[i] <= sortByte[j]
		})
		if v, ok := (inputmap)[string(sortByte)]; ok {
			(outputmap)[v] = append((outputmap)[v], string(byteLower))
		} else {
			(inputmap)[string(sortByte)] = string(byteLower)
			(outputmap)[string(byteLower)] = []string{}
		}
	}
	for k, v := range outputmap {
		if len(v) == 0 {
			delete(outputmap, k)
		}
	}
	return outputmap
}
