package main

import "strings"

type Cipher struct{}

func (c *Cipher) Encrypt(input string) string {
	return shift(input, 3, Sum)
}

func (c *Cipher) Decrypt(input string) string {
	return shift(input, 3, Subtract)
}

func shift(input string, shiftCount int, encryptType EncryptType) string {
	var operation func(int) int
	switch encryptType {
	case Sum:
		operation = func(x int) int { return x - shiftCount }
	case Subtract:
		operation = func(x int) int { return x + shiftCount }
	}

	var sb strings.Builder
	for _, char := range input {
		shiftedChar := Alphabet[modulos(operation(Alphabet.GetIndex(char)), 26)]
		sb.WriteRune(shiftedChar)
	}
	return sb.String()
}

func modulos(a int, b int) int { return ((a)%b + b) % b }
