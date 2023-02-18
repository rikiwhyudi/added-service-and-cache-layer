package service

import (
	authcache "backend-api/cache"
	authdto "backend-api/dto/auth"
	"backend-api/models"
	"backend-api/pkg/bcrypt"
	jwtToken "backend-api/pkg/jwt"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	Register(request authdto.RegisterRequest) (*authdto.RegisterResponse, error)
	Login(request authdto.LoginRequest) (*authdto.LoginResponse, error)
	GetUserID(ID int) (*authdto.CheckAutResponse, error)
}

type authService struct {
	authCache authcache.AuthCache
}

func NewAuthService(authCache authcache.AuthCache) *authService {
	return &authService{authCache}
}

func (s *authService) Register(request authdto.RegisterRequest) (*authdto.RegisterResponse, error) {

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
		Role:     "customer",
	}

	data, err := s.authCache.Register(user)
	if err != nil {
		return nil, err
	}

	response := authdto.RegisterResponse{
		Name:  data.Name,
		Email: data.Email,
		Role:  data.Role,
	}

	return &response, nil
}

func (s *authService) Login(request authdto.LoginRequest) (*authdto.LoginResponse, error) {

	user, err := s.authCache.Login(request.Email)
	if err != nil {
		return nil, fmt.Errorf("email not registered!")
	}

	//check password for auth
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return nil, fmt.Errorf("wrong password!")
	}

	//generate token login
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hour expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("unauthorized")
		return nil, errGenerateToken

	}

	loginResponse := authdto.LoginResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
		Role:  user.Role,
	}

	return &loginResponse, nil
}

func (s *authService) GetUserID(ID int) (*authdto.CheckAutResponse, error) {

	//getUserById
	user, err := s.authCache.GetUserID(ID)
	if err != nil {
		return nil, err
	}

	response := authdto.CheckAutResponse{
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role,
	}

	return &response, nil
}
