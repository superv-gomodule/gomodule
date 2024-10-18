package users

type UserService interface {
	FindAll() []string
}

type userServiceImpl struct {
}

func NewUserService() UserService {
	return &userServiceImpl{}
}

func (s *userServiceImpl) FindAll() []string {
	return []string{"user1", "user2"}
}
