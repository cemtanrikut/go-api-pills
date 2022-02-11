package user

import "time"

type user struct {
	name      string
	email     string
	password  string
	gender    string
	birthdate time.Time
	active    bool
}
