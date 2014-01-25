package main

import (
	"flag"
	"fmt"
	//"strconv"
	"strings"
)

var (
	input		= flag.String("input", "", "plaintext / ciphertext")
	key			= flag.String("key", "", "encryption key string")
	decrypt = flag.Bool("decrypt", false, "decrypt input with key instead of encrypt")
)

func vigenereCipher(input string, key string, decrypt bool) (string, error) {
	return "", nil
}

func main() {
	flag.Parse()
	input := strings.ToUpper(strings.Replace(*input, " ", "", -1))
	key := strings.ToUpper(strings.Replace(*key, " ", "", -1))
	fmt.Printf("%v %v\n", input, key)
	fmt.Printf("%d\n", 's' - 'a')
}
