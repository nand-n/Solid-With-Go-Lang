package main

/*
A Trade processor class should be designed to be open for extension but cloased
for modification . This means that if new trade types  are added, the tade processor
class should be able to handle them without needing to modify the existing code.
This can be achieved by defining an interface for processing trades and implementing it
for each trade type
*/

type Trade struct {
	ID       int
	Symbol   string
	Quantity float64
	Price    float64
}

type TradeProcessor interface {
	Process(trade *Trade) error
}

type FutureTradeProcessor struct{}

func (ftp *FutureTradeProcessor) Process(trade *Trade) error {
	//process future trade
	return nil
}

type OptionTradeProcessor struct{}

func (otp *OptionTradeProcessor) Process(trade *Trade) error {
	//Process trade
	return nil
}
