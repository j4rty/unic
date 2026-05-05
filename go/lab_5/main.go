package main

import (
	"fmt"
	"sync"
)

type device interface {
	PrintInfo(wg *sync.WaitGroup)
	Resail() int
	PrintResail()
	PrintShortInfo(mu *sync.Mutex, wg *sync.WaitGroup)
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

func (c *Computer) PrintInfo(wg *sync.WaitGroup) {
	fmt.Printf("[1] Computer Model: %s\n", c.model)
	fmt.Printf("[1] Computer Year: %d\n", c.year)
	fmt.Printf("[1] Computer Brand: %s\n", c.brand)
	fmt.Printf("[1] Computer Price: %d $\n ", c.price)
	fmt.Printf("\n")
	wg.Done()

}

func (l *Laptop) PrintInfo(wg *sync.WaitGroup) {
	fmt.Printf("[2] Laptop Model: %s\n", l.model)
	fmt.Printf("[2] Laptop Year: %d\n", l.year)
	fmt.Printf("[2] Laptop Brand: %s\n", l.brand)
	fmt.Printf("[2] Laptop Price: %d $\n", l.price)
	fmt.Printf("[2] Laptop Work Time: %d hour \n", l.WorkTime)
	fmt.Printf("[2] Laptop Monitor: %.1f inch\n", l.Monitor)
	fmt.Printf("\n")
	wg.Done()
}

func (v *View) PrintInfo(wg *sync.WaitGroup) {
	fmt.Printf("[3] View Model: %s\n", v.model)
	fmt.Printf("[3] View Year: %d\n", v.year)
	fmt.Printf("[3] View Brand: %s\n", v.brand)
	fmt.Printf("[3] View Price: %d $\n", v.price)
	fmt.Printf("[3] View Work Time: %d hour\n", v.WorkTime)
	fmt.Printf("[3] View Monitor: %.1f inch\n", v.Monitor)
	fmt.Printf("[3] View Gpu: %s \n", v.Gpu)
	fmt.Printf("[3] View Ram: %d Gb\n", v.Ram)
	fmt.Printf("[3] View Oc ID: %s\n", v.Oc)
	fmt.Printf("\n")
	wg.Done()
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

func (c *Computer) PrintShortInfo(mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("[4] Computer Model: %s\n", c.model)
	fmt.Printf("[4] Computer Year: %d\n", c.year)
	fmt.Printf("[4] Computer Brand: %s\n", c.brand)
	fmt.Printf("\n")
}

func (l *Laptop) PrintShortInfo(mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("[5] Laptop Computer Model: %s\n", l.model)
	fmt.Printf("[5] Laptop Year: %d\n", l.year)
	fmt.Printf("[5] Laptop Brand: %s\n", l.brand)
	fmt.Printf("\n")
}

func (v *View) PrintShortInfo(mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("[6] View Computer Model: %s\n", v.model)
	fmt.Printf("[6] View Year: %d\n", v.year)
	fmt.Printf("[6] ViewBrand: %s\n", v.brand)
	fmt.Printf("\n")
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	computer1 := NewComputer("MackBook", 2025, "Apple", 1500)
	laptop1 := NewLaptop("Tuf A15", 2023, "ASUS", 1100, 1.9, 3, 15.6)
	view1 := NewView("ThinkPad", 2018, "Lenovo", 300, 1.4, 6, 13.0, "Fedora", 16, "Integrated")

	fmt.Printf("\n")
	fmt.Printf("Горутины без синхронизации завершены\n")
	fmt.Printf("\n")

	wg.Add(3)
	go computer1.PrintInfo(&wg)
	go laptop1.PrintInfo(&wg)
	go view1.PrintInfo(&wg)

	wg.Wait()

	fmt.Printf("Горутины c синхронизации завершены\n")
	fmt.Printf("\n")

	wg.Add(3)
	go computer1.PrintShortInfo(&mu, &wg)
	go laptop1.PrintShortInfo(&mu, &wg)
	go view1.PrintShortInfo(&mu, &wg)
	wg.Wait()

	computer1.PrintResail()
	computer1.Discount()

	laptop1.PrintResail()
	laptop1.Discount()

	view1.PrintResail()
	view1.Discount()

	fmt.Printf("=====Полиморфизм====\n")
	fmt.Printf("\n")
	device := []device{computer1, laptop1, view1}
	for _, d := range device {
		d.PrintInfo(&wg)
	}

}
