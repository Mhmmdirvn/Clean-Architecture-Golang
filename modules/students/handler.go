package students

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Usecase UseCase
}

func (handler Handler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.Usecase.CreateStudent(student)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Create New Data Success"))
}

func (handler Handler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	students, err := handler.Usecase.GetAllStudents()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func (handler Handler) GetStudentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	student, err := handler.Usecase.GetStudentById(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func (handler Handler) UpdateStudentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = handler.Usecase.UpdateStudentById(id, student)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Update Success"))
}

func (handler Handler) DeleteStudentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	err := handler.Usecase.DeleteStudentById(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Delete Success"))
}