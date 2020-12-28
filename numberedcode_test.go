package onetimecode

import (
	"fmt"
	"testing"
)

func TestNewNumberedCode(t *testing.T) {
	//nc := NewNumberedCode()

}

func TestNewNumberedCodeWithLength(t *testing.T) {
	cases := []struct{
		length int
	}{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
		{7},
		{8},
		{9},
	}
	for _, c := range cases {
		nc := NewNumberedCode(
			WithLength(c.length),
			)
		//fmt.Println(nc.length, nc.min, nc.max, nc.code)
		if len(fmt.Sprint(nc.code)) != c.length {
			t.Errorf("Code %d should be %d long but is %d", nc.code, c.length, len(fmt.Sprint(nc.code)))
		}
	}
}

func TestNewNumberedCodeWithMinMax(t *testing.T) {
	cases := []struct{
		min int
		max int
	}{
		{1, 6},
		{2, 10},
		{32, 33},
		{10, 100},
		{5, 21},
		{100, 1000},
		{45, 55},
		{1000, 1002},
	}
	for _, c := range cases {
		nc := NewNumberedCode(
			WithMinMax(c.min, c.max),
		)
		//fmt.Println(nc.min, nc.max, nc.code)
		if nc.code < int64(nc.min) || nc.code > int64(nc.max) {
			t.Errorf("Code %d should be between %d and %d", nc.code, nc.min, nc.max)
		}
	}
}