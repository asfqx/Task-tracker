package tasks

type Storage interface {
	Save([]Task) error
	Load() ([]Task, error)
}
