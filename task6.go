package main

import (
	"fmt"
	"sync"
	"time"
)

func task6() {
	// Создаем канал для сигнала завершения
	done := make(chan struct{})

	// Запускаем горутину
	go myGoroutine(done)

	// Ждем 3 секунды
	time.Sleep(3 * time.Second)

	// 1. Завершение по закрытию канала
	close(done)

	// Ждем завершения горутины
	time.Sleep(1 * time.Second)

	// Создаем новый канал
	done2 := make(chan struct{})

	// Запускаем горутину с использованием канала для сигнала завершения
	go myGoroutineWithChannel(done2)

	// Ждем 3 секунды
	time.Sleep(3 * time.Second)

	// 2. Завершение по закрытию канала
	close(done2)

	// Ждем завершения горутины
	time.Sleep(1 * time.Second)

	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Запускаем горутины с использованием WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go myGoroutineWithWaitGroup(&wg)
	}

	// Ждем 3 секунды
	time.Sleep(3 * time.Second)

	// 3. Завершение по WaitGroup
	wg.Wait()

	fmt.Println("Программа завершает работу.")
}

// 1. Завершение по закрытию канала
func myGoroutine(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Горутина завершает работу (закрытие канала).")
			return
		default:
			fmt.Println("Горутина выполняет работу.")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// 2. Завершение по закрытию канала
func myGoroutineWithChannel(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Горутина завершает работу.")
			return
		default:
			fmt.Println("Горутина выполняет работу.")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// 3. Завершение по WaitGroup
func myGoroutineWithWaitGroup(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Горутина выполняет работу.")
		time.Sleep(500 * time.Millisecond)
	}
}
