package database

import (
	"github.com/joseMarciano/pos-golang-expert/apis/internal/entity/user"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{
		DB: db,
	}
}

func (u *UserDB) Create(user *user.User) error {
	return u.DB.Create(user).Error
}

func (u *UserDB) FindByEmail(email string) (*user.User, error) {
	var user user.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
