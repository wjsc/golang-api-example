package models

type Guitar struct {
	Id    int    `form:"id" json:"id" xml:"id"  binding:"required"`
	Year  int    `form:"year" json:"year" xml:"year"  binding:"required"`
	Model string `form:"model" json:"model" xml:"model"  binding:"required"`
}
