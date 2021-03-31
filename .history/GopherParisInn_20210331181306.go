package GopherParisInn

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const hotelName = "The Gopher Paris Inn"
	const rooms = 134
	const firstRoomNumber = 110

	rand.Seed(time.Now().UTC().UnixNano())
	roomsOccupied := rand.Intn(rooms)
	roomsAvailable := rooms - roomsOccupied

	occupancyRate := (float32(roomsOccupied) / float32(rooms)) * 100
	var occupancyLevel string
	if occupancyRate > 70 {
		occupancyLevel = "High"
	} else if occupancyRate > 20 {
		occupancyLevel = "Medium"
	} else {
		occupancyLevel = "Low"
	}

	// OccupancyLevel : switch case alternative
	switch {
	case occupancyRate > 70:
		occupancyLevel = "High"
	case occupancyRate > 20:
		occupancyLevel = "Medium"
	default:
		occupancyLevel = "Low"
	}
	// 	occupancyLevel := "Low"
	// switch {
	// case occupancyRate > 70:
	// 	occupancyLevel = "High"
	// case occupancyRate > 20:
	// 	occupancyLevel = "Medium"
	// }

	fmt.Println("Hotel:", hotelName)
	fmt.Println("Number of rooms:", rooms)
	fmt.Println("Rooms available", roomsAvailable)
	fmt.Println("               Occupancy Level:", occupancyLevel)
	fmt.Printf("               Occupancy Rate:%0.2f %%\n", occupancyRate)

	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		// for i := 0; i < roomsAvailable; i++ {
		for i := 0; roomsAvailable > i; i++ {
			roomNumber := firstRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
		}
	} else {
		fmt.Println("No rooms availbale for tonight")
	}

	// switch occupancyRate {
	// case < 30%:

	// }

}
