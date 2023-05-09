package database

import (
	"context"
	"miniProject/models"
)

func GetArrives(ctx context.Context) ([]models.Arrive, error) {
	var arrives []models.Arrive

	err := DB.WithContext(ctx).Find(&arrives).Error
	if err != nil {
		return nil, err
	}

	return arrives, nil
}

func GetArrivesByID(ctx context.Context, id int) (models.Arrive, error) {
	var arrives models.Arrive

	err := DB.WithContext(ctx).Where("id = ?", id).First(&arrives).Error
	if err != nil {
		return models.Arrive{}, err
	}

	return arrives, nil
}

func CreateArrives(ctx context.Context, arrives models.Arrive) (models.Arrive, error) {
	err := DB.WithContext(ctx).Create(&arrives).Error

	if err != nil {
		return models.Arrive{}, err
	}

	return arrives, nil
}

func DeleteArrives(ctx context.Context, id int) error {
	var arrives models.Arrive

	result := DB.WithContext(ctx).Where("id = ?", id).Delete(&arrives)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func UpdateArrives(ctx context.Context, id int, arrives models.Arrive) (interface{}, error) {
	result := DB.WithContext(ctx).Model(&models.Arrive{}).Where("id = ?", id).Updates(arrives)
	if result.Error != nil {
		return arrives, result.Error
	}

	if result.RowsAffected == 0 {
		return arrives, ErrIDNotFound
	}

	return arrives, nil
}
