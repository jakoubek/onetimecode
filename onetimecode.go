package onetimecode

import (
	"fmt"
	"github.com/google/uuid"
	"math"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

type OnetimecodeType string

const (
	ANumberedCode OnetimecodeType = "ANumberedCode"
	AnAlphaNumericCode = "AnAlphaNumericCode"
	AnAlphaNumericUpperCaseCode = "AnAlphaNumericUpperCaseCode"
)

type OnetimecodeConfig func(code *Onetimecode)

type Onetimecode struct {
	codeType OnetimecodeType
	length int
	min int
	max int
	code int64
	stringCode string
}

func WithLength(length int) OnetimecodeConfig {
	return func(code *Onetimecode) {
		if length != -1 {
			code.length = length
			code.min = int(math.Pow(10, float64(length-1)))
			code.max = int(math.Pow(10, float64(length))) - 1
		}
	}
}

func WithMinMax(min, max int) OnetimecodeConfig {
	return func(code *Onetimecode) {
		code.min = min
		code.max = max
	}
}

func WithAlphaNumericCode() OnetimecodeConfig {
	return func(code *Onetimecode) {
		code.codeType = AnAlphaNumericCode
	}
}

func WithAlphaNumericUpperCaseCode() OnetimecodeConfig {
	return func(code *Onetimecode) {
		code.codeType = AnAlphaNumericUpperCaseCode
	}
}

func NewOnetimecode(opts ...OnetimecodeConfig) *Onetimecode {
	otc := &Onetimecode{
		codeType: ANumberedCode,
		length: 6,
		min: 1,
		max: 999999,
	}
	for _, opt := range opts {
		opt(otc)
	}
	otc.defineValue()
	return otc
}

func (otc *Onetimecode) Code() string {
	switch otc.codeType {
	case ANumberedCode:
		return fmt.Sprint(otc.code)
	case AnAlphaNumericCode:
		return otc.stringCode
	case AnAlphaNumericUpperCaseCode:
		return otc.stringCode
	}
	return ""
}

func (otc *Onetimecode) defineValue() {
	if otc.codeType == ANumberedCode {
		rand.Seed(time.Now().UnixNano())
		rndNr := rand.Intn(otc.max-otc.min) + otc.min
		otc.code = int64(rndNr)
	} else {
		if otc.codeType == AnAlphaNumericCode {
			otc.stringCode = AlphaNumberCode(otc.length)
		} else if otc.codeType == AnAlphaNumericUpperCaseCode {
			otc.stringCode = AlphaNumberUcCode(otc.length)
		}
	}
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