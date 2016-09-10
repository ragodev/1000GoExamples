// Скорость первого автомобиля V 1 км/ч, второго — V 2 км/ч, расстояние
// между ними S км. Определить расстояние между ними через T часов, если
// автомобили удаляются друг от друга. Данное расстояние равно сумме на-
// чального расстояния и общего пути, проделанного автомобилями; общий
// путь = время · суммарная скорость.
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y float64
	x = util.NoNNumber("Введите расстояние между автомобилями")
	y = util.NoNNumber("Введите время езды")
	s := x + (y * 3)
	fmt.Printf("расстояние между ними через %v (ч): %v км\n", y, s)
}