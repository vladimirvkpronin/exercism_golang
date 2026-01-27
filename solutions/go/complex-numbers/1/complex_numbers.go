package complexnumbers

import (
	"math"
)

// Number представляет комплексное число
// Мы используем неэкспортируемые поля, чтобы обеспечить инкапсуляцию
type Number struct {
	real      float64 // действительная часть
	imaginary float64 // мнимая часть
}

// New создает новое комплексное число
// Эта функция не требуется по заданию, но может быть полезной
func New(real, imaginary float64) Number {
	return Number{real, imaginary}
}

// Real возвращает действительную часть комплексного числа
func (n Number) Real() float64 {
	return n.real
}

// Imaginary возвращает мнимую часть комплексного числа
func (n Number) Imaginary() float64 {
	return n.imaginary
}

// Add складывает два комплексных числа
// z1 + z2 = (a + c) + (b + d)i
func (n1 Number) Add(n2 Number) Number {
	return Number{
		real:      n1.real + n2.real,
		imaginary: n1.imaginary + n2.imaginary,
	}
}

// Subtract вычитает одно комплексное число из другого
// z1 - z2 = (a - c) + (b - d)i
func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		real:      n1.real - n2.real,
		imaginary: n1.imaginary - n2.imaginary,
	}
}

// Multiply умножает два комплексных числа
// z1 * z2 = (a*c - b*d) + (b*c + a*d)i
func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		real:      n1.real*n2.real - n1.imaginary*n2.imaginary,
		imaginary: n1.imaginary*n2.real + n1.real*n2.imaginary,
	}
}

// Times умножает комплексное число на действительный коэффициент
// factor * (a + bi) = (factor*a) + (factor*b)i
func (n Number) Times(factor float64) Number {
	return Number{
		real:      n.real * factor,
		imaginary: n.imaginary * factor,
	}
}

// Divide делит одно комплексное число на другое
// z1 / z2 = (a + bi) / (c + di) = 
// = (a*c + b*d)/(c^2 + d^2) + (b*c - a*d)/(c^2 + d^2)i
func (n1 Number) Divide(n2 Number) Number {
	denominator := n2.real*n2.real + n2.imaginary*n2.imaginary
	
	// Проверка деления на ноль
	if denominator == 0 {
		// В математике деление на ноль не определено
		// В Go мы можем вернуть комплексную бесконечность или NaN
		// Но по условию задачи это не требуется, просто будем осторожны
		return Number{
			real:      math.Inf(1),
			imaginary: math.Inf(1),
		}
	}
	
	return Number{
		real:      (n1.real*n2.real + n1.imaginary*n2.imaginary) / denominator,
		imaginary: (n1.imaginary*n2.real - n1.real*n2.imaginary) / denominator,
	}
}

// Conjugate возвращает комплексно-сопряженное число
// Для z = a + bi сопряженное: a - bi
func (n Number) Conjugate() Number {
	return Number{
		real:      n.real,
		imaginary: -n.imaginary,
	}
}

// Abs возвращает модуль (абсолютное значение) комплексного числа
// |z| = √(a² + b²)
func (n Number) Abs() float64 {
	return math.Sqrt(n.real*n.real + n.imaginary*n.imaginary)
}

// Exp вычисляет e^z, где z - комплексное число, используя формулу Эйлера
// e^(a + bi) = e^a * (cos(b) + i*sin(b))
func (n Number) Exp() Number {
	// Вычисляем e^a
	expA := math.Exp(n.real)
	
	// Вычисляем cos(b) и sin(b)
	cosB := math.Cos(n.imaginary)
	sinB := math.Sin(n.imaginary)
	
	return Number{
		real:      expA * cosB,
		imaginary: expA * sinB,
	}
}