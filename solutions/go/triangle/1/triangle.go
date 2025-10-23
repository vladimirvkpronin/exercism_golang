package triangle

import (
	"math"
)

// Kind представляет тип треугольника
type Kind int

// Константы для типов треугольников
const (
	NaT Kind = iota // Not a Triangle - не треугольник
	Equ             // Equilateral - равносторонний
	Iso             // Isosceles - равнобедренный
	Sca             // Scalene - разносторонний
)

// KindFromSides определяет тип треугольника по длинам его сторон
func KindFromSides(a, b, c float64) Kind {
	// Проверяем, являются ли стороны допустимыми для треугольника
	if !isValidTriangle(a, b, c) {
		return NaT
	}

	// Проверяем равносторонний треугольник (все стороны равны)
	if a == b && b == c {
		return Equ
	}

	// Проверяем равнобедренный треугольник (хотя бы две стороны равны)
	if a == b || b == c || a == c {
		return Iso
	}

	// Если не выполнились предыдущие условия - треугольник разносторонний
	return Sca
}

// isValidTriangle проверяет, могут ли стороны образовать треугольник
func isValidTriangle(a, b, c float64) bool {
	// Проверяем, что все стороны положительные числа
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	// Проверяем неравенство треугольника:
	// Сумма любых двух сторон должна быть больше третьей стороны
	// Используем math.IsInf для обработки особых случаев с бесконечностями
	if math.IsInf(a+b, 0) || math.IsInf(b+c, 0) || math.IsInf(a+c, 0) {
		return false
	}

	return a+b > c && b+c > a && a+c > b
}