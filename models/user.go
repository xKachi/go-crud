package models

import (

	"example.com/rest-api/db"
)

type User struct {
	ID          int64
	Name        string    `binding:"required"`
	Email 		string    `binding:"required"`
}

//var users = []User{}

func (u User) Save() (User, error) {
	query := `
	INSERT INTO users(name, email) 
	VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Name, u.Email)
	if err != nil {
		return User{}, err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return u, err
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(id int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (user User) Update() error {
	query := `
	UPDATE events
	SET name = ?, email = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email)
	return err
}

func (user User) Delete() error {
	query := "DELETE FROM users WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	return err
}
