package user

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

type Repository interface {
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}
