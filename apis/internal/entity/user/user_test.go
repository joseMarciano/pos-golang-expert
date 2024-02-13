package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John", "email@email.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmptyf(t, user.ID, "ID should not be empty")
	assert.NotEmptyf(t, user.Password, "Password should not be empty")
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, "email@email.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John", "email@email.com", "123456")
	assert.Nil(t, err)
	assert.Truef(t, user.ValidatePassword("123456"), "Password should be valid")
	assert.False(t, user.ValidatePassword("1234567"), "Password should be valid")
	assert.NotEqualf(t, "123456", user.Password, "Password must be encrypted")
}
