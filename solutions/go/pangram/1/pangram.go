package pangram

import (
	"strings"
	//"unicode"
)

func IsPangram(sentence string) bool {
	// Создаем карту для отслеживания встреченных букв
	seen := make(map[rune]bool)
	
	// Приводим предложение к нижнему регистру и итерируемся по символам
	for _, char := range strings.ToLower(sentence) {
		// Проверяем, является ли символ буквой английского алфавита
		if char >= 'a' && char <= 'z' {
			seen[char] = true
		}
	}
	
	// Проверяем, что все 26 букв встретились
	return len(seen) == 26
}
