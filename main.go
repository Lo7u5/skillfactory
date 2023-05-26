package main

import (
	"fmt""log"
)

func main() {
	n:=0
	fmt.Print("Ввудите целое число: ")
	_, err:= fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы вели число: %d\n", n)
}
