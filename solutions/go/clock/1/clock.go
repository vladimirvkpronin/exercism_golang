package clock
import "fmt"
// Define the Clock type here.
// Clock представляет время в минутах с полуночи
type Clock int

func New(h, m int) Clock {
	// Преобразуем часы и минуты в общее количество минут
    totalMinutes := h*60 + m
    // Нормализуем время, приводя его к диапазону 0-1439 минут
    return normalize(totalMinutes)
    //panic("Please implement the New function")
}

// Add добавляет указанное количество минут к текущему времени
func (c Clock) Add(minutes int) Clock {
    // Добавляем минуты и нормализуем результат
    return normalize(int(c) + minutes)
}

// Subtract вычитает указанное количество минут из текущего времени
func (c Clock) Subtract(minutes int) Clock {
    // Вычитаем минуты и нормализуем результат
    return normalize(int(c) - minutes)
}

// String возвращает строковое представление времени в формате HH:MM
func (c Clock) String() string {
    // Вычисляем часы и минуты из общего количества минут
    hours := int(c) / 60
    minutes := int(c) % 60
    
    // Форматируем строку с ведущими нулями
    return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func normalize(minutes int) Clock {
    // Вычисляем остаток от деления на количество минут в сутках (1440)
    normalized := minutes % 1440
    
    // Если результат отрицательный, добавляем 1440 для корректировки
    if normalized < 0 {
        normalized += 1440
    }
    
    return Clock(normalized)
}
