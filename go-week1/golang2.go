package main

import "fmt"

func main(){
	power := getPower()
	fmt.Printf("It's over %d\n", power)

	power1, power := 9002, 8989

	fmt.Printf("It's value is varying to %d %d\n ", power, power1)
}

func getPower() int {
	return 90001
}
