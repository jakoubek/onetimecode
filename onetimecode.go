package onetimecode

import (
	"github.com/google/uuid"
	"math"
	"math/rand"
	"strings"
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
	ulmcase int
	withoutDashes bool
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

func WithMin(min int) OnetimecodeConfig {
	return func(code *Onetimecode) {
		code.min = min
	}
}

func WithMax(max int) OnetimecodeConfig {
	return func(code *Onetimecode) {
		code.max = max
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

func WithUpperCase() OnetimecodeConfig {
	return func(code *Onetimecode) {
		code.ulmcase = 1
	}
}

func WithLowerCase() OnetimecodeConfig {
	return func(code *Onetimecode) {
		code.ulmcase = -1
	}
}

func WithoutDashes() OnetimecodeConfig {
	return func(code *Onetimecode) {
		code.withoutDashes = true
	}
}

func NewNumericalCode(opts ...OnetimecodeConfig) *Onetimecode {
	otc := &Onetimecode{
		codeType: ANumberedCode,
		length: 6,
		min: 1,
		max: 999999,
	}
	for _, opt := range opts {
		opt(otc)
	}
	otc.defineValueNumeric()
	return otc
}

func NewAlphanumericalCode(opts ...OnetimecodeConfig) *Onetimecode {
	otc := &Onetimecode{
		codeType: AnAlphaNumericCode,
		length: 6,
		ulmcase: 0,
	}
	for _, opt := range opts {
		opt(otc)
	}
	otc.defineValueAlphanumeric()
	return otc
}

func NewUuidCode(opts ...OnetimecodeConfig) *Onetimecode {
	otc := &Onetimecode{
		withoutDashes: false,
	}
	for _, opt := range opts {
		opt(otc)
	}
	otc.stringCode = Uuid()
	if otc.withoutDashes == true {
		otc.stringCode = strings.Replace(otc.stringCode, "-", "", -1)
	}
	return otc
}

func (otc *Onetimecode) Code() string {
	if otc.stringCode == "" {
		otc.stringCode = string(otc.code)
	}
	return otc.stringCode
}

func (otc *Onetimecode) NumberCode() int64 {
	return otc.code
}

func (otc *Onetimecode) defineValueNumeric() {
	rand.Seed(time.Now().UnixNano())
	rndNr := rand.Intn(otc.max-otc.min+1) + otc.min
	otc.code = int64(rndNr)
}

func (otc *Onetimecode) defineValueAlphanumeric() {
	otc.stringCode = alphaNumberCode(otc.length, otc.ulmcase)
}

// AlphaNumberCode returns an alphanumeric randomized
// code of the given length with numbers, uppercase
// and lowercase characters.
func alphaNumberCode(length int, ulmcase int) string {
	const charsetMixed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const charsetUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const charsetLower = "abcdefghijklmnopqrstuvwxyz0123456789"
	var charset string
	switch ulmcase {
	case -1:
		charset = charsetLower
	case 1:
		charset = charsetUpper
	default:
		charset = charsetMixed
	}

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