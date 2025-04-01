package main

import "fmt"

func main() {
	fmt.Println("Hello IoT Edge device")
	fmt.Println(string(GetIotData()))
}
