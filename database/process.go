package database

import (
	"context"
	"miniProject/models"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"gorm.io/gorm"
)

func GetProcess(ctx context.Context) ([]models.Process, error) {
	var process []models.Process

	err := DB.WithContext(ctx).Preload("Arrives").Find(&process).Error
	if err != nil {
		return nil, err
	}

	return process, nil
}

func GetprocessByID(ctx context.Context, id int) (models.Process, error) {
	var process models.Process

	err := DB.WithContext(ctx).Preload("Arrives").Where("id = ?", id).First(&process).Error
	if err != nil {
		return models.Process{}, err
	}

	return process, nil
}

func Createprocess(ctx context.Context, process models.Process) (models.Process, error) {
	err := DB.WithContext(ctx).Create(&process).Error

	if err != nil {
		return models.Process{}, err
	}

	var newProcess models.Process
	err = DB.WithContext(ctx).Preload("Arrives").Where("id = ?", process.ID).First(&newProcess).Error
	if err != nil {
		return models.Process{}, err
	}

	jsonResponse, err := json.Marshal(newProcess)
	if err != nil {
		fmt.Println("Gagal mengonversi response ke JSON:", err)
		return models.Process{}, err
	}

	// Menyimpan JSON ke file
	err = ioutil.WriteFile("response.csv", jsonResponse, 0644)
	if err != nil {
		fmt.Println("Gagal menyimpan JSON ke file:", err)
		return models.Process{}, err
	}

	return newProcess, nil
}

func Deleteprocess(ctx context.Context, id int) error {
	var process models.Process

	result := DB.WithContext(ctx).Where("id = ?", id).Delete(&process)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func Updateprocess(ctx context.Context, id int, process models.Process) (models.Process, error) {
	result := DB.WithContext(ctx).Model(&models.Process{}).Where("id = ?", id).Updates(process)
	if result.Error != nil {
		return models.Process{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.Process{}, ErrIDNotFound
	}

	var newProcess models.Process
	err := DB.WithContext(ctx).Where("id = ?", process.ID).First(&newProcess).Error
	if err != nil {
		return models.Process{}, err
	}

	return newProcess, nil
}

func DecreaseStock(ctx context.Context, id int, count int) error {
	result := DB.WithContext(ctx).Model(&models.Process{}).Where(id).Update("count", gorm.Expr("count - ?", count))
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}
