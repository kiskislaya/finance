package entity

import "time"

type Transaction struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Amount    int64     `db:"amount"`
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"created_at"`
}
