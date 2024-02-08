package main

import (
	"fmt"
	"sync"
)

func task9() {
	// Создаем каналы
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутину для записи чисел в первый канал
	wg.Add(1)
	go writeToChannel(inputChannel, &wg)

	// Запускаем горутину для умножения чисел и записи во второй канал
	wg.Add(1)
	go multiplyAndWriteToChannel(inputChannel, outputChannel, &wg)

	// Запускаем горутину для вывода результатов в stdout
	wg.Add(1)
	go readAndPrintFromChannel(outputChannel, &wg)

	// Ожидаем завершения всех горутин
	wg.Wait()
}

func writeToChannel(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()

	// Записываем числа в первый канал 
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		ch <- num
	}
}

func multiplyAndWriteToChannel(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer close(output)
	defer wg.Done()

	// Читаем числа из первого канала, умножаем на 2 и записываем во второй канал
	for num := range input {
		output <- num * 2
	}
}

func readAndPrintFromChannel(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Читаем и выводим результаты из второго канала
	for result := range ch {
		fmt.Println(result)
	}
}
