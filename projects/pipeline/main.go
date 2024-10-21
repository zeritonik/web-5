package main

import "fmt"

// реализовать removeDuplicates(in, out chan string)
func removeDuplicates(in, out chan string) {
	prev := ""
	for val := range in {
		if val != prev {
			out <- val
			prev = val
		}
	}
	close(out)
}

func main() {
	// здесь должен быть код для проверки правильности работы функции removeDuplicates(in, out chan string)
	in, out := make(chan string), make(chan string)

	go func() {
		values := [...]string{"123", "321", "321", "111"}
		for val := range values {
			in <- values[val]
		}
		close(in)
	}()

	go removeDuplicates(in, out)

	for val := range out {
		fmt.Println(val)
	}
}
