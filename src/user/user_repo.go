package user

import (
	userattributes "go-url-shortener-api/src/user/atrributes"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

type UserRepo interface {
	CreateUser(user *userattributes.CreateUserDTO) (*userattributes.UserDTO, error)
	GetUserByEmail(email string) (*userattributes.UserCredentialsDTO, error)
	GetUserByID(id string) (*userattributes.UserDTO, error)
	GetCountOfUsersByEmail(email string) (int64, error)
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		DB: db,
	}
}

func (repo *userRepo) CreateUser(user *userattributes.CreateUserDTO) (*userattributes.UserDTO, error) {
	newUser := User{
		ID:        uuid.New(),
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	result := repo.DB.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return toUserDTO(&newUser), nil
}

func (repo *userRepo) GetUserByEmail(email string) (*userattributes.UserCredentialsDTO, error) {
	var user User
	result := repo.DB.Select("id", "email", "password").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCredentialsDTO(&user), nil
}

func (repo *userRepo) GetUserByID(id string) (*userattributes.UserDTO, error) {
	var user User
	result := repo.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return toUserDTO(&user), nil
}

func (repo *userRepo) GetCountOfUsersByEmail(email string) (int64, error) {
	var count int64
	result := repo.DB.Model(&User{}).Where("email = ?", email).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func toUserDTO(user *User) *userattributes.UserDTO {
	return &userattributes.UserDTO{
		ID:        user.ID.String(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

func toCredentialsDTO(user *User) *userattributes.UserCredentialsDTO {
	return &userattributes.UserCredentialsDTO{
		ID:       user.ID.String(),
		Email:    user.Email,
		Password: user.Password,
	}
}
