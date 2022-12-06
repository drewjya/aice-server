package services

import (
	"log"

	"github.com/drewjya/aice-server/dto"
	"github.com/drewjya/aice-server/entity"
	"github.com/drewjya/aice-server/repository"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) (*entity.User, error)
	ProfileUser(userId string) entity.User
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepository: userRepo}
}

func (auth *authService) VerifyCredential(email string, password string) interface{} {
	res := auth.userRepository.VerifyCredential(email, password)

	log.Println(res)
	if v, ok := res.(entity.User); ok {
		comparePass := compareHashPassword([]byte(password), v.Password)
		if comparePass && v.Email == email {
			return res

		}
		return false
	}
	return false
}
func (auth *authService) CreateUser(user dto.RegisterDTO) (*entity.User, error) {
	userCreate := entity.User{}

	err := smapping.FillStruct(&userCreate, smapping.MapFields(user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res, err := auth.userRepository.CreateUser(&userCreate)
	if err != nil {
		return nil, err
	}
	return res, nil

}
func (auth *authService) ProfileUser(userId string) entity.User {
	res := auth.userRepository.ProfileUser(userId)
	return res
}

func compareHashPassword(password []byte, hashedPasword string) bool {
	byteHash := []byte(hashedPasword)
	err := bcrypt.CompareHashAndPassword(byteHash, password)

	return err == nil
}
