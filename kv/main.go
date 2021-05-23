package main

import (
	"fmt"

	"kv/data"
)

func main() {
	fmt.Println("Here we go")
	data.Put("testy", "pav")
	val, ok := data.Get("testy")
	fmt.Println("Value for testy is ", val, ok)
	data.Delete("testy")
	val, ok = data.Get("testy")
	fmt.Println("Value for testy is ", val, ok)
}
