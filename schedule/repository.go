package schedule

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Schedule, error)
	FindAllByUser(UserID uint) ([]Schedule, error)
	FindByID(ID int) (Schedule, error)
	Create(schedule Schedule) (Schedule, error)
	Update(schedule Schedule) (Schedule, error)
	Delete(schedule Schedule) (Schedule, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db, // Assign the provided db to the struct field
	}
}

func (r *repository) FindAll() ([]Schedule, error) {
	var sch []Schedule
	err := r.db.Find(&sch).Error

	return sch, err
}

func (r *repository) FindByID(ID int) (Schedule, error) {
	var sch Schedule
	err := r.db.First(&sch, ID).Error
	return sch, err
}

func (r *repository) Create(sch Schedule) (Schedule, error) {
	err := r.db.Create(&sch).Error
	return sch, err
}

func (r *repository) Delete(sch Schedule) (Schedule, error) {
	err := r.db.Delete(&sch).Error
	return sch, err
}

func (r *repository) FindAllByUser(UserID uint) ([]Schedule, error) {
	var sch []Schedule
	err := r.db.Where("user_id = ?", UserID).Find(&sch).Error
	return sch, err
}
