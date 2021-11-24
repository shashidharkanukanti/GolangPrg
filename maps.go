package main

import "fmt"

func main() {
	mymap := make(map[int]int)
	newmap := make(map[int]int)
	mymap[1] = 1
	mymap[2] = 2
	mymap[3] = 3
	fmt.Println(mymap)
	for k, v := range mymap {
		newmap[k] = v
	}
	fmt.Println(newmap)
}
