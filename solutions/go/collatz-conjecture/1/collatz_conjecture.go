package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
        return 0, errors.New("input must be a positive integer")
    }
    count:=0
    sum:= n
    for sum !=1  {
        count++
        //if sum == 1 {
        //    break
        //} 
        if sum%2 == 0 {
            sum = sum / 2
        } else {
            sum = sum *3 +1
        }
    }
    return count, nil
    //panic("Please implement the CollatzConjecture function")
}
