package user

import "time"

type RegisterForm struct {
	Name     string
	Email    string
	Password string
}

type User struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}
