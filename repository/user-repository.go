package repository

import (
	"github.com/drewjya/aice-server/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	ProfileUser(userId string) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}
func (db *userConnection) CreateUser(user *entity.User) (*entity.User, error) {
	user.Password = hashPassword([]byte(user.Password))
	errTrx := db.connection.Save(&user)
	if errTrx.Error != nil {
		return nil, errTrx.Error
	}
	return user, nil
}
func (db *userConnection) UpdateUser(user entity.User) entity.User {
	user.Password = hashPassword([]byte(user.Password))
	db.connection.Save(&user)
	return user
}
func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) ProfileUser(userId string) entity.User {
	var user entity.User
	db.connection.Find(&user, userId)
	return user
}

func hashPassword(pass []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		panic("failed to hash password")
	}
	return string(hash)
}
