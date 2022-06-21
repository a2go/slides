package main

import "log"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func main() {
	log.Println("Inferred:", MapKeys(map[string]string{"foo": "bar"}))

	log.Println("Explicit:", MapKeys[string, string](map[string]string{"foo": "bar"}))
}
