package main

type AlphabetLetters []rune

var Alphabet = AlphabetLetters{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
	'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func (a *AlphabetLetters) GetIndex(c rune) int {
	for i := 0; i < len(Alphabet); i++ {
		if Alphabet[i] == c {
			return i
		}
	}

	panic("Index not found")
}
