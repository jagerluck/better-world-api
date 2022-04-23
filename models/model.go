package main

import (
	"database/sql"
	"reflect"
	"time"
	// "errors"
)

type pin struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Address   string    `json: "address"`
	Lat       float32   `json:"lat"`
	Lng       float32   `json:"lng"`
}

func getPins(db *sql.DB, order string) ([]pin, error) {
	rows, err := db.Query("SELECT * FROM pin ORDER BY $1",
		order)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pins := []pin{}

	for rows.Next() {
		var p pin
		// mb some generic handling instead
		if err := rows.Scan(&p.ID); err != nil {
			return nil, err
		}
		pins = append(pins, p)
	}

	return pins, nil
}

func getPin(db *sql.DB) (p *pin) {
	err := db.QueryRow(`
	SELECT address FROM pin
	LEFT JOIN pin_photo AS pp
	LEFT JOIN pin_video AS pv
	ON pp.fk_pin_id = pv.fk_pin_id
	WHERE id=$1
	`, p.ID).Scan(p)

	if err != nil {
		panic(err)
	}

	return p
}

func (p *pin) updatePin(db *sql.DB) error {
	fields := reflect.VisibleFields(p)
	var updateStr string
	for _, f := range fields {

	}
	_, err := db.Exec(`UPDATE pin SET name=$1, price=$2 WHERE id=$3`,
		p.Name, p.Price, p.ID)

	if err != nil {
		panic(err)
	}

	return p
}

func (p *pin) deletePin(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

	return err
}

func (p *pin) createPin(db *sql.DB) error {
	// err := db.QueryRow(
	// 	"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
	// 	p.Name, p.Price).Scan(&p.ID)

	createdPin := db(p)

	if err != nil {
		return err
	}

	return nil
}
