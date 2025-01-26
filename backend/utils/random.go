package utils

import "math/rand"


var alphabets string = "abcdefghijklmnopqrstuvwxyz" 

func RandomString(r int) string {
	bits := []rune{} 
	k := len(alphabets)

	for i := 0; i < r; i++ {
		index := rand.Intn(k) 
		bits = append(bits, rune(alphabets[index])) 
	}

	return string(bits) // Convert the slice of runes to a string and return it
}

func RandomEmail() string {
	return RandomString(8) + "@example.com"
}