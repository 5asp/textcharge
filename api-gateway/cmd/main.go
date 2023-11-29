package main

import (
	"log"

	"github.com/jaevor/go-nanoid"
)

func main() {
	// The canonic NanoID is nanoid.Standard(21).
	canonicID, err := nanoid.Standard(6)
	if err != nil {
		panic(err)
	}

	id1 := canonicID()
	log.Printf("ID 1: %s", id1) // eLySUP3NTA48paA9mLK3V
}
