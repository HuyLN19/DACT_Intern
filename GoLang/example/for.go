package main

import "fmt"

func main() {
	i := 1
	for i<=3 {
		fmt.Println(i)
		i++
	}
	for j:=7; j<=9; j++{
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}
	for i := 1; i <= 5; i++ {
		if i%2 == 0  {
			continue
		}
		fmt.Println(i)
	}
}