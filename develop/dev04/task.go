package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println(findAnagram(&[]string{"Пятак", "Пятка", "ТяПка", "листок"}))
}

func findAnagram(strs *[]string) *map[string]*[]string {
	anagrams := make(map[string][]string)
	result := make(map[string]*[]string)
	for _, s := range *strs {
		s = strings.ToLower(s)
		word := strings.Split(s, "")
		sort.Strings(word)
		key := strings.Join(word, "")
		if _, ok := anagrams[key]; !ok {
			anagrams[key] = make([]string, 0)
			anagrams[key] = append(anagrams[key], s)
		} else {
			anagrams[key] = append(anagrams[key], s)
		}
	}
	sortMap(&anagrams)
	for _, k := range anagrams {
		key := k[0]
		result[key] = new([]string)
		for _, v := range k {
			*result[key] = append(*result[key], v)
		}
	}
	return &result
}

func sortMap(anagrams *map[string][]string) {
	for _, v := range *anagrams {
		sort.Strings(v)
	}
}
