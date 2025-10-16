package parsinglogfiles
import (
    "regexp"
    "strings"
    )

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
    //re, err = regexp.Compile(text) 
    return re.MatchString(text)
    //panic("Please implement the IsValidLine function")
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`(<\W*>)`)
    return re.Split(text, -1) 
    //panic("Please implement the SplitLogLine function")
}

func CountQuotedPasswords(lines []string) int {
    count:=0
	re := regexp.MustCompile(`"[^"]*password[^"]*"`)
	
	for _, line := range lines {
		// Приводим строку к нижнему регистру для поиска без учета регистра
		if re.MatchString(strings.ToLower(line)) {
			count++
		}
	}
	return count
    //panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(end-of-line\d*)`)
    return re.ReplaceAllString(text, "") 
    //panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`(User\s+\w+\d*\s)`)
    re2 := regexp.MustCompile(`\s{2,}`)
    re3 := regexp.MustCompile(`User\s`)
    var out []string
    for _, line := range lines {
        if re.MatchString(line) {
            //str := re2.ReplaceAllString(
              str:=  re2.ReplaceAllString("[USR] "+ re3.ReplaceAllString(re.FindString(line),"") + " "," ") + line //, " " )
            
            out = append(out,str)
        } else {
            out = append(out, line)
        }
    }
    return out
    //panic("Please implement the TagWithUserName function")
}
