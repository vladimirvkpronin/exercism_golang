package wordcount
import (
    "strings"
    "regexp"
    )

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	// Создаем регулярное выражение для извлечения слов
	// Слово может содержать буквы, цифры и апострофы внутри слова
	re := regexp.MustCompile(`[a-zA-Z0-9]+('[a-zA-Z0-9]+)?`)
	
	// Приводим фразу к нижнему регистру
	lowerPhrase := strings.ToLower(phrase)
	
	// Ищем все совпадения с регулярным выражением
	matches := re.FindAllString(lowerPhrase, -1)
	
	// Создаем карту для подсчета
	out := make(Frequency)
	
	// Подсчитываем каждое слово
	for _, word := range matches {
		out[word]++
	}
	
	return out
    //panic("Please implement the WordCount function")
}
