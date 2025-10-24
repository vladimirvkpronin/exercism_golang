package lsproduct

import (
	"errors"
)

// LargestSeriesProduct вычисляет наибольшее произведение последовательных цифр заданной длины
func LargestSeriesProduct(digits string, span int) (int64, error) {
    // Проверка корректности входных данных
    if span < 0 {
        return 0, errors.New("span must be non-negative")
    }
    
    if span > len(digits) {
        return 0, errors.New("span must be smaller than string length")
    }
    
    if span == 0 {
        return 1, nil
    }
    
    // Проверка, что все символы - цифры
    for _, char := range digits {
        if char < '0' || char > '9' {
            return 0, errors.New("digits input must only contain digits")
        }
    }

    maxProduct := int64(0)
    
    // Проходим по всем возможным сериям длины span
    for i := 0; i <= len(digits)-span; i++ {
        product := int64(1)
        
        // Вычисляем произведение для текущей серии
        for j := 0; j < span; j++ {
            digit := int64(digits[i+j] - '0')
            product *= digit
        }
        
        // Обновляем максимальное произведение
        if product > maxProduct {
            maxProduct = product
        }
    }
    
    return maxProduct, nil
}