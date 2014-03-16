This go Repo will be used to learn basics about the google go programming language

go build "Mein Package"
go install github.com/user/"ORDNERNAME"
run binaries directly

gofmt -w yourcode.go
go fmt path/to/your/package

Go Basic Types

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

That is a slice:

var av = []int{1,5,2,3,7}

And those are arrays:

var av = [...]int{1,5,2,3,7}
var bv = [5]int{1,5,2,3,7}

Slices are pointers with a lengthmethod len(); arrays type consists of a length and the type