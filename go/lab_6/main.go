package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func server(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var count = 1

	for link := range ch {
		fmt.Printf("Принятая сервером ссылка %d: %s\n", count, link)

		UrlReplacer := strings.NewReplacer(
			"%", "%25",
			"@", "%40",
			"$", "%24",
			"+", "%2B",
			",", "%2C",
			"#", "%23",
			" ", "%20",
			"[", "%5B",
			"]", "%5D",
			"{", "%7B",
			"}", "%7D",
			"|", "%7C",
			"\\", "%5C",
			"^", "%5E",
			"~", "%7E",
			"`", "%60",
			"\"", "%22",
			"<", "%3C",
			">", "%3E",
		)

		processed := UrlReplacer.Replace(link)

		fmt.Printf("Обработанная сервером ссылка %d: %s\n", count, processed)
		count++
	}
}

func user1(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	var link string

	for i := 1; i <= 5; i++ {
		fmt.Printf("\n")
		fmt.Printf("Введите ссылку %d:", i)
		fmt.Scanln(&link)
		ch <- link
		time.Sleep(10 * time.Millisecond)
	}
	close(ch)
}

func user(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	var link string

	for i := 1; i <= 5; i++ {
		fmt.Printf("\n")
		fmt.Printf("Введите ссылку %d:", i)
		fmt.Scanln(&link)
		ch <- link
		time.Sleep(10 * time.Millisecond)
	}
	close(ch)
}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)
	go server(ch, &wg)
	go user(ch, &wg)
	wg.Wait()
}
