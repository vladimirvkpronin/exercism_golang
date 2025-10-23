package resistorcolorduo
import "strconv"
// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
    if len(colors)== 0 {
        return 0
    } 
    sum:= ""
    for k,_ := range colors{ 
	  if k < 2 { 
          switch colors[k] {
        case "black":
        	sum += "0"
        case "brown":
        	sum +=  "1"
        case "red":
        	sum +=  "2"
        case "orange":
        	sum +=  "3"
        case "yellow":
        	sum +=  "4"
        case "green":
        	sum +=  "5"
        case "blue":
        	sum +=  "6"
        case "violet":
        	sum +=  "7"
        case "grey":
        	sum +=  "8"
        case "white":
        	sum +=  "9"
        default:
        sum +=  "0"
    }
        }
        }
    out, err:= strconv.Atoi(sum)
	if err != nil {
        return 0
    }
    return out
    //panic("Implement the Value function")
}
