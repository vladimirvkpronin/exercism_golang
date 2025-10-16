package raindrops
import 	"strconv"
func Convert(number int) string {
     result := ""
    
    // Проверяем делимость на 3, 5, 7 и формируем строку
    if number%3 == 0 {
        result += "Pling"
    }
    if number%5 == 0 {
        result += "Plang"
    }
    if number%7 == 0 {
        result += "Plong"
    }
    
    // Если не делится ни на одно из чисел, возвращаем число как строку
    if result == "" {
        return strconv.Itoa(number)
    }
    
    return result
    //panic("Please implement the Convert function")
}
