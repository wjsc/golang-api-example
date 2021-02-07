package storage
import (
	models "app-v1/models"
	"errors"
)

type GuitarStorageType struct {
	elements []models.GuitarModel
}

func (guitarStorage *GuitarStorageType) Get() ([]models.GuitarModel, error) {
	return guitarStorage.elements, nil
}

func (guitarStorage *GuitarStorageType) GetById(Id int) (models.GuitarModel , error) {

	for _, currentGuitar := range guitarStorage.elements {
		if currentGuitar.Id == Id {
			return currentGuitar, nil
		}
	}
	return models.GuitarModel{}, errors.New("Guitar not found")
}

func (guitarStorage *GuitarStorageType) Create(guitar models.GuitarModel) (models.GuitarModel, error) {
	guitarStorage.elements = append(guitarStorage.elements, guitar)
	return guitar, nil
}


func (guitarStorage *GuitarStorageType) Update(guitar models.GuitarModel)(models.GuitarModel, error) {
	
	for index, currentGuitar := range guitarStorage.elements {
		if currentGuitar.Id == guitar.Id {
			guitarStorage.elements[index] = guitar
			return guitar, nil
		}
	}
	return models.GuitarModel{}, errors.New("Guitar not found")
}


func (guitarStorage *GuitarStorageType) Delete(Id int)(error) {
	for index, currentGuitar := range guitarStorage.elements {

		if currentGuitar.Id == Id {
			guitarStorage.elements = append(guitarStorage.elements[:index], guitarStorage.elements[index+1:]...)
			return nil
		}
	}
	return errors.New("Guitar not found")
}
