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
	var mode string
	var length int
	var min int
	var max int
	var code string

	flag.StringVar(&mode, "mode", "numbers", "mode for code generation")
	flag.IntVar(&length, "length", -1, "length of generated code")
	flag.IntVar(&min, "min", -1, "minimal number")
	flag.IntVar(&max, "max", -1, "maximal number")
	flag.Parse()

	switch mode {
	case "numbers":
		var tmp *onetimecode.Onetimecode
		if length != -1 {
			tmp = onetimecode.NewNumericalCode(onetimecode.WithLength(length))
		} else {
			if min != -1 && max != -1 {
				tmp = onetimecode.NewNumericalCode(onetimecode.WithMinMax(min, max))
			} else {
				tmp = onetimecode.NewNumericalCode()
			}
		}
		code = tmp.Code()
	case "alphanum":
		var tmp *onetimecode.Onetimecode
		tmp = onetimecode.NewAlphanumericalCode(
			onetimecode.WithLength(length),
			)
		code = tmp.Code()
	case "alphanumuc":
		var tmp *onetimecode.Onetimecode
		tmp = onetimecode.NewAlphanumericalCode(
			onetimecode.WithUpperCase(),
			onetimecode.WithLength(length),
		)
		code = tmp.Code()
	case "alphanumlc":
		var tmp *onetimecode.Onetimecode
		tmp = onetimecode.NewAlphanumericalCode(
			onetimecode.WithLowerCase(),
			onetimecode.WithLength(length),
		)
		code = tmp.Code()
	case "uuid":
		code = onetimecode.Uuid()
	}

	fmt.Println("CODE: ",code)
}
