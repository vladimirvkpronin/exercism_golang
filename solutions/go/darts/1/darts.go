package darts

func Score(x, y float64) int {
	   squaredDistance := x*x + y*y
    
    switch {
    case squaredDistance <= 1:   // Внутренний круг (радиус 1)
        return 10
    case squaredDistance <= 25:  // Средний круг (радиус 5, 5²=25)
        return 5
    case squaredDistance <= 100: // Внешний круг (радиус 10, 10²=100)
        return 1
    default:                     // Промах
        return 0
    }
    //panic("Please implement the Score function")
}
