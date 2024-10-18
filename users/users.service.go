package users

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) FindAll() []string {
	return []string{"user1", "user2"}
}
