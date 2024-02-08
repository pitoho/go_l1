package main

import (
	"fmt"
)

func setBit(num int64, pos int, value bool) int64 {
    // Устанавливаем или сбрасываем i-й бит в зависимости от значения value
    if value {
        num |= 1 << pos
    } else {
        num &^= 1 << pos
    }
    return num
}

func task8() {
    // Исходное число
    var number int64 = 34
    fmt.Printf("Исходное число в двоичном формате: %b\n", number)
	var bit int = 3
    // Устанавливаем 3-й бит в 1
    number = setBit(number, bit, true)
    fmt.Printf("Число после установки %v-го бита в 1: %b\n", bit, number)
}
