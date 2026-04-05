package good_code

import "fmt"

type FareStrategy interface {
	calculate(float64) float64
}

type DiscountStrategy interface {
	discount(fair float64, distance float64) float64
}

type RideType string

const (
	CAR  RideType = "CAR"
	BIKE RideType = "BIKE"
	SUV  RideType = "SUV"
)

type BikeRideType struct{}
type SuvRideType struct{}
type CarRideType struct{}

func (bike *BikeRideType) calculate(distance float64) float64 {
	return distance * 5
}

func (bike *BikeRideType) discount(fare float64, distance float64) float64 {
	if distance > 10 {
		return fare * 0.9
	}
	return fare
}

func (car *CarRideType) calculate(distance float64) float64 {
	return distance * 10
}

func (suv *SuvRideType) calculate(distance float64) float64 {
	return distance * 15
}

var FareStrategyRegistry = map[RideType]func() FareStrategy{
	BIKE: func() FareStrategy { return &BikeRideType{} },
	CAR:  func() FareStrategy { return &CarRideType{} },
	SUV:  func() FareStrategy { return &SuvRideType{} },
}

var DiscountStrategyRegistry = map[RideType]func() DiscountStrategy{
	BIKE: func() DiscountStrategy { return &BikeRideType{} },
}

func DiscountStrategyFactory(rideType RideType) (DiscountStrategy, error) {
	if creator, ok := DiscountStrategyRegistry[rideType]; ok {
		return creator(), nil
	}
	return nil, fmt.Errorf("invalid ridetype provided")
}

func FareStrategyFactory(rideType RideType) (FareStrategy, error) {
	if creator, ok := FareStrategyRegistry[rideType]; ok {
		return creator(), nil
	}
	return nil, fmt.Errorf("invalid ridetype provided")
}

type RideService struct{}

func TimeBaseCharges(time int) float64 {
	if time > 60 {
		return 50
	}
	return 0
}

func (r *RideService) CalculateFare(rideType RideType, distance float64, time int, isPeak bool) (float64, error) {

	fareStrategy, err := FareStrategyFactory(rideType)
	if err != nil {
		return 0, err
	}

	fare := fareStrategy.calculate(distance)
	fare += TimeBaseCharges(time)
	if isPeak {
		fare *= 1.5
	}
	discountStrategy, err := DiscountStrategyFactory(rideType)
	if err != nil {
		return 0, err
	}
	fare = discountStrategy.discount(fare, distance)
	return fare, nil
}
