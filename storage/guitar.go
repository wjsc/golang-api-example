package storage
import (
	models "app-v1/models"
	"errors"
)

type GuitarStorage struct {
	elements []models.GuitarModel
}

func (guitarStorage *GuitarStorage) Get() ([]models.GuitarModel, error) {
	return guitarStorage.elements, nil
}

func (guitarStorage *GuitarStorage) GetById(Id int) (models.GuitarModel , error) {

	for _, currentGuitar := range guitarStorage.elements {
		if currentGuitar.Id == Id {
			return currentGuitar, nil
		}
	}
	return models.GuitarModel{}, errors.New("Guitar not found")
}

func (guitarStorage *GuitarStorage) Create(guitar models.GuitarModel) (models.GuitarModel, error) {
	guitarStorage.elements = append(guitarStorage.elements, guitar)
	return guitar, nil
}


func (guitarStorage *GuitarStorage) Update(guitar models.GuitarModel)(models.GuitarModel, error) {
	
	for index, currentGuitar := range guitarStorage.elements {
		if currentGuitar.Id == guitar.Id {
			guitarStorage.elements[index] = guitar
			return guitar, nil
		}
	}
	return models.GuitarModel{}, errors.New("Guitar not found")
}


func (guitarStorage *GuitarStorage) Delete(Id int)(error) {
	for index, currentGuitar := range guitarStorage.elements {

		if currentGuitar.Id == Id {
			guitarStorage.elements = append(guitarStorage.elements[:index], guitarStorage.elements[index+1:]...)
			return nil
		}
	}
	return errors.New("Guitar not found")
}
