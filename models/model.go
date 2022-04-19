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

//  // fails
// func (p *pin) getPin(db *sql.DB) error {
// 	return db.QueryRow(`
// 	SELECT address FROM pin
// 	LEFT JOIN pin_photo AS pp
// 	LEFT JOIN pin_video AS pv
// 	ON pp.fk_pin_id = pv.fk_pin_id
// 	WHERE id=$1
// 	`, p.ID)
// }

// func (p *product) updatePin(db *sql.DB) error {
// 	_, err :=
// 		db.Exec("UPDATE pin SET name=$1, price=$2 WHERE id=$3",
// 			p.Name, p.Price, p.ID)

// 	return err
// }

// func (p *product) deletePin(db *sql.DB) error {
// 	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

// 	return err
// }

// func (p *product) createPin(db *sql.DB) error {
// 	err := db.QueryRow(
// 		"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
// 		p.Name, p.Price).Scan(&p.ID)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
