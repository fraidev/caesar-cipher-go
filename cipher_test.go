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
		{input:"abc", want:"def"},
		{input:"XYZ", want:"ABC"},
		{input:"xYz", want:"aBc"},
		{input:"x$z", want:"a$c"},
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
		{input:"def", want:"abc"},
		{input:"ABC", want:"XYZ"},
		{input:"aBc", want:"xYz"},
		{input:"a$c", want:"x$z"},
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
