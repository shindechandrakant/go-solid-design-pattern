package best_code

import "fmt"

type RideType string

const (
	CAR  RideType = "CAR"
	BIKE RideType = "BIKE"
	SUV  RideType = "SUV"
)

type RideContext struct {
	Distance float64
	Time     int
	IsPeak   bool
}

type FareStrategy interface {
	calculate(float64) float64
}

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

var FareRegistry = map[RideType]func() FareStrategy{
	BIKE: func() FareStrategy { return &BikeRideType{} },
	CAR:  func() FareStrategy { return &CarRideType{} },
	SUV:  func() FareStrategy { return &SuvRideType{} },
}

func FareFactory(rideType RideType) (FareStrategy, error) {
	if creator, ok := FareRegistry[rideType]; ok {
		return creator(), nil
	}
	return nil, fmt.Errorf("invalid ridetype provided")
}

type PricingRule interface {
	Apply(float64, RideContext) float64
}

type TimeRule struct{}

func (t *TimeRule) Apply(fare float64, ctx RideContext) float64 {
	if ctx.Time > 60 {
		return fare + 50
	}
	return fare
}

type SurchargeRule struct{}

func (s *SurchargeRule) Apply(fare float64, ctx RideContext) float64 {
	if ctx.IsPeak {
		return fare * 1.5
	}
	return fare
}

type BikeDiscountRule struct{}

func (d *BikeDiscountRule) Apply(fare float64, ctx RideContext) float64 {
	if ctx.Distance > 10 {
		return fare * 0.9
	}
	return fare
}

type RideService struct {
	fareStrategy FareStrategy
	rules        []PricingRule
}

func NewRideService(strategy FareStrategy, rules []PricingRule) *RideService {
	return &RideService{
		fareStrategy: strategy,
		rules:        rules,
	}
}

func (r *RideService) CalculateFare(ctx RideContext) float64 {
	fare := r.fareStrategy.calculate(ctx.Distance)

	for _, rule := range r.rules {
		fare = rule.Apply(fare, ctx)
	}
	return fare
}

func main() {

	fareStrategy, err := FareFactory(BIKE)
	if err != nil {
		panic(err)
	}

	rules := []PricingRule{
		&TimeRule{},
		&SurchargeRule{},
		&BikeDiscountRule{},
	}

	service := NewRideService(fareStrategy, rules)
	ctx := RideContext{
		Distance: 20,
		Time:     80,
		IsPeak:   true,
	}

	fare := service.CalculateFare(ctx)
	fmt.Println("Final Fare: ", fare)
}
