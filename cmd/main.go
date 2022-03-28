// Affine Cipher encrypt and decrypt
// Source: https://en.wikipedia.org/wiki/Affine_cipher

package main

import (
	"fmt"
	"log"
)

const dictionary = "abcdefghijklmnopqrstuvwxyz "

// These two maps are only for optimization purposes 
var runeWithIndex map[rune]int
var indexWithRune map[int]rune

func initialize() {
	runeWithIndex = make(map[rune]int)
	indexWithRune = make(map[int]rune)

	for index, ch := range dictionary {
		runeWithIndex[ch] = index
		indexWithRune[index] = ch
	}
}

func main() {
	// Input data
	str := "hello world"
	// Keys
	a, b := 4, 4

	// Initialize populates two maps with dictionary data combined with indexes
	initialize()

	fmt.Println("Initial text: " + str)
	fmt.Println("---------------------------------")

	encrypted := Encrypt(str, a, b)
	fmt.Printf("Encrypted: %s\n", encrypted)
	fmt.Println("---------------------------------")

	fmt.Printf("Decrypted: %s\n", Decrypt(encrypted, a, b))
}

// Encrypt given text with given keys a and b
func Encrypt(text string, a, b int) string {
	if !coprime(a, len(dictionary)) {
		log.Fatal("Key A and dictionary length should be Co-Prime numbers!")
	}

	dictionaryLength := len(dictionary)
	result := ""

	for _, ch := range text {
		nextIndex := (a*runeWithIndex[ch] + b) % dictionaryLength
		// Optimization map indexWithRune helps to remain O(1) when finding letter by index
		nextLetter := indexWithRune[nextIndex]
		result += string(nextLetter)
	}

	return result
}

func Decrypt(encrypted string, a, b int) string {
	dictionaryLength := len(dictionary)
	inverse := 0
	result := ""
	// Trying to count inverse value (a^-1)
	for ; inverse * a % dictionaryLength != 1; inverse++ {}
	
	for _, ch := range encrypted {
		nextIndex := inverse * (runeWithIndex[ch] + dictionaryLength - b) % dictionaryLength
		nextLetter := indexWithRune[nextIndex]
		result += string(nextLetter)
	}

	return result
}

func coprime(a, b int) bool {
	if gcd(a,b) == 1 {
		return true
	}

	return false
}

func gcd(a, b int) int {
    if a == 0 || b == 0 {
		return 0
	}
     
    if a == b{
		return a
	} 
     
    if a > b{
		return gcd(a - b, b)
	}
        
             
    return gcd(a, b - a)
}
