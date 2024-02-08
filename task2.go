package main

import (
	"fmt"
	"sync"
)

func task2() {
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутины для каждого числа
	for _, num := range numbers {
		wg.Add(1)
		go squareCalculator1(num, &wg)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}

func squareCalculator1(number int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Рассчитываем квадрат числа
	result := number * number

	// Выводим результат в stdout
	fmt.Printf("%d squared is %d\n", number, result)
}
