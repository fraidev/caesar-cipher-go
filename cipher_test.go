package main

import (
	"testing"
)

func TestCipher_Decrypt(t *testing.T) {
	tests := []struct {
		input string
		want string
	}{
		{input:"ABC", want:"DEF"},
		{input:"XYZ", want:"ABC"},
	}
	for _, tt := range tests {
		t.Run(tt.input + " must be " + tt.want, func(t *testing.T) {
			c := &Cipher{}
			if got := c.Decrypt(tt.input); got != tt.want {
				t.Errorf("Cipher.Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCipher_Encrypt(t *testing.T) {
	tests := []struct {
		input string
		want string
	}{
		{input:"DEF", want:"ABC"},
		{input:"ABC", want:"XYZ"},
	}
	for _, tt := range tests {
		t.Run(tt.input + " must be " + tt.want, func(t *testing.T) {
			c := &Cipher{}
			if got := c.Encrypt(tt.input); got != tt.want {
				t.Errorf("Cipher.Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
