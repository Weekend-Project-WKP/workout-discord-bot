package main

import (
	"bytes"
	"fmt"
	"strings"
)

type printer interface{
	Print() string
}

type user struct {
	username string
	id int
}

func (u user) Print() string {
	return fmt.Sprintf("%v [%v]\n", u.username, u.id)
}

func (mi menuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(mi.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range mi.prices {
		fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
	}
	return b.String()
}

type menuItem struct {
	name string
	prices map[string]float64
}



func main() {
	// Project init
	//services.Hello()

	// Interface demo
/*
	var p printer = user{username: "user1", id: 1}
	fmt.Println(p.Print())

	p = menuItem{name: "Coffee",
		prices: map[string]float64	{"small": 1.65,
			"medium": 1.80,
			"large": 1.95,
		},
	}
	fmt.Println(p.Print())

	u, ok := p.(user)
	fmt.Println(u, ok)
	mi, ok := p.(menuItem)
	fmt.Println(mi, ok)

	switch v:= p.(type) {
	case user:
		fmt.Println("Found a user!", v)
	case menuItem:
		fmt.Println("Found a menueItem!", v)
	default:
		fmt.Println("I'm not sure what this is....")
	}
*/

	// Generics Demo
/*
	testScores := map[string]float64{
		"June": 87.4,
		"Burg": 34.4,
		"Steve": 56.4,
		"Ally": 105,
	}

	c := clone(testScores)
	fmt.Println(c)
*/
	// Switch case here for different slash options that we pull from services?

	// Custom Type Constraints
/*
	a1 := []int{1,2,3}
	a2 := []float64{1.32, 2.98, 3.65}
	a3 := []string{"cat", "dog", "bird"}

	s1 := add(a1)
	s2 := add(a2)
	s3 := add(a3)

	fmt.Printf("Sum of %v: %v\n", a1, s1)
	fmt.Printf("Sum of %v: %v\n", a2, s2)
	fmt.Printf("Sum of %v: %v\n", a3, s3)
*/

	// Errors vs Panics
	// err := errors.New("this is an error")
	// fmt.Println(err)

	// Wrap error in another error
	// err2 := fmt.Errorf("this error wraps the first one: %w", err)
	// fmt.Println(err2)

	//fmt.Println(divide(10, 5))

	// result, err := divide1(10,0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("result:", result)

	// result, err := divide2(10,0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("result:", result)

	// // Goroutines and WaitGroups and channel
	// var wg sync.WaitGroup

	// // create a channel
	// ch := make(chan int)
	
	// // add a wait group
	// wg.Add(1)

	// // create a go function and push a number to the channel
	// go func () {
	// 	// send a message
	// 	ch <- 42
	// 	fmt.Println("This happens async")
	// }()

	// go func () {
	// 	// recieve the message and print it
	// 	msg := <- ch
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }()
	// wg.Wait()

	// // Channel switch
	// ch1, ch2 := make(chan string), make(chan string)
	
	// go func() {
	// 	ch1 <- "message to channel 1"
	// }()

	// go func() {
	// 	ch2 <- "message to channel 2"
	// }()

	// select {
	// case msg := <-ch1:
	// 	fmt.Println(msg)
	// case msg := <-ch2:
	// 	fmt.Println(msg)
	// default:
	// 	fmt.Println("no messages available")
	// }

	// // Loopping channels
	// ch3 := make(chan int)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		ch3 <- i
	// 	}
	// 	close(ch3)
	// }()

	// for msg := range ch3 {
	// 	fmt.Println(msg)
	// }
}

func Add(l, r int) int {
	return l+ r
}

// func divide(l, r int) int {
// 	return l/r
// }

// func divide1(l, r int) (result int, err error) {
// 	if r == 0 {
// 		return 0, errors.New("invalid divisor: must not be zero")
// 	}
// 	return l/r, nil
// }

// func divide2(l, r int) (result int, err error) {
// 	defer func(){
// 		if msg:= recover(); msg != nil{
// 			result = 0
// 			err = fmt.Errorf("%v", msg)
// 		}
// 	}()

// 	return l/r, nil
// }


// type addable interface {
// 	int | float64 | string
// }

// func add[V addable](s []V) V {
// 	var result V
// 	for _, v := range s{
// 		result += v
// 	}
// 	return result
// }

// func clone[K comparable, V any](m map[K]V) map[K]V{
// 	result := make(map[K]V, len(m))
// 	for k, v := range m {
// 		result[k] = v
// 	}
// 	return result
// }