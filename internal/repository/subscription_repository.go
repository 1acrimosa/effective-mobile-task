package repository

import (
	"effective-mobile-task/internal/model"
	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func (r *ItemRepository) FindAll() ([]model.Item, error) {
	var items []model.Item
	r.DB.Find(&items)
	return items, nil
}

func (r *ItemRepository) FindById(id uint) (model.Item, error) {
	var item model.Item
	r.DB.First(&item, id)
	return item, nil
}

func (r *ItemRepository) Create(item *model.Item) error {
	return r.DB.Create(item).Error
}

func (r *ItemRepository) Update(item *model.Item) error {
	return r.DB.Save(item).Error
}

func (r *ItemRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Item{}, id).Error
}
