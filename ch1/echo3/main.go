// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"

	"strings"
	// "strconv"
	"time"
)

//!+
// func main() {
// 	fmt.Println(strings.Join(os.Args[1:], " "))
// }

// func main() {
// 	fmt.Println(strings.Join(os.Args[1:], " "))
// }

// Exercise 1.1 modify the echo program to also print
// os.Args[0], the name of the command that invoked it.
// func main() {
// 	fmt.Println(strings.Join(os.Args, " "))
// }

// Exercise 1.2 modify the echo program to print
// the index and value of each of its arguments, one per line.
// func main() {
// 	for i, arg := range os.Args[1:] {
// 		fmt.Println(strconv.Itoa(i) + " " + arg)
// 	}
// }

// Exercise 1.3 Experiment to measure the difference in
// runing time between our potentially inefficient versions
// and the one that uses strings.Join

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	start := time.Now()
	echo1()
	x := time.Since(start).Seconds()
	fmt.Printf("%.2fs elapsed\n", x)
	echo2()
	fmt.Println("FOO FAN FOOM")
	y := time.Since(start).Seconds() - x
	fmt.Printf("%.2fs elapsed\n", y)
	fmt.Println("FOO FAN FOOM")
	z := time.Since(start).Seconds() - x - y
	fmt.Printf("%.2fs elapsed\n", z)
	fmt.Println("FOO FAN FOOM")
}

//!-
