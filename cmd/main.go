package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/jakoubek/onetimecode"
)

var source = rand.NewSource(time.Now().UnixNano())

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var mode string
	var length int
	var code string

	flag.StringVar(&mode, "mode", "numbers", "mode for code generation")
	flag.IntVar(&length, "length", 6, "length of generated code")
	flag.Parse()

	fmt.Println("Length", length)

	switch mode {
	case "numbers":
		code = onetimecode.NumberCode(length)
	case "alphanum":
		code = onetimecode.AlphaNumberCode(length)
	case "alphanumuc":
		code = onetimecode.AlphaNumberUcCode(length)
	}

	fmt.Println(code)
}

// func numberCode(length int) string {
// 	const charset = "0123456789"
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = charset[source.Int63()%int64(len(charset))]
// 	}
// 	return string(b)
// }

// func alphaNumberCode(length int) string {
// 	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = charset[source.Int63()%int64(len(charset))]
// 	}
// 	return string(b)
// }

// func alphaNumberUcCode(length int) string {
// 	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = charset[source.Int63()%int64(len(charset))]
// 	}
// 	return string(b)
// }
