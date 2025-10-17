package isogram

import "strings"

func IsIsogram(word string) bool {
    // Создаем и инициализируем карту для отслеживания встреченных букв
    seen := make(map[rune]bool)
    
    // Приводим слово к нижнему регистру для нечувствительности к регистру
    lowerWord := strings.ToLower(word)
    
    for _, char := range lowerWord {
        // Пропускаем пробелы и дефисы
        if char == ' ' || char == '-' {
            continue
        }
        
        // Если буква уже встречалась, это не изограмма
        if seen[char] {
            return false
        }
        
        // Отмечаем букву как встреченную
        seen[char] = true
    }
    
    // Если дошли до конца без повторений, это изограмма
    return true
}