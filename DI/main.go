package main

// import (
// 	"database/sql"
// 	"fmt"

// 	mgo "gopkg.in/mgo.v2"
// )

// /*
// 	A trade  processor class should  not depend on an interface like Trade service , rather than a concreate implementation
// 	like SqlServiceTradeResponsiblity.

// 	This way d/t implementaiton of TradeService interface can be swapped in and out with out affecting  the trade Processor class ,
// 	which can make it easier to maintain and test .

// 	MongoDBTradeRepository could be used instead of the SQLServiceTradeRepository without modifying the TradeProcesor class
// */

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	mgo "gopkg.in/mgo.v2"
)

type Trade struct {
	ID       int
	Symbol   string
	Quantity float64
	Price    float64
}

type TradeService interface {
	Save(trade *Trade) error
}

type TradeProcessor struct {
	tradeService TradeService
}

func (tp *TradeProcessor) Process(trade *Trade) error {
	err := tp.tradeService.Save(trade)
	if err != nil {
		return err
	}
	// process trade
	return nil
}

type SQLServiceTradeRepository struct {
	db *sql.DB
}

func (str *SQLServiceTradeRepository) Save(trade *Trade) error {
	_, err := str.db.Exec("INSERT INTO trade (trade_id, symbol, quantity, price) VALUES (?, ?, ?, ?)", trade.ID, trade.Symbol, trade.Quantity, trade.Price)
	if err != nil {
		return err
	}

	return nil
}

type MongoDBTradeRepository struct {
	session *mgo.Session
}

func (mdtr *MongoDBTradeRepository) Save(trade *Trade) error {
	collection := mdtr.session.DB("trades").C("trade")

	err := collection.Insert(trade)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Setup SQLite repository
	sqlDb, err := sql.Open("sqlite3", "./trades.db")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDb.Close()

	// Create table if not exists
	_, err = sqlDb.Exec(`CREATE TABLE IF NOT EXISTS trade (
		trade_id INTEGER PRIMARY KEY,
		symbol TEXT,
		quantity REAL,
		price REAL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	sqlRepo := &SQLServiceTradeRepository{db: sqlDb}
	sqlTradeProcessor := &TradeProcessor{tradeService: sqlRepo}

	// Process a trade using SQL repository
	sqlTrade := &Trade{
		ID:       2,
		Symbol:   "AAPL",
		Quantity: 10,
		Price:    150.0,
	}

	err = sqlTradeProcessor.Process(sqlTrade)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Trade processed and saved to SQLite")

	// Setup MongoDB repository
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer mongoSession.Close()

	mongoRepo := &MongoDBTradeRepository{session: mongoSession}
	mongoTradeProcessor := &TradeProcessor{tradeService: mongoRepo}

	// Process a trade using MongoDB repository
	mongoTrade := &Trade{
		ID:       3,
		Symbol:   "GOOGL",
		Quantity: 5,
		Price:    2800.0,
	}

	err = mongoTradeProcessor.Process(mongoTrade)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Trade processed and saved to MongoDB")
}
