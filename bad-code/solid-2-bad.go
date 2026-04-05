package bad_code

import "fmt"

type RideService struct{}

func (r *RideService) CalculateFare(rideType string, distance float64, time int, isPeak bool) float64 {

	var fare float64

	if rideType == "BIKE" {
		fare = distance * 5
	} else if rideType == "CAR" {
		fare = distance * 10
	} else if rideType == "SUV" {
		fare = distance * 15
	} else {
		panic("invalid ride type")
	}

	// time-based charges
	if time > 60 {
		fare += 50
	}

	// surge pricing
	if isPeak {
		fare *= 1.5
	}

	// discount
	if rideType == "BIKE" && distance > 10 {
		fare *= 0.9
	}

	fmt.Println("Final Fare:", fare)
	return fare
}
