package atbash

import (
	"strings"
	"unicode"
)

// Atbash кодирует или декодирует строку с помощью шифра Атбаш
func Atbash(s string) string {
	var result strings.Builder
	
	// Приводим строку к нижнему регистру и обрабатываем каждый символ
	for _, char := range strings.ToLower(s) {
		switch {
		case unicode.IsLetter(char):
			// Для букв применяем шифр Атбаш: 'a'->'z', 'b'->'y', etc.
			encoded := 'a' + ('z' - char)
			result.WriteRune(encoded)
		case unicode.IsDigit(char):
			// Цифры остаются без изменений
			result.WriteRune(char)
		}
	}
	
	// Получаем строку без пробелов и знаков препинания
	clean := result.String()
	
	// Разбиваем на группы по 5 символов
	return splitIntoGroups(clean, 5)
}

// splitIntoGroups разбивает строку на группы по n символов
func splitIntoGroups(s string, n int) string {
	if n <= 0 {
		return s
	}
	
	var result strings.Builder
	count := 0
	
	for _, char := range s {
		if count == n {
			result.WriteRune(' ')
			count = 0
		}
		result.WriteRune(char)
		count++
	}
	
	return result.String()
}