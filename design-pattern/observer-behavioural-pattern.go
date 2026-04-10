package design_pattern

type Observer interface {
	Update(symbol string, price float64)
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	NotifyUser()
}

type Stock struct {
	observers []Observer
	symbol    string
	price     float64
}

func NewStock(symbol string) *Stock {
	return &Stock{
		symbol:    symbol,
		observers: []Observer{},
	}
}

func (s *Stock) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Stock) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s.symbol, s.price)
	}
}

func (s *Stock) Detach(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Stock) SetPrice(price float64) {
	s.price = price
	s.NotifyObservers()
}

type PriceDisplay struct{}

func (PriceDisplay) Update(symbol string, price float64) {
	_ = price
	_ = symbol
}

type PriceAlert struct {
	threshold float64
}

func NewPriceAlert(threshold float64) *PriceAlert {
	return &PriceAlert{
		threshold: threshold,
	}
}

func (p *PriceAlert) Update(symbol string, price float64) {
	if price > p.threshold {

	}
	_ = symbol
}

func main() {
	stock := NewStock("$")
	display := PriceDisplay{}
	alert := NewPriceAlert(230)
	stock.Attach(display)
	stock.Attach(alert)
	stock.SetPrice(123)
	stock.SetPrice(234)
}
