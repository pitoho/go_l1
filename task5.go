package main

import (
	"fmt"
	"time"
)

func task5() {
	// Создаем канал для передачи данных
	dataChannel := make(chan int)

	// Время выполнения программы в секундах
	N := 5

	// Запускаем горуту для отправки данных в канал
	go sendData(dataChannel)

	// Запускаем горуту для чтения данных из канала
	go receiveData(dataChannel)

	// Ждем N секунд
	time.Sleep(time.Duration(N) * time.Second)

	// Завершаем программу
	fmt.Println("Программа завершает работу.")
}

func sendData(ch chan<- int) {
	// Посылаем значения в канал
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(time.Second) 
	}
}

func receiveData(ch <-chan int) {
	// Читаем значения из канала
	for {
		select {
		case data := <-ch:
			fmt.Printf("Принято значение из канала: %d\n", data)
		}
	}
}
