package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["key"] = 7

	value := m["key"]

	fmt.Println(value)

	fmt.Println(m)
}
