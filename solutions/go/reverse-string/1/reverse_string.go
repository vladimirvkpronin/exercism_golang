package reverse

func Reverse(input string) string {
    // Преобразуем строку в срез рун для корректной работы с Unicode
    runes := []rune(input)
    length := len(runes)
    
    // Создаем срез для результата такой же длины
    result := make([]rune, length)
    
    // Переворачиваем строку
    for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
        result[i] = runes[j]
    }
    
    return string(result)
}