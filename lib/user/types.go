package user

import "time"

type UserStruct struct {
	Id string
	FullName string
	CreatedAt time.Time
	BirthDate time.Time
	Status bool
}

type UserSaveStruct struct{
	FullName string
	BirthDate string
}