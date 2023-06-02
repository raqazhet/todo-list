package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"todolist"
	"todolist/pkg/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	repo repository.Authorization
}

const (
	salt      = "ejbjdewblfwekdskjdf'awsl"
	signInKey = "#sgagsvas#$hsagdasl#sa$jwdsdhks(*&asgdhsg"
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user todolist.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	password = generatePasswordHash(password)
	user, err := s.repo.GetUser(username, password)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&tokenClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			UserId: user.Id,
		})
	text, err := token.SignedString([]byte(signInKey))
	if err != nil {
		logrus.Printf("error jwt token: %v", err)
		return "", err
	}
	return text, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
