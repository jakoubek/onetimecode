package onetimecode

import (
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

// NumberCode returns a purely numeric randomized
// code of the given length.
func NumberCode(length int) string {
	const charset = "0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

// AlphaNumberCode returns an alphanumeric randomized
// code of the given length with numbers, uppercase
// and lowercase characters.
func AlphaNumberCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

// AlphaNumberUcCode returns an alphanumeric randomized
// code of the given length with numbers and uppercase
// characters.
func AlphaNumberUcCode(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}
