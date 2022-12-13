package app

import "github.com/taujago/gotoko/models"

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
	}
}
