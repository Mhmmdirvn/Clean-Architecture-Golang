package class

type UseCase struct {
	Repo Repository
}

func (usecase UseCase) GetAllClass() ([]Class, error) {
	class, err := usecase.Repo.GetAllClass()
	return class, err
}

func (usecase UseCase) GetClassById(id string) (*Class, error) {
	class, err := usecase.Repo.GetClassById(id)
	return class, err
}

func (usecase UseCase) CreateClass(class Class) error {
	err := usecase.Repo.CreateClass(class)
	return err
}

func (usecase UseCase) UpdateClassById(id string, class Class) error {
	err := usecase.Repo.UpdateClassById(id, class)
	return err
}

func (usecase UseCase) DeleteClassById(id string) error {
	err := usecase.Repo.DeleteClassById(id)
	return err
}