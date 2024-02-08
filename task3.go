package main

import (
	"fmt"
	"sync"
)

func task3() {
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Канал для передачи результатов
	resultChannel := make(chan int, len(numbers))

	// Запускаем горутины для каждого числа
	for _, num := range numbers {
		wg.Add(1)
		go squareCalculator2(num, &wg, resultChannel)
	}

	// Запускаем горуту для закрытия канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Считываем результаты из канала и суммируем их
	sum := 0
	for result := range resultChannel {
		sum += result
	}

	// Выводим результат
	fmt.Println("Сумма квадратов чисел:", sum)
}

func squareCalculator2(number int, wg *sync.WaitGroup, resultChannel chan<- int) {
	defer wg.Done()

	// Рассчитываем квадрат числа
	result := number * number

	// Отправляем результат в канал
	resultChannel <- result
}
