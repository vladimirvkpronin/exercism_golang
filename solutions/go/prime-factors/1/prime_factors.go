package prime

func Factors(n int64) []int64 {
	factors := []int64{}
    divisor := int64(2)
    
    for n > 1 {
        for n%divisor == 0 {
            factors = append(factors, divisor)
            n /= divisor
        }
        divisor++
    }
    
    return factors
	panic("Please implement the Factors function")
}
