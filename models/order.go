package models

import (
	"galon-app/database"
	"time"
)

type Order struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	GalonID    int       `json:"galon_id"`
	TotalOrder int       `json:"total_order"`
	Status     string    `json:"status"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

func (o *Order) MakeOrder() error {
	sqlStatement := `
		INSERT INTO orders (user_id,galon_id, total_order, status)
		VALUES ($1,$2,$3,$4)
		Returning *
	`
	err := database.DB.QueryRow(sqlStatement, o.UserID, o.GalonID, o.TotalOrder, o.Status).
		Scan(&o.ID,
			&o.UserID,
			&o.GalonID,
			&o.TotalOrder,
			&o.Status,
			&o.UpdatedAt,
			&o.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (o *Order) UpdateStatus() error {
	sqlStatement := `UPDATE orders SET status=$2, updated_at=$3
					WHERE id=$1
					Returning *`

	err := database.DB.QueryRow(sqlStatement, o.ID, o.Status, time.Now()).
		Scan(&o.ID,
			&o.UserID,
			&o.GalonID,
			&o.TotalOrder,
			&o.Status,
			&o.UpdatedAt,
			&o.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (o *Order) GetById() error {
	sqlStatement := `SELECT * FROM orders WHERE id = $1`
	err := database.DB.QueryRow(sqlStatement, o.ID).
		Scan(&o.ID,
			&o.UserID,
			&o.GalonID,
			&o.TotalOrder,
			&o.Status,
			&o.UpdatedAt,
			&o.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}

type Orders []Order

func (o *Orders) GetByUserId(userID int) error {
	sqlStatement := `SELECT * FROM orders WHERE user_id = $1 ORDER BY id DESC`
	rows, err := database.DB.Query(sqlStatement, userID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var order = Order{}
		err = rows.Scan(
			&order.UserID,
			&order.GalonID,
			&order.TotalOrder,
			&order.Status,
			&order.UpdatedAt,
			&order.CreatedAt)
		if err != nil {
			return err
		}
		*o = append(*o, order)
	}
	return nil
}
func (o *Orders) GetAll() error {
	sqlStatement := `SELECT * FROM orders ORDER BY id DESC`
	rows, err := database.DB.Query(sqlStatement)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var order = Order{}
		err = rows.Scan(
			&order.UserID,
			&order.GalonID,
			&order.TotalOrder,
			&order.Status,
			&order.UpdatedAt,
			&order.CreatedAt)
		if err != nil {
			return err
		}
		*o = append(*o, order)
	}
	return nil
}
