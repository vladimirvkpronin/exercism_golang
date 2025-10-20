package space
//import "time"

type Planet string

func Age(seconds float64, planet Planet) float64 {
    sec := ((((seconds / 60 ) / 60) / 24) / 365.24 )
    switch planet {
        case "Mercury":
        	return sec / 0.2408467
        case "Venus":
        	return sec / 0.61519726
        case "Mars":
        	return sec / 1.8808158
        case "Jupiter":
        	return sec / 11.862615
        case "Saturn":
        	return sec / 29.447498
        case "Uranus":
        	return sec / 84.016846
        case "Neptune":
        	return sec / 164.79132
        case "Earth":
        	return sec
        default:
        	return -1.00
    }
    //panic("Please implement the Age function")
}
