package storage
import (
	models "app-v1/models"
	"errors"
)

type GuitarStorageType struct {
	elements []models.Guitar
}

var GuitarStorage GuitarStorageType

func (guitarStorage GuitarStorageType) Get() ([]models.Guitar, error) {
	return GuitarStorage.elements, nil
}

func (guitarStorage GuitarStorageType) GetById(Id int) (models.Guitar , error) {

	for _, currentGuitar := range GuitarStorage.elements {
		if currentGuitar.Id == Id {
			return currentGuitar, nil
		}
	}
	return models.Guitar{}, errors.New("Guitar not found")
}

func (guitarStorage GuitarStorageType) Create(guitar models.Guitar) (models.Guitar, error) {
	GuitarStorage.elements = append(GuitarStorage.elements, guitar)
	return guitar, nil
}


func (guitarStorage GuitarStorageType) Update(guitar models.Guitar)(models.Guitar, error) {
	
	for index, currentGuitar := range GuitarStorage.elements {
		if currentGuitar.Id == guitar.Id {
			GuitarStorage.elements[index] = guitar
			return guitar, nil
		}
	}
	return models.Guitar{}, errors.New("Guitar not found")
}


func (guitarStorage GuitarStorageType) Delete(Id int)(error) {
	for index, currentGuitar := range GuitarStorage.elements {

		if currentGuitar.Id == Id {
			GuitarStorage.elements = append(GuitarStorage.elements[:index], GuitarStorage.elements[index+1:]...)
			return nil
		}
	}
	return errors.New("Guitar not found")
}
