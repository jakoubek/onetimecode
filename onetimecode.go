package onetimecode

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

// RandomNumber returns a random number between
// min and max.
func RandomNumber(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	rndNr := rand.Intn(max-min) + min
	return strconv.Itoa(rndNr)
}

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

// Uuid returns an UUID code.
func Uuid() string {
	return uuid.New().String()
}