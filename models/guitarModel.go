package models

type GuitarModel struct {
	Id    int    `json:"id" binding:"required"`
	Year  int    `json:"year" binding:"required"`
	Model string `json:"model" binding:"required"`
}
