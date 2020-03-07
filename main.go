package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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
		code = numberCode(length)
	case "alphanum":
		code = alphaNumberCode(length)
	case "alphanumuc":
		code = alphaNumberUcCode(length)
	}

	fmt.Println(code)
}

func numberCode(length int) string {
	const charset = "0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

func alphaNumberCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

func alphaNumberUcCode(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

func numberCode1(length int) string {
	max, _ := strconv.Atoi("1" + strings.Repeat("0", length-1))
	nr := rand.Intn(max)
	return strconv.Itoa(nr)

	// var ret string
	// for i := 0; i <= length-1; i++ {
	// 	nr := rand.Intn(9)
	// 	fmt.Println(i, nr, ret)
	// 	ret += string(nr)
	// }
	// return ret
}
