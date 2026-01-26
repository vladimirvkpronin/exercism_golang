package sublist

// Relation type is defined in relations.go file.

// Sublist анализирует отношение между двумя списками целых чисел
func Sublist(a, b []int) Relation {
    /*
    Алгоритм:
    1. Проверяем на равенство (одинаковая длина и элементы)
    2. Проверяем, является ли a подсписком b:
       - Пустой список - подсписок любого списка
       - Если a длиннее b, он не может быть подсписком
       - Ищем непрерывное вхождение a в b
    3. Если a не подсписок b, проверяем обратное
       (является ли b подсписком a, что делает a суперсписком)
    4. Если ничего не найдено, возвращаем "unequal"
    */
    
    // Сначала проверяем тривиальные случаи для оптимизации
    switch {
    case len(a) == 0 && len(b) == 0:
        return RelationEqual
    case len(a) == 0:
        // Пустой список является подсписком любого списка
        return RelationSublist
    case len(b) == 0:
        // Непустой список не может быть подсписком пустого списка,
        // но может быть его суперсписком
        return RelationSuperlist
    }
    
    // Проверяем равенство
    if areEqual(a, b) {
        return RelationEqual
    }
    
    // Проверяем, является ли a подсписком b
    if isContained(a, b) {
        return RelationSublist
    }
    
    // Проверяем, является ли b подсписком a
    if isContained(b, a) {
        return RelationSuperlist
    }
    
    return RelationUnequal
}

// areEqual проверяет полное равенство двух списков
func areEqual(a, b []int) bool {
    // Быстрая проверка по длине
    if len(a) != len(b) {
        return false
    }
    
    // Поэлементное сравнение
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }
    
    return true
}

// isContained проверяет, содержится ли список sub в списке seq
func isContained(sub, seq []int) bool {
    // Если sub длиннее seq, он не может содержаться в нем
    if len(sub) > len(seq) {
        return false
    }
    
    // Если списки одинаковой длины, нужно проверить равенство
    if len(sub) == len(seq) {
        return areEqual(sub, seq)
    }
    
    // Ищем sub в seq как непрерывную подпоследовательность
    // Максимальная стартовая позиция для поиска
    maxStart := len(seq) - len(sub)
    
    for start := 0; start <= maxStart; start++ {
        // Проверяем совпадение начиная с позиции start
        match := true
        for i := 0; i < len(sub); i++ {
            if seq[start+i] != sub[i] {
                match = false
                break
            }
        }
        
        if match {
            return true
        }
    }
    
    return false
}