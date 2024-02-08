package main

import "fmt"

// Определение структуры Human
type Human struct {
    Name   string
    Age    int
    Gender string
}

// Метод PrintInfo для структуры Human
func (h *Human) PrintInfo() {
    fmt.Printf("Name: %s, Age: %d, Gender: %s\n", h.Name, h.Age, h.Gender)
}

// Определение структуры Action с встраиванием Human
type Action struct {
    Human // Встраивание структуры Human
}

// Дополнительный метод для структуры Action
func (a *Action) DoSomething() {
    fmt.Println("The action is being performed!")
}

func print(){
    fmt.Print("ashbfdh")
}

func task1() {
    // Создание объекта типа Action
    person := Action{
        Human: Human{
            Name:   "John",
            Age:    25,
            Gender: "Male",
        },
    }

    // Использование методов из встроенной структуры Human
    person.PrintInfo()

    // Использование метода из структуры Action
    person.DoSomething()
}
