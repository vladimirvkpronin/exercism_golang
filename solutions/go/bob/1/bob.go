package bob

import (
	"strings"
	"unicode"
)

// Hey возвращает ответ Боба на заданную ремарку
func Hey(remark string) string {
	// Убираем пробелы в начале и конце строки
	remark = strings.TrimSpace(remark)

	// Проверяем на пустую строку или пробелы
	if remark == "" {
		return "Fine. Be that way!"
	}

	// Проверяем, является ли строка вопросом (заканчивается на ?)
	isQuestion := strings.HasSuffix(remark, "?")

	// Проверяем, является ли строка криком (все буквы в верхнем регистре и есть хотя бы одна буква)
	isShouting := isShouting(remark)

	// Обрабатываем комбинации условий
	switch {
	case isShouting && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isShouting:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}

// isShouting проверяет, является ли строка криком
// Криком считается строка, в которой есть хотя бы одна буква и все буквы в верхнем регистре
func isShouting(s string) bool {
	hasLetters := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetters = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return hasLetters
}