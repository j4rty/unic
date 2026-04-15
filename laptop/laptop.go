package laptop

import (
	"fmt"
	"lab_4/computer"
)

type Laptop struct {
	computer.Computer
	Width    float64
	WorkTime int
	Monitor  float64
}

func NewLaptop(model string, year int, brand string, price int, width float64, worktime int, monitor float64) *Laptop {
	comp := computer.Computer{Model: model, Year: year, Brand: brand, Price: price}
	return &Laptop{comp, width, worktime, monitor}
}

func (l *Laptop) PrintInfo() {
	fmt.Printf("Laptop Model: %s\n", l.Model)
	fmt.Printf("Laptop Year: %d\n", l.Year)
	fmt.Printf("Laptop Brand: %s\n", l.Brand)
	fmt.Printf("Price: %d $\n", l.Price)
	fmt.Printf("Laptop Work Time: %d hour \n", l.WorkTime)
	fmt.Printf("Laptop Monitor: %.1f inch\n", l.Monitor)
	fmt.Printf("\n")
}

func (l *Laptop) Resail() int {
	return int(float64(l.Price) * (1.0 - (0.1 * float64(2026-l.Year))))
}

func (l *Laptop) PrintResail() {
	fmt.Printf("Resale price: %d $\n", l.Resail())
}

func (l *Laptop) Discount() {
	dis := l.Price - 101
	fmt.Printf("Discount price: %d $\n", dis)
	fmt.Printf("\n")
}

func (l *Laptop) PrintShortInfo() {
	fmt.Printf("Laptop Model: %s\n", l.Model)
	fmt.Printf("Year: %d\n", l.Year)
	fmt.Printf("Brand: %s\n", l.Brand)
	fmt.Printf("\n")
}
