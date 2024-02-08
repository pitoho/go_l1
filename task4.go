package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func task4() {
	// Создаем канал для передачи данных
	dataChannel := make(chan int)

	// Создаем канал для сигнала завершения
	doneChannel := make(chan struct{})

	// Количество воркеров 
	numWorkers := 3

	// Создаем WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Запускаем воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, doneChannel, &wg)
	}

	// Запускаем горуту для постоянной записи данных в канал
	go dataProducer(dataChannel, doneChannel)

	// Ожидаем получения сигнала завершения (Ctrl+C)
	waitForCtrlC()

	// Закрываем канал завершения, чтобы уведомить воркеров о завершении
	close(doneChannel)

	// Ожидаем завершения всех воркеров
	wg.Wait()
}

func dataProducer(ch chan<- int, done chan struct{}) {
	// Постоянно записываем данные в канал
	for i := 1; ; i++ {
		select {
		case ch <- i:
		case <-done:
			fmt.Println("Производитель завершает работу.")
			return
		}
		time.Sleep(time.Second)
	}
}

func worker(id int, ch <-chan int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Читаем данные из канала и выводим их в stdout
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				fmt.Printf("Worker %d завершает работу.\n", id)
				return
			}
			fmt.Printf("Worker %d получил данные: %d\n", id, data)
		case <-done:
			fmt.Printf("Worker %d получил сигнал завершения. Завершение работы.\n", id)
			return
		}
	}
}

func waitForCtrlC() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT)
	<-signalChannel
	fmt.Println("Программа завершает работу по сигналу Ctrl+C.")
}
