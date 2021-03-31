package MyHotelApp

import (
	"fmt"
	"math/rand" //*\label{cond1diff}
	"time"
)

func main() {
	hotelName := "The Gopher Hotel"
	fmt.Println("Hotel " + hotelName)
	var roomsAvailable int
	rand.Seed(time.Now().UTC().UnixNano())
	var rooms, roomsOccupied int = 100, rand.Intn(100)
	fmt.Println("rooms :", rooms, " roomsOccupied :", roomsOccupied)

	// Example 1: is there more rooms than roomsOccupied
	fmt.Println("boolean expression : 'rooms > roomsOccupied' is equal to :")
	fmt.Println(rooms > roomsOccupied) //*\label{condExpBool1}

	// Example 2: is there more roomsOccupied than rooms
	fmt.Println("boolean expression : 'roomsOccupied > rooms' is equal to :")
	fmt.Println(roomsOccupied > rooms) //*\label{condExpBool2}

	// Example 3: is roomsOccupied equal to rooms
	fmt.Println("boolean expression : 'roomsOccupied == rooms' is equal to :")
	fmt.Println(roomsOccupied == rooms) //*\label{condExpBool3}

	roomsAvailable = rooms - roomsOccupied
	if roomsAvailable == 0 {
		fmt.Println("No available rooms now")
	} else {
		fmt.Println(roomsAvailable, "room available")
	}

	a := 21
	b := 42

	fmt.Println(a == b) // false
	fmt.Println(a == a) // true
	fmt.Println(a != b) // true
	fmt.Println(a > b)  // false
	fmt.Println(a < b)  // true
	fmt.Println(a <= b) // true
	fmt.Println(a <= a) // true
	fmt.Println(a >= a) // true
}
