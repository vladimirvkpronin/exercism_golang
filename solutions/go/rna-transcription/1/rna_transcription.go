package strand

func ToRNA(dna string) string {
    if dna == "" {
        return ""
    }
    
    str := []rune(dna)
    
    for i := 0; i < len(str); i++ {
        switch str[i] {
        case 'G':
            str[i] = 'C'
        case 'C':
            str[i] = 'G'
        case 'T':
            str[i] = 'A'
        case 'A':
            str[i] = 'U'
        }
    }
    
    return string(str)
}
