package romannumerals
import (
    "strings"
    "errors"
    )
func ToRomanNumeral(input int) (string, error) {
	// Проверка допустимого диапазона
    if input < 1 || input > 3999 {
        return "", errors.New("Invalid input")
    }

    // Определяем таблицу преобразований
    type romanStruct struct {
        value  int
        symbol string
    }

    conversions := []romanStruct{
        {1000, "M"},
        {900, "CM"},
        {500, "D"},
        {400, "CD"},
        {100, "C"},
        {90, "XC"},
        {50, "L"},
        {40, "XL"},
        {10, "X"},
        {9, "IX"},
        {5, "V"},
        {4, "IV"},
        {1, "I"},
    }

    var result strings.Builder

    // Проходим по всем преобразованиям от большего к меньшему
    for _, conversion := range conversions {
        // Пока текущее число больше или равно значению преобразования
        for input >= conversion.value {
            result.WriteString(conversion.symbol) // Добавляем символ
            input -= conversion.value            // Уменьшаем число
        }
    }

    return result.String(), nil
    //panic("Please implement the ToRomanNumeral function")
}
