package main

import "fmt"

func main() {
	chipher := Cipher{}

	result := chipher.Encrypt("ABC")

	fmt.Println(result)
}
