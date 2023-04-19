package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	Usecase UseCase
}

func (handler Handler) Login(w http.ResponseWriter, r *http.Request) {
	var users User
	json.NewDecoder(r.Body).Decode(&users)
	
	token, err := handler.Usecase.Login(users)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Login failed"))
		} else {
			w.Write([]byte(err.Error()))
		}

		return


	}

	fmt.Println("token :", token)
	w.Write([]byte(token))

}