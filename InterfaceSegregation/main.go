package main

type Trade interface {
	Process() error
}

type OptionTrade interface {
	CalculateImpliedVolatility() error
}

type FutureTrade struct {
	Trade
}

func (ft *FutureTrade) Process() error {
	// process future trade
	return nil
}

type OptionTrade struct {
	Trade
}

func (ot *OptionTrade) Process() error {
	// process option trade
	return nil
}

func (ot *OptionTrade) CalculateImpliedVolatility() error {
	// calculate implied volatility
	return nil
}
