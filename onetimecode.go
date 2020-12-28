package onetimecode

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type Onetimecode interface {
	String() string
}

var source = rand.NewSource(time.Now().UnixNano())

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

// Uuid returns an UUID code.
func Uuid() string {
	return uuid.New().String()
}