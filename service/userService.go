package service

import (
	"log"

	"github.com/Sutrisno-14/skill-test-Nusatech/entity"
	"github.com/Sutrisno-14/skill-test-Nusatech/repository"
	"github.com/Sutrisno-14/skill-test-Nusatech/transport"
	"github.com/mashingan/smapping"
)

//UserService 
type UserService interface {
	Update(user transport.UserUpdate) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user transport.UserUpdate) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}