package models

import (
	"galon-app/database"
	"time"
)

type Galon struct {
	ID            int       `json:"galon_id"`
	Brandname     string    `json:"brand_name"`
	Stock         int       `json:"stock"`
	UpdatestockAt time.Time `json:"updatestok_at"`
	CreatedAt     time.Time `json:"created_at"`
}
type Gallons []Galon

func (g *Galon) AddStock() error {
	sqlStatement := `INSERT INTO item_galon (brand_name,stock)
					VALUES ($1,$2)
					Returning *`
	err := database.DB.QueryRow(sqlStatement, g.Brandname, g.Stock).
		Scan(
			&g.ID,
			&g.Brandname,
			&g.Stock,
			&g.UpdatestockAt,
			&g.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (g *Galon) UpdateStock() error {
	sqlStatement := `UPDATE item_galon
					SET stock=$2, updatestock_at=$3
					WHERE id = $1
					Returning *`
	err := database.DB.QueryRow(sqlStatement, g.ID, g.Stock, time.Now()).
		Scan(&g.ID, &g.Brandname, &g.Stock, &g.UpdatestockAt, &g.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (g *Galon) DeleteGalon() error {
	sqlStatement := `DELETE from item_galon WHERE id=$1`
	_, err := database.DB.Exec(sqlStatement, g.ID)
	if err != nil {
		return err
	}
	return nil
}
func (g *Gallons) GetAll() error {
	sqlStatement := `SELECT * FROM item_galon ORDER BY id`

	rows, err := database.DB.Query(sqlStatement)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var galon = Galon{}
		err = rows.Scan(
			&galon.ID,
			&galon.Brandname,
			&galon.Stock,
			&galon.UpdatestockAt,
			&galon.CreatedAt,
		)

		if err != nil {
			return err
		}

		*g = append(*g, galon)
	}

	return nil
}
func (g *Galon) GetById() error {
	sqlStatement := `SELECT * FROM item_galon WHERE id = $1`
	err := database.DB.QueryRow(sqlStatement, g.ID).
		Scan(&g.ID, &g.Brandname, &g.Stock, &g.UpdatestockAt, &g.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}
