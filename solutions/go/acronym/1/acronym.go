// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	// Заменим все дефисы и тире
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", " ", -1)	
	// Разбиваем строку на слова
	words := strings.Split(s, " ")
	// Создаем переменную для хранения акронимы
	var acronym string
	// Проходимся по каждому слову
	for _, word := range words {
		if len(word) > 0 {
		
			// Добавляем первую букву слова в акроним
			// Переводим букву в верхний регистр
			// Проверим. что это буквы
			char := word[0]
			if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			acronym += strings.ToUpper(string(char))
			}
		}	
	}
	return acronym
}
