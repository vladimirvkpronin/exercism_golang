package grains

import "errors"

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("number must be between 1 and 64")
	}
	
	// Вычисляем 2^(number-1)
	result := uint64(1)
	for i := 1; i < number; i++ {
		result *= 2
	}
	return result, nil
}

func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		// Поскольку мы знаем, что Square не вернет ошибку для i от 1 до 64,
		// мы можем игнорировать ошибку
		val, _ := Square(i)
		total += val
	}
	return total
}