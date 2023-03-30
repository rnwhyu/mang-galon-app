package models

import (
	"galon-app/database"
	token "galon-app/utils"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"user_id"`
	Fullname  string    `json:"full_name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role_id   int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}
type Users []User

func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func LoginCheck(username string, password string) (string, error) {
	u := User{}
	sqlStatement := `SELECT * FROM users WHERE username = $1`
	err := database.DB.QueryRow(sqlStatement, username).
		Scan(
			&u.ID,
			&u.Fullname,
			&u.Username,
			&u.Email,
			&u.Password,
			&u.Role_id,
			&u.CreatedAt)
	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := token.GenerateToken(u.ID, u.Role_id)
	if err != nil {
		return "", err
	}
	return token, nil

}
func (u *User) Create() error {
	hashedPw, errs := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errs != nil {
		return errs
	}
	u.Password = string(hashedPw)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	sqlStatement := `INSERT INTO users (full_name, username, email, password, role_id)
	VALUES($1,$2,$3,$4,$5)
	Returning *
	`
	err := database.DB.
		QueryRow(
			sqlStatement,
			u.Fullname,
			u.Username,
			u.Email,
			u.Password,
			u.Role_id).
		Scan(
			&u.ID,
			&u.Fullname,
			&u.Username,
			&u.Email,
			&u.Password,
			&u.Role_id,
			&u.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
func GetUserByID(id int) (User, error) {
	var u User
	sqlStatement := `SELECT * FROM users WHERE id = $1`
	err := database.DB.QueryRow(sqlStatement, id).
	Scan(
		&u.ID,
		&u.Fullname,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.Role_id,
		&u.CreatedAt)
	if err != nil{
		return u, err
	}
	u.PrepareGive()
	return u, nil
}
