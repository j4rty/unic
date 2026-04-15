package computer

import "fmt"

// Интерфейс Device
type Device interface {
	PrintInfo()
	Resail() int
	PrintResail()
	PrintShortInfo()
	Discount()
}

// базовый тип Computer
type Computer struct {
	Model string
	Year  int
	Brand string
	Price int
}

// Конструктор
func NewComputer(model string, year int, brand string, price int) *Computer {
	return &Computer{model, year, brand, price}
}

// Методы Computer
func (c *Computer) PrintInfo() {
	fmt.Printf("Computer Model: %s\n", c.Model)
	fmt.Printf("Year: %d\n", c.Year)
	fmt.Printf("Brand: %s\n", c.Brand)
	fmt.Printf("Price: %d $\n", c.Price)
	fmt.Printf("\n")
}

func (c *Computer) Resail() int {
	return int(float64(c.Price) * (1.0 - (0.1 * float64(2026-c.Year))))
}

func (c *Computer) PrintResail() {
	fmt.Printf("Resale price: %d $\n", c.Resail())
}

func (c *Computer) Discount() {
	dis := c.Price - 101
	fmt.Printf("Discount price: %d $\n", dis)
	fmt.Printf("\n")
}

func (c *Computer) PrintShortInfo() {
	fmt.Printf("Computer Model: %s\n", c.Model)
	fmt.Printf("Year: %d\n", c.Year)
	fmt.Printf("Brand: %s\n", c.Brand)
	fmt.Printf("\n")
}
