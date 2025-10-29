package armstrong

import (
	"math"
)

// IsNumber проверяет, является ли число числом Армстронга
func IsNumber(n int) bool {
	if n < 0 {
		return false
	}

	// Находим количество цифр в числе
	original := n
	numDigits := countDigits(n)

	// Вычисляем сумму цифр, возведенных в степень количества цифр
	sum := 0
	temp := n

	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		temp /= 10
	}

	return sum == original
}

// countDigits возвращает количество цифр в числе
func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	
	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	return count
}