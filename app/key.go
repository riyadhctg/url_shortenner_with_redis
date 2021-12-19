package main

import (
	"math/rand"
)

/**
array vs slice:

var a [10]int
declares a variable a as an array of ten integers.
An array's length is part of its type, so arrays cannot be resized.


A slice, on the other hand, is a dynamically-sized. slices are much more common than arrays.
The type []T is a slice with elements of type T.

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]

ref: go.dev/tour
*/

var charSlice = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func genKey() string {
	s := make([]byte, 10)
	for i := range s {
		s[i] = charSlice[rand.Intn(len(charSlice))]
	}
	return string(s)
}
