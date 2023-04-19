package login

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UseCase struct {
	Repo Repository
}

func (usecase UseCase) Login(user User) (string ,error) {
	err := usecase.Repo.Login(user)

	claims := MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("MySecret"))
	return signedToken, err

}

