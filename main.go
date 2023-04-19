package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Clean-Architecture/modules/class"
	"Clean-Architecture/modules/login"
	"Clean-Architecture/modules/students"
)

func main() {
	// Connection to database student
	db, errr := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_student"), &gorm.Config{})
	if errr != nil {
		panic("failed connect to database")
	}

	db.AutoMigrate(students.Student{}, class.Class{})

	studentRepo := students.Repository{DB: db}
	studentUseCase := students.UseCase{Repo: studentRepo}
	StudentHandler := students.Handler{Usecase: studentUseCase}

	classRepo := class.Repository{DB: db}
	classUseCase := class.UseCase{Repo: classRepo}
	classHandler := class.Handler{Usecase: classUseCase}


	// Connection to database users
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_users"), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}

	loginRepo := login.Repository{DB: db}
	loginUseCase := login.UseCase{Repo: loginRepo}
	loginHandler := login.Handler{Usecase: loginUseCase}

	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler.Login).Methods("POST")

	// Handler students
	r.HandleFunc("/students", MiddlewareJWTAuthorization(StudentHandler.CreateStudent)).Methods("POST")
	r.HandleFunc("/students", MiddlewareJWTAuthorization(StudentHandler.GetAllStudents)).Methods("GET")
	r.HandleFunc("/students/{id}", MiddlewareJWTAuthorization(StudentHandler.GetStudentById)).Methods("GET")
	r.HandleFunc("/students/{id}", MiddlewareJWTAuthorization(StudentHandler.UpdateStudentById)).Methods("PUT")
	r.HandleFunc("/students/{id}", MiddlewareJWTAuthorization(StudentHandler.DeleteStudentById)).Methods("DELETE")

	// Handler Class
	r.HandleFunc("/classes", MiddlewareJWTAuthorization(classHandler.CreateClass)).Methods("POST")
	r.HandleFunc("/classes", MiddlewareJWTAuthorization(classHandler.GetAllClass)).Methods("GET")
	r.HandleFunc("/classes/{id}", MiddlewareJWTAuthorization(classHandler.GetClassById)).Methods("GET")
	r.HandleFunc("/classes/{id}", MiddlewareJWTAuthorization(classHandler.UpdateClassById)).Methods("PUT")
	r.HandleFunc("/classes/{id}", MiddlewareJWTAuthorization(classHandler.DeleteClassById)).Methods("DELETE")

	// Set Port
	fmt.Println("server starter at localhost:9000")
	http.ListenAndServe(":9000", r)
}
