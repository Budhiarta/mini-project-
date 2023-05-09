package database

import (
	"context"
	"miniProject/models"
	"miniProject/utils"
)

func GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	err := DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(ctx context.Context, id int) (models.User, error) {
	var user models.User

	err := DB.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CreateUser(ctx context.Context, user models.User) (models.User, error) {
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return user, err
	}

	user.Password = hash

	err = DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUser(ctx context.Context, id int) error {
	var user models.User

	result := DB.WithContext(ctx).Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func UpdateUser(ctx context.Context, id int, user models.User) (models.User, error) {
	result := DB.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.User{}, ErrIDNotFound
	}

	return user, nil
}

func LoginUser(ctx context.Context, requestUser models.User) (models.User, error) {
	var user models.User

	err := DB.WithContext(ctx).Where("email = ?", requestUser.Email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
