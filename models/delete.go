package models

import "github.com/leonardodelira/go-postgres/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE todos WHERE id=$4`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}