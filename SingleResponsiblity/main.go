package main

/*
A trade class should be represponsible for storing and processing trade data.
Another class like TradeValidator , could be responsible for validating trade based on bussiness rules.
By separating these concerns , each class can be more teasily tested 	and maintained
*/
import (
	"database/sql"
	"errors"
)

type Trade struct {
	ID       int
	Symbol   string
	Quantity float64
	Price    float64
}

type TradeReppository struct {
	db *sql.DB
}

func (tr *TradeReppository) Save(trade *Trade) error {
	_, err := tr.db.Exec("INSERT INTO trades (trade_id, symbol, quantity, price ) VALUES (?,?,?,?)", trade.ID, trade.Symbol, trade.Quantity, trade.Price)
	if err != nil {
		return err
	}
	return nil
}

type TradeValidator struct{}

func (tv *TradeValidator) Validate(trade *Trade) error {
	if trade.Quantity <= 0 {
		return errors.New("Trade quantity must be greater than zero")
	}
	if trade.Price < 0 {
		return errors.New("The trade price must be greater than Zero")
	}
	return nil
}
