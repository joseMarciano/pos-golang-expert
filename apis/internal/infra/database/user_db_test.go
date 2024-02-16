package database

import (
	"github.com/joseMarciano/pos-golang-expert/apis/internal/entity/user"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{}) // mock gorm.DB
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&user.User{})
	u, _ := user.NewUser("John", "john@example.com", "123456")
	userDB := NewUserDB(db)

	err = userDB.Create(u)
	assert.Nil(t, err)

	var userFound user.User

	err = db.First(&userFound, "email = ?", u.Email).Error
	assert.Nil(t, err)
	assert.Equal(t, u.ID, userFound.ID)
	assert.Equal(t, u.Name, userFound.Name)
	assert.Equal(t, u.Email, userFound.Email)
	assert.NotNil(t, u.Password)
	// assert user was inserted
}

//func TestNewUserDB(t *testing.T) {
//  db := // mock gorm.DB
//
//  userDB := NewUserDB(db)
//
//  if userDB.DB != db {
//    t.Errorf("DB is not set correctly")
//  }
//}
//
//func TestFindByEmail(t *testing.T) {
//	db := // mock gorm.DB
//
//	email := "test@test.com"
//	user := // sample user
//
//	userDB := NewUserDB(db)
//
//	// insert user
//
//	foundUser, err := userDB.FindByEmail(email)
//
//	if err != nil {
//		t.Errorf("FindByEmail returned error: %v", err)
//	}
//
//	if foundUser.Email != user.Email {
//		t.Errorf("Found incorrect user")
//	}
//}
