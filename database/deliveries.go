package database

import (
	"context"
	"miniProject/models"
)

func GetDeliveries(ctx context.Context) ([]models.Delivery, error) {
	var deliveries []models.Delivery

	err := DB.WithContext(ctx).Find(&deliveries).Error
	if err != nil {
		return nil, err
	}

	return deliveries, nil
}

func GetDeliveriesByID(ctx context.Context, id int) (models.Delivery, error) {
	var deliveries models.Delivery

	err := DB.WithContext(ctx).Where("id = ?", id).First(&deliveries).Error
	if err != nil {
		return models.Delivery{}, err
	}

	return deliveries, nil
}

func CreateDeliveries(ctx context.Context, deliveries models.Delivery) (models.Delivery, error) {
	err := DB.WithContext(ctx).Omit("Process.*").Create(&deliveries).Error

	if err != nil {
		return models.Delivery{}, err
	}

	var NewDelivery models.Delivery
	err = DB.WithContext(ctx).Preload("Process").Preload("Process.Arrives").Where("id = ?", deliveries.ID).First(&NewDelivery).Error
	if err != nil {
		return models.Delivery{}, err
	}

	return NewDelivery, nil
}

func DeleteDeliveries(ctx context.Context, id int) error {
	var deliveries models.Delivery

	result := DB.WithContext(ctx).Where("id = ?", id).Delete(&deliveries)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func UpdateDeliveries(ctx context.Context, id int, deliveries models.Delivery) (models.Delivery, error) {
	result := DB.WithContext(ctx).Model(&models.Delivery{}).Where("id = ?", id).Updates(deliveries)
	if result.Error != nil {
		return models.Delivery{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.Delivery{}, ErrIDNotFound
	}

	return deliveries, nil
}
