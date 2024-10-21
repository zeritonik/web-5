package main

import (
	"fmt"
	"sync"
)

// реализовать calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
func calculator(firstChan, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resultChan := make(chan int, 1)
	select {
	case sqr := <-firstChan:
		resultChan <- sqr * sqr
	case mult := <-secondChan:
		resultChan <- 3 * mult
	case <-stopChan:
	}
	close(resultChan)
	return resultChan
}

func main() {
	// здесь должен быть код для проверки правильности работы функции calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
	ch1, ch2 := make(chan int), make(chan int)
	ch3 := make(chan struct{})

	var res <-chan int
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		res = calculator(ch1, ch2, ch3)
		wg.Done()
	}()

	ch1 <- 5
	ch2 <- 3
	close(ch1)
	close(ch2)
	close(ch3)

	wg.Wait()
	fmt.Println(<-res)
}
