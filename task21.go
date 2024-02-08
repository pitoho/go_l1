package main

import "fmt"

// Интерфейс Converter с методом ConvertMetersToKilometers
type Converter interface {
	ConvertMetersToKilometers(meters float64) float64
}

// Структура, которая реализует интерфейс Converter
type MetersToKilometersAdapter struct{}

// Реализация метода ConvertMetersToKilometers
func (a *MetersToKilometersAdapter) ConvertMetersToKilometers(meters float64) float64 {
	return meters / 1000
}

// Функция, которая использует интерфейс Converter
func useConverter(converter Converter, meters float64) {
	kilometers := converter.ConvertMetersToKilometers(meters)
	fmt.Printf("%.2f метров равно %.2f километров\n", meters, kilometers)
}

func ефыл21() {
	// Создаем экземпляр адаптера
	adapter := &MetersToKilometersAdapter{}

	// Используем адаптер для перевода метров в километры
	useConverter(adapter, 5000.0)
}
