package main

import (
	"fmt"
	"sync"
)

// Структура Counter с встроенным мьютексом
type Counter struct {
	value int
	mu    sync.Mutex
}

// Метод для инкрементации счетчика
func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Метод для получения текущего значения счетчика
func (c *Counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func task18() {
	// Создаем экземпляр структуры-счетчика
	counter := Counter{}

	// Количество горутин для инкрементации
	numGoroutines := 10

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутины для инкрементации счетчика
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.getValue())
}
