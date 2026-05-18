package main

import (
	"container/list"
	"fmt"
)

type Table[T any] struct {
	name    string
	columns []string
	rows    *list.List
}

type Database struct {
	tables map[string]interface{} // Хранит таблицы любого типа
}

func NewDatabase() *Database {
	return &Database{
		tables: make(map[string]interface{}),
	}
}

func CreateTable[T any](db *Database, name string, columns []string) *Table[T] {
	table := &Table[T]{
		name:    name,
		columns: columns,
		rows:    list.New(),
	}
	db.tables[name] = table
	return table
}

func GetTable[T any](db *Database, name string) (*Table[T], error) {
	table, exists := db.tables[name]
	if !exists {
		return nil, fmt.Errorf("table '%s' not found", name)
	}

	typedTable, ok := table.(*Table[T])
	if !ok {
		return nil, fmt.Errorf("table '%s' has wrong type", name)
	}
	return typedTable, nil
}

func (t *Table[T]) Insert(row T) {
	t.rows.PushBack(row)
	fmt.Printf("%s: %+v\n", t.name, row)
}

func (t *Table[T]) Find(predicate func(T) bool) []T {
	var result []T
	for e := t.rows.Front(); e != nil; e = e.Next() {
		row := e.Value.(T)
		if predicate(row) {
			result = append(result, row)
		}
	}
	return result
}

func (t *Table[T]) GetAll() []T {
	var result []T
	for e := t.rows.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(T))
	}
	return result
}

type JoinedRow[T1, T2 any] struct {
	Table1 T1
	Table2 T2
}

// Объединение двух таблиц по условию
func JoinTables[T1, T2 any](db *Database, name1, name2 string,
	condition func(T1, T2) bool) ([]JoinedRow[T1, T2], error) {

	// Получаем первую таблицу
	table1, err := GetTable[T1](db, name1)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения таблицы %s: %v", name1, err)
	}

	// Получаем вторую таблицу
	table2, err := GetTable[T2](db, name2)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения таблицы %s: %v", name2, err)
	}

	var result []JoinedRow[T1, T2]

	rows1 := table1.GetAll()
	rows2 := table2.GetAll()

	for _, row1 := range rows1 {
		for _, row2 := range rows2 {
			if condition(row1, row2) {
				result = append(result, JoinedRow[T1, T2]{
					Table1: row1,
					Table2: row2,
				})
			}
		}
	}

	fmt.Printf("Объединение таблиц '%s' и '%s': найдено %d совпадений\n",
		name1, name2, len(result))

	return result, nil
}

type Computer struct {
	Brand string
	Model string
	Year  int
	Price int
}

type Laptop struct {
	Brand  string
	Model  string
	Year   int
	Price  int
	Weight float64
}

func main() {
	db := NewDatabase()

	computers := CreateTable[Computer](db, "computers", []string{"Brand", "Model", "Year", "Price"})
	laptops := CreateTable[Laptop](db, "laptops", []string{"Brand", "Model", "Year", "Price", "Weight"})

	computers.Insert(Computer{Brand: "Apple", Model: "Mac Pro", Year: 2025, Price: 2499})
	computers.Insert(Computer{Brand: "Apple", Model: "Mac Mini", Year: 2024, Price: 799}) // ДОБАВИЛИ
	computers.Insert(Computer{Brand: "Dell", Model: "XPS", Year: 2024, Price: 1599})
	computers.Insert(Computer{Brand: "Asus", Model: "ROG Desktop", Year: 2024, Price: 1499}) // ДОБАВИЛИ
	laptops.Insert(Laptop{Brand: "Apple", Model: "MacBook Pro", Year: 2024, Price: 1999, Weight: 1.6})
	laptops.Insert(Laptop{Brand: "Apple", Model: "MacBook Air", Year: 2024, Price: 1199, Weight: 1.2}) // ДОБАВИЛИ
	laptops.Insert(Laptop{Brand: "Asus", Model: "ROG", Year: 2024, Price: 1799, Weight: 2.1})
	laptops.Insert(Laptop{Brand: "Dell", Model: "XPS Laptop", Year: 2024, Price: 1699, Weight: 1.8}) // ДОБАВИЛИ
	fmt.Println("\n")

	// Параметризованный поиск
	appleComputers := computers.Find(func(c Computer) bool {
		return c.Brand == "Apple"
	})
	fmt.Printf("Поиск компьютеров Apple: %+v\n", appleComputers)
	fmt.Println("\n")

	joinedDevices, err := JoinTables(db, "computers", "laptops", func(c Computer, l Laptop) bool {
		return c.Brand == l.Brand
	})

	if err == nil {
		fmt.Printf("Найдено объединений: %d\n", len(joinedDevices))
		fmt.Println("-------------------------------------")
		for _, j := range joinedDevices {
			fmt.Printf("Бренд: %s\n", j.Table1.Brand)
			fmt.Printf("Компьютер: %s %s ($%d, %d год)\n",
				j.Table1.Brand, j.Table1.Model, j.Table1.Price, j.Table1.Year)
			fmt.Printf("Ноутбук: %s %s ($%d, %.1fкг, %d год)\n",
				j.Table2.Brand, j.Table2.Model, j.Table2.Price, j.Table2.Weight, j.Table2.Year)
			fmt.Println("-------------------------------------")
		}
	} else {
		fmt.Printf("Ошибка объединения: %v\n", err)
	}
}
