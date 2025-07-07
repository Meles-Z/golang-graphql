package persistence

import (
	"context"
	"fmt"

	"github.com/meles-z/golang-graphql/app/domain/repository"
	"github.com/meles-z/golang-graphql/app/models"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) Create(ctx context.Context, input models.UserInput) (*models.User, error) {
	user := &models.User{
		Name:  input.Name,
		Email: input.Email,
	}
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) Update(ctx context.Context, id string, input models.UserInput) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	user.Name = input.Name
	user.Email = input.Email
	if err := r.db.WithContext(ctx).Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id string) (bool, error) {
	user, err := r.GetByID(ctx, id)
	if err != nil {
		return false, fmt.Errorf("user not found")
	}

	if err := r.db.WithContext(ctx).Delete(user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *userRepositoryImpl) GetByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetAll(ctx context.Context, filter *models.UserFilter) ([]*models.User, error) {
	var users []*models.User
	query := r.db.WithContext(ctx).Model(&models.User{})

	if filter != nil {
		if filter.Name != nil {
			query = query.Where("name ILIKE ?", "%"+*filter.Name+"%")
		}
		if filter.Email != nil {
			query = query.Where("email ILIKE ?", "%"+*filter.Email+"%")
		}
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
