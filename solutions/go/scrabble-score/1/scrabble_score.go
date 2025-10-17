package scrabble
import "strings"
func Score(word string) int {
    sum := 0
	lowerWord := strings.ToLower(word)
    for char := range lowerWord {
        switch lowerWord[char] {
        case 'q', 'z':
            sum += 10
        case 'j', 'x':
            sum += 8
        case 'k':
            sum += 5
        case 'f', 'h', 'v', 'w', 'y':
            sum += 4
        case 'b', 'c', 'm', 'p':
            sum += 3
        case 'd', 'g':
            sum += 2
        case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
            sum += 1
        }
    }
    return sum
    //panic("Please implement the Score function")
}
