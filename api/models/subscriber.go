package models

type Subscriber struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}
