package main

import (
	"fmt"
	"reflect"
)

func task14() {
	var myVar interface{}

	myVar = 42
	printVariableType(myVar)

	myVar = "Hello!"
	printVariableType(myVar)

	myVar = true
	printVariableType(myVar)

	myVar = make(chan int)
	printVariableType(myVar)
}

func printVariableType(v interface{}) {
	// Используем рефлексию для получения типа переменной
	varType := reflect.TypeOf(v)

	// Выводим информацию о типе переменной
	fmt.Printf("Значение: %v, Тип: %v\n", v, varType)
}
