package onetimecode

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type OnetimecodeConfig func(*NumberedCode)

type NumberedCode struct {
	length int
	min int
	max int
	code int64
}

func WithLength(length int) OnetimecodeConfig {
	return func(code *NumberedCode) {
		code.length = length
		code.min = int(math.Pow(10, float64(length-1)))
		code.max = int(math.Pow(10, float64(length)))-1
	}
}

func WithMinMax(min, max int) OnetimecodeConfig {
	return func(code *NumberedCode) {
		code.min = min
		code.max = max
	}
}

func NewNumberedCode(opts ...OnetimecodeConfig) *NumberedCode {
	otc := &NumberedCode{
		length: 6,
	}
	for _, opt := range opts {
		opt(otc)
	}
	otc.defineValue()
	return otc
}

func (otc *NumberedCode) String() string {
	return fmt.Sprint(otc.code)
}

func (otc *NumberedCode) defineValue() {
	rand.Seed(time.Now().UnixNano())
	rndNr := rand.Intn(otc.max-otc.min) + otc.min
	otc.code = int64(rndNr)
}