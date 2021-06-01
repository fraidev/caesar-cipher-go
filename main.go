package main

import "fmt"

func main() {
	chipher := Cipher{}

	result := chipher.Encrypt("abc")

	fmt.Println(result)
}
