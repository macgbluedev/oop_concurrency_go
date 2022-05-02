package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	//Explicit Error handling
	myValue, err := strconv.ParseInt("No", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	//Map
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])

	//Slice
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	//Channels, GoRutines
	//c := make(chan int)
	
	//go doSomething(c)
	//<-c

	//Pointers
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
