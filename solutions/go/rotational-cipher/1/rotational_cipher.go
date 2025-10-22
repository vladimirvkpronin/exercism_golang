package rotationalcipher

// RotationalCipher применяет ротационный шифр к входной строке с заданным ключом
func RotationalCipher(plain string, shiftKey int) string {
    // Нормализуем ключ сдвига в диапазон 0-25
    shiftKey = shiftKey % 26
    if shiftKey < 0 {
        shiftKey += 26
    }

    result := make([]rune, len(plain))
    
    for i, char := range plain {
        switch {
        case char >= 'a' && char <= 'z':
            // Обработка строчных букв
            result[i] = 'a' + (char-'a'+rune(shiftKey))%26
        case char >= 'A' && char <= 'Z':
            // Обработка заглавных букв
            result[i] = 'A' + (char-'A'+rune(shiftKey))%26
        default:
            // Все остальные символы (пробелы, знаки препинания, цифры) остаются без изменений
            result[i] = char
        }
    }
    
    return string(result)
}