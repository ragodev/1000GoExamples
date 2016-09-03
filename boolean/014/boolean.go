// Даны три целых числа: A , B , C . Проверить истинность высказыва-
// ния: «Ровно одно из чисел A , B , C положительное»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var a, b, c int
	a = util.Integer("целое число A")
	b = util.Integer("целое число B")
	c = util.Integer("целое число C")
	fmt.Println((a > 0 && b < 0 && c < 0) || (a < 0 && b > 0 && c < 0) || (a < 0 && b < 0 && c > 0))
}
