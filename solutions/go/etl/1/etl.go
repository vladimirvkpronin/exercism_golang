package etl
import "strings"
func Transform(in map[int][]string) map[string]int {
	out := make(map [string]int)
    for j, i := range in {
        for _, char := range i {
        	out[strings.ToLower(char)] = j    
        }
    }
    return out
    //panic("Please implement the Transform function")
}
