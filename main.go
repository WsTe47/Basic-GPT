package main

import "fmt"

func main() {
	g := NewGPTClient()
	res, err := g.SendMsg("哈喽")
	fmt.Println(res, err)
}
