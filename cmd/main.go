// Affine Cipher encrypt and decrypt
// Source: https://en.wikipedia.org/wiki/Affine_cipher

package main

import (
	"fmt"
	"log"
)

const dictionary = "abcdefghijklmnopqrstuvwxyz "

func main() {
	// Input data
	str := "serega krasava"
	// Keys
	a, b := 4, 4

	// Initialize populates two maps with dictionary data combined with indexes
	st := NewAffineCipherManager(dictionary)

	fmt.Println("Initial text: " + str)
	fmt.Println("---------------------------------")

	encrypted := st.Encrypt(str, a, b)
	fmt.Printf("Encrypted: %s\n", encrypted)
	fmt.Println("---------------------------------")

	fmt.Printf("Decrypted: %s\n", st.Decrypt(encrypted, a, b))
}

type AffineCipherManager struct {
	dictionary string

	// These two maps are only for optimization purposes 
	runeWithIndex map[rune]int
	indexWithRune map[int]rune
}

func NewAffineCipherManager(dictionary string) *AffineCipherManager {
	st := &AffineCipherManager{
		dictionary:    dictionary,
		runeWithIndex: make(map[rune]int),
		indexWithRune: make(map[int]rune),
	}

	for index, ch := range st.dictionary {
		st.runeWithIndex[ch] = index
		st.indexWithRune[index] = ch
	}

	return st
}

// Encrypt given text with given keys a and b
func (m *AffineCipherManager) Encrypt(text string, a, b int) string {
	if !coprime(a, len(m.dictionary)) {
		log.Fatal("Key A and dictionary length should be Co-Prime numbers!")
	}

	dictionaryLength := len(m.dictionary)
	result := ""

	for _, ch := range text {
		nextIndex := (a*m.runeWithIndex[ch] + b) % dictionaryLength
		// Optimization map indexWithRune helps to remain O(1) when finding letter by index
		nextLetter := m.indexWithRune[nextIndex]
		result += string(nextLetter)
	}

	return result
}

func (m *AffineCipherManager) Decrypt(encrypted string, a, b int) string {
	dictionaryLength := len(m.dictionary)
	inverse := 0
	result := ""
	// Trying to count inverse value (a^-1)
	for ; inverse * a % dictionaryLength != 1; inverse++ {}
	
	for _, ch := range encrypted {
		nextIndex := inverse * (m.runeWithIndex[ch] + dictionaryLength - b) % dictionaryLength
		nextLetter := m.indexWithRune[nextIndex]
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
