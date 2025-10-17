package luhn

func Valid(id string) bool {
    // Создаем срез для хранения цифр без пробелов
    var clean []byte
    for i := 0; i < len(id); i++ {
        c := id[i]
        if c == ' ' {
            continue // Пропускаем пробелы
        }
        if c < '0' || c > '9' {
            return false // Недопустимый символ
        }
        clean = append(clean, c)
    }

    n := len(clean)
    if n < 2 {
        return false // Строка должна содержать как минимум 2 цифры
    }

    sum := 0
    for i := 0; i < n; i++ {
        digit := int(clean[i] - '0')
        // Удваиваем цифру, если ее позиция с конца четная (начиная с 1 для последней цифры)
        if (n - i) % 2 == 0 {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }
        sum += digit
    }

    return sum % 10 == 0
}