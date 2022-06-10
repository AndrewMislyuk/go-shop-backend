package domain

import "time"

type CategoriesList struct {
	Id        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateCategoryInput struct {
	Name *string `json:"name"`
}
