package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("BWA")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}

	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedtToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedtToken, err
	}

	return signedtToken, nil
}
