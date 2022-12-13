package app

import "github.com/taujago/gotoko/app/models"

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{

		{Model: models.Address{}},
		{Model: models.User{}},
		{Model: models.Product{}},
		{Model: models.ProductImage{}},
		{Model: models.Section{}},
		{Model: models.Category{}},
	}
}
