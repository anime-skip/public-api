package models

type User struct {
	BaseModelSoftDelete
	username string
	email    string
}
