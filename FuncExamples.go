package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const hotelName = "Golang Hotel"
	const totalRooms = 134
	const firstRoomNumber = 110

	rand.Seed(time.Now().UTC().UnixNano())
	roomsOccupied := rand.Intn(totalRooms)
	roomsAvailable := totalRooms - roomsOccupied

	occupancyRate := occupancyRate(roomsOccupied, totalRooms)
	occupancyLevel := occupancyLevel(occupancyRate)

	fmt.Println("Hotel:", hotelName)
	fmt.Println("Number of rooms", totalRooms)
	fmt.Println("Rooms available", roomsAvailable)
	fmt.Println("                  Occupancy Level:", occupancyLevel)
	fmt.Printf("                  Occupancy Rate: %0.2f %%\n", occupancyRate)

	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; roomsAvailable > i; i++ {
			roomNumber := firstRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			printsroomDetails(roomNumber, size, nights)
		}
	} else {
		fmt.Println("No room available for tonight")
	}

	printHeader()

	fmt.Println("Vacant rooms :", vacantRooms())
	johnPrice := computePrice(145.90, 3)   //*\label{funcEx1Us1}
	paulPrice := computePriceV2(26.32, 10) //*\label{funcEx1Us2}
	robPrice := computePriceV3(189.21, 2)  //*\label{funcEx1Us3}

	total := johnPrice + paulPrice + robPrice
	fmt.Printf("Total : %0.2f $", total)
}

// compute the price of a room based on its rate per night
// and the number of night
func computePrice(rate float32, nights int) float32 {
	n := float32(nights)
	return rate * n
}

// Go gives you the ability to give a name to results. In the previous section, we had an anonymous result. We knew only the type. Here we specify its name and its type.
// Note also that we simply call return. We could have written return price, but it’s not necessary when you name parameters.
func computePriceV2(rate float32, nights int) (price float32) {
	price = rate * float32(nights)
	return
}

// compute the price with a 200% margin
func computePriceV3(rate float32, nights int) (price float32) {
	p := rate * float32(nights)
	p = p * 2
	return
}

// 0 parameter, 1 result
func vacantRooms() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(100)
}

// 0 parameter, 0 result
func printHeader() {
	fmt.Println("Hotel Golang")
	fmt.Println("San Francisco, CA")
}

func occupancyRate(roomsOccupied int, totalRooms int) float32 {
	return (float32(roomsOccupied) / float32(totalRooms)) * 100
}

func printsroomDetails(roomNumber, size, nights int) {
	night := "nights"
	if nights == 1 {
		night = "night"
	}
	fmt.Println(-roomNumber, ":", size, "people /", nights, night)
}

// display information about a room
// This alternative works, but the code is duplicated...
func printRoomDetails2(roomNumber, size, nights int) {
	if nights == 1 {
		fmt.Println(roomNumber, ":", size, "people /", nights, "night")
	} else {
		// practically the same line as before!
		// avoid this
		fmt.Println(roomNumber, ":", size, "people /", nights, "nights")
	}
}

func occupancyLevel(occupancyRate float32) string {
	switch {
	case occupancyRate > 70:
		return "High"
	case occupancyRate > 20:
		return "Medium"
	default:
		return "Low"
	}
}

// From STDLib
// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
// func TrimSpace(s string) string {
// 	// Fast path for ASCII: look for the first ASCII non-space byte
// 	start := 0
// 	for ; start < len(s); start++ {
// 		c := s[start]
// 		if c >= utf8.RuneSelf {
// 			// If we run into a non-ASCII byte, fall back to the
// 			// slower Unicode-aware method on the remaining bytes
// 			return TrimFunc(s[start:], unicode.IsSpace)
// 		}
// 		if asciiSpace[c] == 0 {
// 			break
// 		}
// 	}

// 	// Now look for the first ASCII non-space byte from the end
// 	stop := len(s)
// 	for ; stop > start; stop-- {
// 		c := s[stop-1]
// 		if c >= utf8.RuneSelf {
// 			return TrimFunc(s[start:stop], unicode.IsSpace)
// 		}
// 		if asciiSpace[c] == 0 {
// 			break
// 		}
// 	}

// 	// At this point s[start:stop] starts and ends with an ASCII
// 	// non-space bytes, so we're done. Non-ASCII cases have already
// 	// been handled above.
// 	return s[start:stop]
// }
