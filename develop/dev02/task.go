/*
### Задача на распаковку

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
- "a4bc2d5e" => "aaaabccddddde"
- "abcd" => "abcd"
- "45" => "" (некорректная строка)
- "" => ""

Дополнительное задание: поддержка escape - последовательностей
- qwe\4\5 => qwe45 (*)
- qwe\45 => qwe44444 (*)
- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package dev02

import "unicode"

func Unpack(l string) string {

	var result []rune
	letters := []rune(l)

	if len(letters) == 0 {
		return ""
	}

	for len(letters) > 0 {
		if len(letters) == 1 {
			result = append(result, letters[0])
			break
		}
		a, b := letters[0], letters[1]
		if unicode.IsDigit(a) && unicode.IsDigit(b) {
			return "Error"

		}
		if unicode.IsDigit(b) {
			var runes []rune
			for i := 0; i < int(b-'0'); i++ {
				runes = append(runes, a)
			}
			result = append(result, runes...)
			letters = letters[2:]
		} else {
			result = append(result, a)
			letters = letters[1:]
		}
	}
	return string(result)
}
