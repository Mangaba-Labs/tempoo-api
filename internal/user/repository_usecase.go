package user

import (
	"errors"
	"gorm.io/gorm"
)

// Repository concrete type
type Repository struct {
	DB *gorm.DB // this can be any gorm instance
}

func (r Repository) FindAll() (users *gorm.DB, err error) {
	users = r.DB.Find(&users)
	err = users.Error
	return
}

func (r Repository) FindOneByEmail(email string) (user User, err error) {
	result := r.DB.First(&user, "email = ?", email)
	err = result.Error
	return
}

func (r Repository) FindById(id int) (user User, err error) {
	result := r.DB.First(&user, "id = ?", id)
	err = result.Error
	return
}

func (r Repository) Delete(id int) (err error) {
	result := r.DB.Delete(User{}, "id = ?", id)
	err = result.Error
	return
}

func (r Repository) Create(user *User) error {
	result := r.DB.Create(user)
	err := result.Error
	rowsCount := result.RowsAffected
	if err != nil || rowsCount <= 0 {
		return errors.New("cannot create user")
	}
	return nil
}
