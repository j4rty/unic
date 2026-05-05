package main

import (
	"fmt"
	"math"
	"time"
)

func calculat(x, y float64) float64 {
	return (math.Pow(x, 3) + math.Sin(math.Pow(x, 3)+math.Pow(y, 3)) + math.Pow(math.Log(math.Cos(y)/math.Sin(y)), 3)) / ((math.Pow(x, 3) + 1) * math.Pow(y, 3) * math.Sqrt(y)) * 1.5 * math.Pow(10, -5)
}

func main() {
	var x, y float64 = 0.5, 1.0

	for i := 0; i < 10; i++ {
		start := time.Now()
		result := calculat(x, y)
		duration := time.Since(start)
		fmt.Println("Итерация", i, ":", result)
		fmt.Println("В наносекундах:", duration.Nanoseconds())
	}
}
