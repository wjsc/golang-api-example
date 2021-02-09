package storage
import (
	"app-v1/models"
	"errors"
)

type GuitarMemoryStorage struct {
	elements []models.GuitarModel
}

func (guitarMemoryStorage *GuitarMemoryStorage) Get() ([]models.GuitarModel, error) {
	if len(guitarMemoryStorage.elements) == 0 {
		guitarMemoryStorage.elements = make([]models.GuitarModel, 0)
	}
	return guitarMemoryStorage.elements, nil
}

func (guitarMemoryStorage *GuitarMemoryStorage) GetById(Id int) (models.GuitarModel , error) {

	for _, currentGuitar := range guitarMemoryStorage.elements {
		if currentGuitar.Id == Id {
			return currentGuitar, nil
		}
	}
	return models.GuitarModel{}, errors.New("Guitar not found")
}

func (guitarMemoryStorage *GuitarMemoryStorage) Create(guitar models.GuitarModel) (models.GuitarModel, error) {
	guitarMemoryStorage.elements = append(guitarMemoryStorage.elements, guitar)
	return guitar, nil
}

func (guitarMemoryStorage *GuitarMemoryStorage) Update(guitar models.GuitarModel)(models.GuitarModel, error) {
	
	for index, currentGuitar := range guitarMemoryStorage.elements {
		if currentGuitar.Id == guitar.Id {
			guitarMemoryStorage.elements[index] = guitar
			return guitar, nil
		}
	}
	return models.GuitarModel{}, errors.New("Guitar not found")
}

func (guitarMemoryStorage *GuitarMemoryStorage) Delete(Id int) (error) {
	for index, currentGuitar := range guitarMemoryStorage.elements {

		if currentGuitar.Id == Id {
			guitarMemoryStorage.elements = append(guitarMemoryStorage.elements[:index], guitarMemoryStorage.elements[index+1:]...)
			return nil
		}
	}
	return errors.New("Guitar not found")
}
