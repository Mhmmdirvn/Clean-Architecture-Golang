package class

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Usecase UseCase
}

func (handler Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var class Class
	err := json.NewDecoder(r.Body).Decode(&class)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.Usecase.CreateClass(class)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Success Add New Data"))
}

func (handler Handler) GetAllClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	class, err := handler.Usecase.GetAllClass()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(&class)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func (handler Handler) GetClassById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	class, err := handler.Usecase.GetClassById(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(&class)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func (handler Handler) UpdateClassById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	var class Class
	err := json.NewDecoder(r.Body).Decode(&class)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = handler.Usecase.UpdateClassById(id, class)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Update Success"))
}

func (handler Handler) DeleteClassById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	err := handler.Usecase.DeleteClassById(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Delete Success"))
}