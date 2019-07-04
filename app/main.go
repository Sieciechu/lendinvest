package main

import (
	"fmt"

	"github.com/Sieciechu/lendinvest"
	"github.com/Sieciechu/lendinvest/calendar"
)

func main() {
	fmt.Println("test")
	x := calendar.NewDate("2022-09-30")
	fmt.Println(x)
	z := lendinvest.Lendinvest{}
	fmt.Println(z)
}
