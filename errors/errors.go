package main

import (
	"errors"
	"log"
)

type DivTerms struct {
	first  int16
	second int16
}

func main() {
	terms := &DivTerms{
		first:  10,
		second: 32,
	}

	divResult, err := MakeDiv(terms)

	log.SetPrefix("errors")
	log.SetFlags(0)

	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	log.Printf("The division result was: %f\n", divResult)
}

func MakeDiv(terms *DivTerms) (division float32, error error) {

	if terms.second == 0 {
		return division, errors.New("Division by 0 imposible")
	}

	division = float32(terms.first) / float32(terms.second)

	return
}
