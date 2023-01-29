package models

import (
	"database/sql"
	
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Created_at sql.NullString `json:"created_at"`
}

func(u *User) CreateUser(db *sql.DB) error{
	_, err := db.Exec("INSERT INTO users (name, email, password) VALUES($1, $2, $3)", u.Name, u.Email, u.Password)
	return err
}

func(u *User) GetUser(db *sql.DB) error {
	err := db.QueryRow("SELECT * FROM users WHERE id=$1",u.Id).Scan(&u.Id,&u.Name, &u.Email, &u.Password, &u.Created_at)
	return err
}

func(u *User) DeleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", u.Id)
	return err
}

func(u *User) UpdateUser(db *sql.DB) error {
	_, err := db.Exec("UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4",  u.Name,u.Email,u.Password,u.Id)
	return err
}

func GetUsers(db *sql.DB) ([]User, error) {
	rows, _ := db.Query("SELECT * FROM users")
	defer rows.Close()
	var users []User
  for rows.Next() {
	var user User
	err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at)
	if err != nil {
		return nil,err
	}
	users = append(users, user)
  }
	return users, nil
}