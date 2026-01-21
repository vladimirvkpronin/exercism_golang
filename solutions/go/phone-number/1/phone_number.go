package phonenumber

import (
	"errors"
	"unicode"
)

var (
	ErrInvalidLength      = errors.New("invalid length")
	ErrInvalidCountryCode = errors.New("country code must be 1")
	ErrInvalidAreaCode    = errors.New("area code must start with 2-9")
	ErrInvalidExchange    = errors.New("exchange code must start with 2-9")
)

// getValidatedNumber - общая функция для получения валидного номера
func getValidatedNumber(phoneNumber string) (string, error) {
	// Извлекаем только цифры
	var digits []rune
	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
	}
	
	// Проверяем длину
	if len(digits) < 10 {
		return "", ErrInvalidLength
	}
	
	if len(digits) > 11 {
		return "", ErrInvalidLength
	}
	
	// Если 11 цифр, проверяем код страны
	if len(digits) == 11 {
		if digits[0] != '1' {
			return "", ErrInvalidCountryCode
		}
		digits = digits[1:] // Убираем код страны
	}
	
	// Теперь должно быть ровно 10 цифр
	if len(digits) != 10 {
		return "", ErrInvalidLength
	}
	
	// Проверяем код зоны (первые 3 цифры)
	if digits[0] < '2' || digits[0] > '9' {
		return "", ErrInvalidAreaCode
	}
	
	// Проверяем код обмена (4-я цифра)
	if digits[3] < '2' || digits[3] > '9' {
		return "", ErrInvalidExchange
	}
	
	return string(digits), nil
}



func Number(phoneNumber string) (string, error) {
	return getValidatedNumber(phoneNumber)
	panic("Please implement the Number function")
}

func AreaCode(phoneNumber string) (string, error) {
	validated, err := getValidatedNumber(phoneNumber)
	if err != nil {
		return "", err
	}
	return validated[:3], nil
	panic("Пожалуйста, реализуйте функцию AreaCode")
}

func Format(phoneNumber string) (string, error) {
	validated, err := getValidatedNumber(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + validated[:3] + ") " + validated[3:6] + "-" + validated[6:], nil
	panic("Please implement the Format function")
}
