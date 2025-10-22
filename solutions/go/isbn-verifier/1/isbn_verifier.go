package isbn

import (
	"strconv"
	"strings"
	"unicode"
)

// IsValidISBN проверяет, является ли строка валидным ISBN-10
func IsValidISBN(isbn string) bool {
	// Удаляем все дефисы из строки
	cleaned := strings.ReplaceAll(isbn, "-", "")

	// Проверяем длину строки
	if len(cleaned) != 10 {
		return false
	}

	sum := 0
	// Обрабатываем первые 9 символов
	for i := 0; i < 9; i++ {
		char := rune(cleaned[i])
		if !unicode.IsDigit(char) {
			return false
		}
		digit, _ := strconv.Atoi(string(char))
		sum += digit * (10 - i)
	}

	// Обрабатываем последний символ (может быть цифрой или 'X')
	lastChar := rune(cleaned[9])
	if unicode.IsDigit(lastChar) {
		digit, _ := strconv.Atoi(string(lastChar))
		sum += digit
	} else if lastChar == 'X' {
		sum += 10
	} else {
		return false // Не цифра и не 'X'
	}

	// Проверяем, делится ли сумма на 11 без остатка
	return sum%11 == 0
}