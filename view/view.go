package view

import (
	"fmt"
	"lab_4/laptop"
)

type View struct {
	laptop.Laptop
	Oc  string
	Ram int
	Gpu string
}

func NewView(model string, year int, brand string, price int, width float64, worktime int, monitor float64, oc string, ram int, gpu string) *View {
	comp := laptop.NewLaptop(model, year, brand, price, width, worktime, monitor)
	return &View{*comp, oc, ram, gpu}
}

func (v *View) PrintInfo() {
	fmt.Printf("View Model: %s\n", v.Model)
	fmt.Printf("View Year: %d\n", v.Year)
	fmt.Printf("View Brand: %s\n", v.Brand)
	fmt.Printf("Price: %d $\n", v.Price)
	fmt.Printf("View Work Time: %d hour\n", v.WorkTime)
	fmt.Printf("View Monitor: %.1f inch\n", v.Monitor)
	fmt.Printf("View Gpu: %s \n", v.Gpu)
	fmt.Printf("View Ram: %d Gb\n", v.Ram)
	fmt.Printf("View Oc ID: %s\n", v.Oc)
	fmt.Printf("\n")
}

func (v *View) Resail() int {
	return int(float64(v.Price) * (1.0 - (0.1 * float64(2026-v.Year))))
}

func (v *View) PrintResail() {
	fmt.Printf("Resale price: %d $\n", v.Resail())
}

func (v *View) Discount() {
	dis := v.Price - 101
	fmt.Printf("Discount price: %d $\n", dis)
	fmt.Printf("\n")
}

func (v *View) PrintShortInfo() {
	fmt.Printf("View Model: %s\n", v.Model)
	fmt.Printf("Year: %d\n", v.Year)
	fmt.Printf("Brand: %s\n", v.Brand)
	fmt.Printf("\n")
}
