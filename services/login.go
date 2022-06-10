package services

type LoginService interface {
	Login(username, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "hiro",
		authorizedPassword: "1234",
	}
}

func (s *loginService) Login(username, password string) bool {
	return s.authorizedUsername == username && s.authorizedPassword == password
}
