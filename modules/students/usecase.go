package students

type UseCase struct {
	Repo Repository
}

func (usecase UseCase) GetAllStudents() ([]Student, error) {
	students, err := usecase.Repo.GetAllStudents()
	return students, err
}

func (usecase UseCase) GetStudentById(id string) (Student, error) {
	student, err := usecase.Repo.GetStudentById(id)
	return student, err
}

func (usecase UseCase) CreateStudent(student Student) error {
	err := usecase.Repo.CreateStudent(student)
	return err
}

func (usecase UseCase) UpdateStudentById(id string, student Student) error {
	err := usecase.Repo.UpdateStudentById(id, student)
	return err
}

func (usecase UseCase) DeleteStudentById(id string) error {
	err := usecase.Repo.DeleteStudentById(id)
	return err
}