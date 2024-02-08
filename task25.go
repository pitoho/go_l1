package main

import (
	"fmt"
	"time"
)

// Функция sleep принимает время в миллисекундах и "засыпает" программу на это время
func sleep(milliseconds int) {
	// Создаем канал
	ch := make(chan bool)

	// Запускаем горутину, которая закроет канал после указанного времени
	go func() {
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)
		close(ch)
	}()

	// Ждем, пока канал не будет закрыт
	<-ch
}

func task25() {
	fmt.Println("Начало выполнения программы")

	// Используем свою функцию sleep
	sleep(2000)

	fmt.Println("Продолжение выполнения программы после задержки")
}
