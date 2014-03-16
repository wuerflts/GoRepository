package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"net"
	"os"
	"sort"
	"time"
)

func main() {
	fmt.Println("The time is", time.Now())
	fmt.Println(os.Open("filename"))
	fmt.Println(net.Dial("tcp", "google.com"))
	rand.Seed(time.Now().UnixNano())
	fmt.Println("a random number: ", rand.Intn(10))

	// big P means public; p means private
	fmt.Println(math.Pi)
	fmt.Println(add(2, 3))
	a, b := swap("world", "hello")
	fmt.Println(a, b)
	fmt.Println(i, c, python, java)

	// in functions var can be used, but := can be used as well with infered type
	k := 3
	fmt.Println(k)

	// for const type is infered as well but := syntax is not allowed
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	var tobecast float64 = 5

	// no implicit casting
	var cast int = int(tobecast)
	fmt.Println(cast)

	// an untyped constant takes the type needed by its context
	fmt.Println(needInt(Small))

	// not possible because go issues an overflow warning !
	// fmt.Println(needInt(Big))

	// if one typed Small as a int it wouldn't compile
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// obviously the () are missing and the type is infered again
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// available with and without semicolons at beginning and end, therefore equivalent to while; thus no while exists in go
	otherSum := 1
	for otherSum < 1000 {
		otherSum += otherSum
	}
	fmt.Println(otherSum)

	// obviously a endless loop
	// for {}

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	q := &v
	fmt.Println(q.X)
	fmt.Println(v, q, r, s, t)

	var aString [2]string
	aString[0] = "Hello"
	aString[1] = "World"
	fmt.Println(aString[0], aString[1])
	fmt.Println(aString)

	// A slice points to an array of values and also includes a length,
	// slices are declared by [] being empty, while an array with infered length has [...]
	arrayForSlice := [...]int{3, 2, 13, 7, 11, 5, 22, 123, 245, 568}

	// [ : ] mimicks a for loop 6 is NOT included
	var slice []int = arrayForSlice[0:6]

	// missing first index means 0 , missing second index means len()
	var sliceFull []int = arrayForSlice[:]
	fmt.Println("slice ==", slice)
	fmt.Println("sliceFull ==", sliceFull)
	fmt.Println("firstElementOfSlice ==", sliceFull[:+1])

	//wouldn't work with an array because sort() expects a slice
	sort.Ints(slice)
	fmt.Println("slice ==", slice)
}

func add(x, y int) int {
	return x + y
}

// multiple return values, multiple declarations of the same type by comma syntax
func swap(x, y string) (string, string) {
	return y, x
}

// default initialization exists, type can be infered from a initialization, var means a list of variables
var i, j = 1, 2
var c, python, java bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	r                 = Vertex{X: 1} // Y:0 is implicit
	s                 = Vertex{}     // X:0 and Y:0
	// new returns a pointer to an zero initialized object of the given type
	t = new(Vertex)
)

// const declaration list
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x float64) string {
	// no surprise for if, no (), but mandatory {}
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// before the condition there can be a statement delimited by a semicolon for if and for
	if v := math.Pow(x, n); v < lim {
		// variables declared before the condition are only in scope until the end of if
		return v
		// end of if v is deleted ( of course in an else block it would still exist )
	}
	// v is no longer existent
	return lim
}

type Vertex struct {
	X, Y int
}
