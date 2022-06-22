package main

import (
	"log"
	"sync"

	"github.com/antoniomika/syncmap"
)

func main() {
	// With Generics
	testMap := syncmap.New[string, string]()
	testMap.Store("foo", "%s")

	data, _ := testMap.Load("foo")

	log.Printf(data, "Hello world!")

	// Without Generics
	testMap2 := &sync.Map{}
	testMap2.Store("foo", "%s")

	data2, _ := testMap2.Load("foo")
	data2string, ok := data2.(string)
	if !ok {
		log.Fatal("data2 is not a string")
	}

	log.Printf(data2string, "Hello world!")
}
