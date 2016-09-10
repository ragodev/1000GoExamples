// Дни недели пронумерованы следующим образом: 0 — воскресенье,
// 1 — понедельник, 2 — вторник, ..., 6 — суббота. Дано целое число K , ле-
// жащее в диапазоне 1–365. Определить номер дня недели для K -го дня года,
// если известно, что в этом году 1 января было понедельником.
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x uint
	for x == 0 || x > 365 || x < 1 {
		x = util.UInteger("введите K, лежащее в диапазоне 1–365")
	}
	x = x % 7
	fmt.Printf("номер дня недели для K-го дня года: %v\n", x)
}