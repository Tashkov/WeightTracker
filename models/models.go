package models

import "database/sql"

var DB *sql.DB

type User struct {
	FirstName string
	LastName  string
	Sex       string
	Age       int64
	HeightCm  int64
}

type WeightLog struct {
	Weight int64
	UserID int64
}

func AllUsers() ([]User, error) {
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usrs []User

	for rows.Next() {
		var usr User

		err := rows.Scan(&usr.FirstName, &usr.LastName, &usr.Sex, &usr.Age, &usr.HeightCm)
		if err != nil {
			return nil, err
		}

		usrs = append(usrs, usr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usrs, nil

}
