package repository

import (
	"errors"

	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func (rp *BaseRepository[T]) GetById(id int) (*T, error) {
	var entity *T

	result := rp.db.First(entity, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return entity, nil
}

func (rp *BaseRepository[T]) GetMany(limit int, offset int) (*[]T, error) {
	var entities []T

	result := rp.db.Limit(limit).Offset(offset).Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (rp *BaseRepository[T]) GetAll() (*[]T, error) {
	return rp.GetMany(-1, -1)
}

func (rp *BaseRepository[T]) Create(entity *T) (*T, error) {
	result := rp.db.Create(entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

func (rp *BaseRepository[T]) CreateBatch(entities []T) ([]T, error) {
	result := rp.db.Create(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func (rp *BaseRepository[T]) Update(entity *T) (*T, error) {
	result := rp.db.Save(entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

func (rp *BaseRepository[T]) Delete(entity *T) error {
	result := rp.db.Delete(entity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
