// Дан номер некоторого года (целое положительное число). Опреде-
// лить соответствующий ему номер столетия, учитывая, что, к примеру, на-
// чалом 20 столетия был 1901 год
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a uint
	for a == 0 {
		a = ioutil.UInteger("номер некоторого года (целое положительное число)")
	}
	if a%10 == 0 {
		a--
	}
	a = (a / 100) + 1
	fmt.Printf("номер столетия: %v\n", a)
}
