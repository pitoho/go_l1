package main

import "fmt"

func task13() {
    a := 5
    b := 10

    fmt.Printf("До обмена: a=%d, b=%d\n", a, b)

    a = a ^ b
    b = a ^ b
    a = a ^ b

    fmt.Printf("После обмена: a=%d, b=%d\n", a, b)
}
