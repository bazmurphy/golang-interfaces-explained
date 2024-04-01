package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type ShopDB struct {
	*sql.DB
}

func (sdb *ShopDB) CountCustomers(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM customers WHERE timestamp > $1", since).Scan(&count)
	return count, err
}

func (sdb *ShopDB) CountSales(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)
	return count, err
}

func main() {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	shopDB := &ShopDB{db}
	sr, err := calculateSalesRate(shopDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sr)
}

func calculateSalesRate(sdb *ShopDB) (string, error) {
	since := time.Now().Add(-24 * time.Hour)

	sales, err := sdb.CountSales(since)
	if err != nil {
		return "", err
	}

	customers, err := sdb.CountCustomers(since)
	if err != nil {
		return "", err
	}

	rate := float64(sales) / float64(customers)
	return fmt.Sprintf("%.2f", rate), nil
}

// Now, what if we want to create a unit test for the calculateSalesRate() function to make sure that the math logic in it is working correctly?

// Currently this is a bit of a pain. We would need to set up a test instance of our PostgreSQL database, along with setup and teardown scripts to scaffold the database with dummy data. That's quite lot of work when all we really want to do is test our math logic.

// So what can we do? You guessed it — interfaces to the rescue!
