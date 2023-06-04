package service

import (
	"crypto/sha1"
	"errors"
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

func NewAuthService(repo repository.Authorization) *AuthService {
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

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signInKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
