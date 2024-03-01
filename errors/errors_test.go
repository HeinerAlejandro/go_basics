package main

import (
	"testing"
)

func TestDiv0(t *testing.T) {

	terms1 := &DivTerms{
		first:  22,
		second: 0,
	}

	_, err := MakeDiv(terms1)

	if err == nil {
		t.Fatalf("La division por 0 no esta permitida %v\n", terms1.second)
	}

}
