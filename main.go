package main

import "fmt"

func main() {
	fmt.Println("Hello IoT Edge device")
	i := 100
	for i > 0 {
		i--
		fmt.Println(string(GetIotData()))
	}
}
