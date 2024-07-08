package main

/*
In a trade class should be subtype of a trade class, which means that it
should be a subtype of a trade class without causing any issues.

Example if a trade Processor class 	expects a trade object , but recieves a
Future trade  object, it should still be able to process the trade wtih out any issue
*/
type Trade interface {
	Process() error
}

type FutureTrade struct {
	Trade
}

func (ft *FutureTrade) Process() error {
	// Process future trade
	return nil
}
