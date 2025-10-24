package prime
import (
	"errors"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be at least 1")
	}

	// Если n = 1, сразу возвращаем первое простое число
	if n == 1 {
		return 2, nil
	}

	// Начинаем с первого простого числа
	primes := []int{2}
	current := 3 // Начинаем проверять с 3

	// Пока не найдем n простых чисел
	for len(primes) < n {
		if isPrime(current, primes) {
			primes = append(primes, current)
		}
		current += 2 // Проверяем только нечетные числа (после 2)
	}

	return primes[n-1], nil
}

// isPrime проверяет, является ли число простым
// primes содержит все найденные простые числа до текущего момента
func isPrime(num int, primes []int) bool {
	// Проверяем делимость на все найденные простые числа
	for _, p := range primes {
		// Если квадрат простого числа больше проверяемого числа,
		// дальше проверять не нужно - число простое
		if p*p > num {
			return true
		}
		// Если число делится на какое-то простое число, оно не простое
		if num%p == 0 {
			return false
		}
	}
	return true
}