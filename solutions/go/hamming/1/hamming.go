    package hamming
    import "errors"
    func Distance(a, b string) (int, error) {
    	if len(a) != len(b) {
            return 0, errors.New("Chain not equal len!")
        } 
       // if len(a) == 0 && len(b) == 0 {
        //    return 0, nil
        //} 
        
        str1 := []byte(a)
        str2 := []byte(b) 
        count := 0
        //lenght := len(str1)
        for i:=0; i < len(str1); i++ {
            if str1[i] != str2[i]{
                count++
            }
        }
        return count, nil
        //panic("Implement the Distance function")
    }
