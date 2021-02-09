package storage
import (
	"app-v1/models"
)

type IGuitarStorage interface {
	Get() ([]models.GuitarModel, error)
	GetById(Id int) (models.GuitarModel , error)
	Create(guitar models.GuitarModel) (models.GuitarModel, error)
	Update(guitar models.GuitarModel) (models.GuitarModel, error)
	Delete(Id int) (error)
}