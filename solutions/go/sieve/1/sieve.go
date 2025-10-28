package sieve

// Sieve возвращает все простые числа до заданного предела n
func Sieve(limit int) []int {
    // Обработка граничных случаев
    if limit < 2 {
        return []int{}
    }

    // Создаем решето (срез булевых значений)
    // Изначально все числа считаются простыми (true)
    sieve := make([]bool, limit+1)
    for i := 2; i <= limit; i++ {
        sieve[i] = true
    }

    // Основной алгоритм решета Эратосфена
    for i := 2; i*i <= limit; i++ {
        if sieve[i] {
            // Помечаем все кратные i как не простые
            for j := i * i; j <= limit; j += i {
                sieve[j] = false
            }
        }
    }

    // Собираем все простые числа
    var primes []int
    for i := 2; i <= limit; i++ {
        if sieve[i] {
            primes = append(primes, i)
        }
    }

    return primes
}