package main

import (
	"fmt"
	"sync"
	"time"
)

func task7() {
	// Создаем мьютекс для блокировки доступа к map
	var mutex sync.Mutex

	// Создаем map для хранения данных
	dataMap := make(map[string]int)

	// Количество горутин для конкурентной записи
	numGoroutines := 3

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутины для конкурентной записи
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go writeData(&mutex, dataMap, i, &wg)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим результат
	fmt.Println("Результат работы программы:", dataMap)
}

func writeData(mutex *sync.Mutex, dataMap map[string]int, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Заблокируем мьютекс для обеспечения безопасного доступа к map
	mutex.Lock()

	// Конкурентная запись данных в map
	key := fmt.Sprintf("Key%d", id)
	dataMap[key] = id

	// Разблокируем мьютекс
	mutex.Unlock()

	// Имитация работы горутины
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("Горутина %d записала данные в map.\n", id)
}
