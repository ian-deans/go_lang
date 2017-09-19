package main

import (
	"fmt"
	// "io/ioutil"
	// m "math"
	// "net/http"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	beyondHello()
}

func beyondHello() {
	// Variables are declared before used
	var x int
	x = 3
	y := 4 // short declarations infer the type from the value assigned

	sum, prod := learnMultiple(x, y) // funcs can take and return multiple arguments
	fmt.Println("sum:", sum, "prod:", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	str := "Learn Go!"

	s2 := `A raw string literal
	can include line breaks.`

	var u uint = 7
	var pi float32 = 22. / 7

	// Conversion syntax with a short declaration
	n := byte('\n')

	var a4 [4]int // arrays have size fixed at compile time

	a5 := [...]int{3, 1, 5, 10, 100} // An array initialized with fixed
																	 // size of 5 elements with the values provided

	// Slices have dynamic size.
	s3 := []int{4, 5, 9}
	s4 := make([]int, 4) // Allocates slice of 4 ints, initialzed to all 0

	var d2 [][]float64
	bs := []byte("a slice") // Type conversion syntax

	// Appending to a slice
	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s)

	p, q := learnMemory() // Declares vars to be type pointer to int.
	fmt.Println(*p, *q) // * follows a pointer. This prints two ints.

	// maps are dynamically growable associative array type
	// like dictionaries in python, objects in javascript, hash in others
	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	// unused variables are an error in Go.
	// The underscore is used to "use" a variable but discard the value.

	_, _, _, _, _, _, _, _ = str, s2, pi, n, a5, s4, bs, u

	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "This is how you write to a file. So much easier than JS.")
	file.Close()

	fmt.Println(s, a4, s3, d2, m)
}

// Go is fullt garbage collected. It has pointers but no pointer arithmetic.
// You can make a mistake with a nil pointer, but not by incrementing a pointer.
func learnMemory() (p, q *int) {
	p = new(int) // the built in function 'new' allocates memory
	// the allocated int is initialized to 0, p is no longer nil

	s := make([]int, 20) // Allocates 20 ints as a single block of memory.
	s[3] = 7	// Assigns a value to one
	r := -2		// Declars another local variable

	return &s[3], &r 	// & takes the address of an object.
}

func learnFlowControl() {
	// If statements use brace brackets but do not require parentheses
	if true {
		fmt.Println("yup")
	}

	if false {
		// nope
	} else {
		// still nope
	}

	x := 42.0
	switch x {
	case 0:
	case 1:
	case 42:
		/* Cases do not fall through in Go, there is a fallthrough keyword however
			 https://github.com/golang/go/wiki/Switch#fall-through */
	case 43:
		// unreached due to lack of fallthroughs
	default:
		// default case is not required but optional
	}

	// Like If statements, For loops also do not need parentheses
	// vars declared in the loop are local to their scope.
	for x := 0; x < 3; x++ {
		fmt.Println("iteration", x)
	}

	// x == 42 here, outside of the for loops's scope

	// While For is the only loop in Go, it has several forms
	// for {
	// 	break
	// 	continue
	// }

	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3}
	{
		// Go supports string substitution
		fmt.Println("key=%s, value=%d\n", key, value)
	}

	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Println("Hello, %s \n", name)
	}

	/*
	if y := expensiveComputation(); y > x {
		x = y
	}
	*/

	// Function literals are closures
	xBig := func() bool {
		return x > 10000 // References x declared above the switch statement
	}
	x = 99999
	fmt.Println("xBig:", xBig()) 	// true
	x = 1.3e3											// Makes x == 1300
	fmt.Println("xBig:", xBig()) 	// false


	// What's more is function literals may be defined and called inline,
	// acting as an argument to function, as long as:
	// a) function literal is called immediately (),
	// b) result type matches expected type of argument.

	fmt.Println("Add + double two numbers: ",
	func(a, b int) int {
		return (a + b) * 2
	}(10, 2) // Called with args 10 and 2
)

}