package main

type Guitar struct {
	Id    int    `form:"id" json:"id" xml:"id"  binding:"required"`
	Year  int    `form:"year" json:"year" xml:"year"  binding:"required"`
	Model string `form:"model" json:"model" xml:"model"  binding:"required"`
}

func (guitar *Guitar) get(id int) (result bool, err error) {
	return true, nil
}

func (guitar *Guitar) insert() (result bool, err error) {
	return true, nil
}

func (guitar *Guitar) update() (result bool, err error) {
	return true, nil
}
