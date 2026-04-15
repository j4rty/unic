package main

import (
	"fmt"
	"lab_4/computer"
	"lab_4/laptop"
	"lab_4/view"
)

func main() {
	computer1 := computer.NewComputer("MackBook", 2025, "Apple", 1500)
	computer1.PrintInfo()
	computer1.PrintResail()
	computer1.Discount()

	laptop1 := laptop.NewLaptop("Tuf A15", 2023, "ASUS", 1100, 1.9, 3, 15.6)
	laptop1.PrintInfo()
	laptop1.PrintResail()
	laptop1.Discount()

	view1 := view.NewView("ThinkPad", 2018, "Lenovo", 300, 1.4, 6, 13.0, "Fedora", 16, "Integrated")
	view1.PrintInfo()
	view1.PrintResail()
	view1.Discount()
	view1.PrintShortInfo()

	fmt.Printf("=====Полиморфизм====\n")
	fmt.Printf("\n")
	devices := []computer.Device{computer1, laptop1, view1}
	for _, d := range devices {
		d.PrintShortInfo()
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Println("Модель:", computer1.Model, "Год выпуска:", computer1.Year, "Бренд:", computer1.Brand, "Цена покупки:", computer1.Price)
	fmt.Println("Модель:", laptop1.Model, "Год выпуска:", laptop1.Year, "Бренд:", laptop1.Brand, "Цена покупки:", laptop1.Price, "Вес ноутбука:", laptop1.Width, "Автономность:", laptop1.WorkTime, "Диагональ:", laptop1.Monitor)
	fmt.Println("Модель:", view1.Model, "Год выпуска:", view1.Year, "Бренд:", view1.Brand, "Цена покупки:", view1.Price, "Вес ноутбука:", view1.Width, "Автономность:", view1.WorkTime, "Диагональ:", view1.Monitor, "Операционная система:", view1.Oc, "Объем Ram", view1.Ram, "Тип Gpu:", view1.Gpu)
}
