package main

import (
	"flag"
	"fmt"
	"strings"
)

const (
	ALPHABET_LENGTH int = 26
	LETTERS string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	input		= flag.String("input", "", "plaintext / ciphertext")
	key			= flag.String("key", "", "encryption key string")
	decrypt = flag.Bool("decrypt", false, "decrypt input with key instead of encrypt")
)

func vigenereCipher(input string, key string, decrypt bool) string {
	chars := []rune(input)
	k := []rune(key)
	keyPos := 0
	output := ""
	for _, m := range chars {
		index := int(m - 'A')
		offset := int(k[keyPos] - 'A') % ALPHABET_LENGTH
		if decrypt {
			index -= offset
		} else {
			index += offset
		}
		if index >= ALPHABET_LENGTH {
			index -= ALPHABET_LENGTH
		} else if index < 0 {
			index += ALPHABET_LENGTH
		}
		output += string(LETTERS[index])

		keyPos++
		if keyPos == len(key) {
			keyPos = 0
		}
	}
	return output
}

func main() {
	flag.Parse()
	input := strings.ToUpper(strings.Replace(*input, " ", "", -1))
	key := strings.ToUpper(strings.Replace(*key, " ", "", -1))
	fmt.Printf("%v\n", vigenereCipher(input, key, *decrypt))
}
