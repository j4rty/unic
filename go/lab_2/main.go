package main

import (
	"fmt"
)

type device interface {
	PrintInfo()
	Resail() int
	PrintResail()
	PrintShortInfo()
	Discount()
}

// базовый тип
type Computer struct {
	model string
	year  int
	brand string
	price int
}

type Laptop struct {
	Computer
	Width    float64
	WorkTime int
	Monitor  float64
}

type View struct {
	Laptop
	Oc  string
	Ram int
	Gpu string
}

// геттеры и сетторы структуры Computer
func (c *Computer) Model() string {
	return c.model
}

func (c *Computer) Year() int {
	return c.year
}

func (c *Computer) Brand() string {
	return c.brand
}

func (c *Computer) Price() int {
	return c.price
}

func (c *Computer) SetPrice(price int) {
	c.price = price
}

func (c *Computer) SetModel(model string) {
	c.model = model
}

func (c *Computer) SetYear(year int) {
	c.year = year
}
func (c *Computer) SetBrand(brand string) {
	c.brand = brand
}

func NewComputer(model string, year int, brend string, price int) *Computer {
	com := Computer{model, year, brend, price}
	return &com
}

func NewLaptop(model string, year int, brend string, price int, width float64, worktime int, monitor float64) *Laptop {
	com := Computer{model, year, brend, price}
	lap := Laptop{com, width, worktime, monitor}
	return &lap
}

func NewView(model string, year int, brend string, price int, width float64, worktime int, monitor float64, oc string, ram int, gpu string) *View {
	com := Computer{model, year, brend, price}
	lap := Laptop{com, width, worktime, monitor}
	vie := View{lap, oc, ram, gpu}
	return &vie
}

func (c *Computer) PrintInfo() {
	fmt.Printf("Computer Model: %s\n", c.model)
	fmt.Printf("Year: %d\n", c.year)
	fmt.Printf("Brand: %s\n", c.brand)
	fmt.Printf("Price: %d $\n ", c.price)
	fmt.Printf("\n")

}

func (l *Laptop) PrintInfo() {
	fmt.Printf("Laptop Model: %s\n", l.model)
	fmt.Printf("Laptop Year: %d\n", l.year)
	fmt.Printf("Laptop Brand: %s\n", l.brand)
	fmt.Printf("Price: %d $\n", l.price)
	fmt.Printf("Laptop Work Time: %d hour \n", l.WorkTime)
	fmt.Printf("Laptop Monitor: %.1f inch\n", l.Monitor)
	fmt.Printf("\n")
}

func (v *View) PrintInfo() {
	fmt.Printf("View Model: %s\n", v.model)
	fmt.Printf("View Year: %d\n", v.year)
	fmt.Printf("View Brand: %s\n", v.brand)
	fmt.Printf("Price: %d $\n", v.price)
	fmt.Printf("View Work Time: %d hour\n", v.WorkTime)
	fmt.Printf("View Monitor: %.1f inch\n", v.Monitor)
	fmt.Printf("View Gpu: %s \n", v.Gpu)
	fmt.Printf("View Ram: %d Gb\n", v.Ram)
	fmt.Printf("View Oc ID: %s\n", v.Oc)
	fmt.Printf("\n")
}

func (c *Computer) Resail() int {
	res := int(float64(c.price) * (1.0 - (0.1 * float64(2026-c.year))))
	return res
}

func (l *Laptop) Resail() int {
	res := int(float64(l.price) * (1.0 - (0.1 * float64(2026-l.year))))
	return res
}

func (v *View) Resail() int {
	res := int(float64(v.price) * (1.0 - (0.1 * float64(2026-v.year))))
	return res
}

func (c *Computer) PrintResail() {
	fmt.Printf("Rasail price: %d $\n", c.Resail())
}

func (l *Laptop) PrintResail() {
	fmt.Printf("Rasail price: %d $\n", l.Resail())
}

func (v *View) PrintResail() {
	fmt.Printf("Rasail price: %d $\n", v.Resail())
}

func (c *Computer) Discount() {
	dis := int(c.price - 101)
	fmt.Printf("Discount price: %d $\n", dis)
	fmt.Printf("\n")
}

func (l *Laptop) Discount() {
	dis := int(l.price - 101)
	fmt.Printf("Discount price: %d $\n", dis)
	fmt.Printf("\n")
}

func (v *View) Discount() {
	dis := int(v.price - 101)
	fmt.Printf("Discount price: %d $\n", dis)
	fmt.Printf("\n")
}

func (c *Computer) PrintShortInfo() {
	fmt.Printf("Computer Model: %s\n", c.model)
	fmt.Printf("Year: %d\n", c.year)
	fmt.Printf("Brand: %s\n", c.brand)
	fmt.Printf("\n")
}

func (l *Laptop) PrintShortInfo() {
	fmt.Printf("Computer Model: %s\n", l.model)
	fmt.Printf("Year: %d\n", l.year)
	fmt.Printf("Brand: %s\n", l.brand)
	fmt.Printf("\n")
}

func (v *View) PrintShortInfo() {
	fmt.Printf("Computer Model: %s\n", v.model)
	fmt.Printf("Year: %d\n", v.year)
	fmt.Printf("Brand: %s\n", v.brand)
	fmt.Printf("\n")
}

func main() {
	computer1 := NewComputer("MackBook", 2025, "Apple", 1500)
	computer1.PrintInfo()
	computer1.PrintResail()
	computer1.Discount()

	laptop1 := NewLaptop("Tuf A15", 2023, "ASUS", 1100, 1.9, 3, 15.6)
	laptop1.PrintInfo()
	laptop1.PrintResail()
	laptop1.Discount()

	view1 := NewView("ThinkPad", 2018, "Lenovo", 300, 1.4, 6, 13.0, "Fedora", 16, "Integrated")
	view1.PrintInfo()
	view1.PrintResail()
	view1.Discount()
	view1.PrintShortInfo()

	fmt.Printf("=====Полиморфизм====\n")
	fmt.Printf("\n")
	device := []device{computer1, laptop1, view1}
	for _, d := range device {
		d.PrintShortInfo()
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Println("Модель:", computer1.Model(), "Год выпуска:", computer1.Year(), "Бренд:", computer1.Brand(), "Цена покупки:", computer1.Price())
	fmt.Println("Модель:", laptop1.Model(), "Год выпуска:", laptop1.Year(), "Бренд:", laptop1.Brand(), "Цена покупки:", laptop1.Price(), "Вес ноутбука:", laptop1.Width, "Автономность:", laptop1.WorkTime, "Диагональ:", laptop1.Monitor)
	fmt.Println("Модель:", view1.Model(), "Год выпуска:", view1.Year(), "Бренд:", view1.Brand(), "Цена покупки:", view1.Price(), "Вес ноутбука:", view1.Width, "Автономность:", view1.WorkTime, "Диагональ:", view1.Monitor, "Операционная система:", view1.Oc, "Объем Ram", view1.Ram, "Тип Gpu:", view1.Gpu)
}
