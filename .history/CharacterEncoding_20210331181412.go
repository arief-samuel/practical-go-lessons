package CharacterEncoding

import "fmt"

func main() {
	n := 2548
	n2 := 0x9F4
	fmt.Printf("0x%X\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%x\n", n2)
	fmt.Printf("%d\n", n2)

	fmt.Printf("Decimal : %d\n", n2)

	// n3 is represented using the octal numeral system
	n3 := 02454
	// alternative : n3 := 0o2454

	// convert in decimal
	fmt.Printf("Decimal : %d\n", n3)

	// Binary
	fmt.Printf("Binary: %b\n", n2)

	// n4 is represented using the decimal numeral system
	n4 := 1324
	// output n4 (decimal) in octal
	fmt.Printf("Octal: %o\n", n4)
	// output n4 (decimal) in octal (with a 0o prefix)
	fmt.Printf("Octal with prefix : %O\n", n4)

	// Let’s see how we can create a slice of
	b := make([]byte, 0)
	b = append(b, 255)
	b = append(b, 10)
	b = append(b, 100)
	b = append(b, 25)
	fmt.Println(b)

	// String
	// raw string literals. They are defined between back quotes. Forbidden characters are back quotes Discarded characters are Carriage returns (\r) interpreted string literals. They are defined between double-quotes. Forbidden characters are  new lines unescaped double quotes
	raw := `spring rain:
	browsing under an umbrella
	at the picture-book store`

	fmt.Println(raw)

	interpreted := "i love spring"
	fmt.Println(interpreted)

	// Runes
	s := " Golang"
	for _, v := range s {
		// v is of type rune
		fmt.Printf("Unicode code point : %U - character '%c' - binary  %b - hex %X - Decimal %d\n ", v, v, v, v, v)
	}

	// rune is an alias for int32 and is equivalent to int32 in all ways. It is
	// used, by convention, to distinguish character values from integer values.
	// type rune = int32
	var aRune rune = 'Z'
	fmt.Printf("Unicode Code point of '%c': %U\n", aRune, aRune)

	// 	b := "hello"
	// for i := 0; i < len(b); i++ {
	//    fmt.Println(b[i])
	// }
	j := "hello"
	for i := 0; i < len(j); i++ {
		fmt.Println(j[i])
	}

	// without initialize
	// var password string
	// fmt.Println(password)
	// var roomNumber, floorNumber int
	// fmt.Println(roomNumber, floorNumber)

	// var roomNumber, floorNumber int = 154, 3
	// fmt.Println(roomNumber, floorNumber)

	var password = "notSecured"
	fmt.Println(password)

	// implicit var
	// short variable declaration cannot be used outside functions !
	roomNumber, floorNumber := 154, 3
	fmt.Println(roomNumber, floorNumber)

	const occupancyLimit = 12

	var occupancyLimit1 uint8
	var occupancyLimit2 uint64
	var occupancyLimit3 float32

	// assign our untyped const to an uint8 variable
	occupancyLimit1 = occupancyLimit
	// assign our untyped const to an int64 variable
	occupancyLimit2 = occupancyLimit
	// assign our untyped const to an float32 variable
	occupancyLimit3 = occupancyLimit

	fmt.Println(occupancyLimit1, occupancyLimit2, occupancyLimit3)

	// 	10.2 ... but a default type when needed
	// To understand the notion of default type, let’s take another example

	var occupancyLimit4 string

	// cannot use occupancyLimit (type untyped int) as type string in assignment
	// occupancyLimit4 = occupancyLimit

	fmt.Println(occupancyLimit4)

	// default type is bool
	const isOpen = true
	// default type is rune (alias for int32)
	const MyRune = 'r'
	// default type is int
	const occupancyLimitV2 = 12
	// default type is float64
	const vatRate = 29.87
	// default type is complex128
	const complexNumber = 1 + 2i
	// default type is string
	const hotelName = "Gopher Hotel"

	// variable-constants-types/constants/untyped/overflow/main.go

	// An untyped constant has no limit
	// maximum value of an int is 9223372036854775807
	// 9223372036854775808 (max + 1 ) overflows int
	const profit = 9223372036854775808
	// the program compiles

	// the program not compiles
	// var profit2 int64 = profit
	// fmt.Println(profit2)

	const hotelNames string = "Gophers Hotel"
	const longitude = 24.806078
	const latitude = -78.243027
	var occupancies int = 12
	// A short variable declaration
	// occupancies := 12

	fmt.Println(hotelNames, longitude, latitude)
	fmt.Println(hotelNames)
	fmt.Println(longitude)
	fmt.Println(latitude)
	fmt.Println(occupancies)

}
