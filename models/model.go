package main

import (
	"database/sql"
	"time"
	// "errors"
)

type pin struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Address		string 		`json: "address"`
	Lat       float32   `json:"lat"`
	Lng       float32   `json:"lng"`
}

func (p *pin) getPins(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM pin",
		p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) getPin(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) updatePin(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
			p.Name, p.Price, p.ID)

	return err
}

func (p *product) deletePin(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

	return err
}

func (p *product) createPin(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
		p.Name, p.Price).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}
