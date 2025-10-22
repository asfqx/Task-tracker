package storage

import (
	"encoding/json"
	"os"
	"task-cli/internal/tasks"
)

type FileStorage struct {
	Name string
}

func (fs *FileStorage) Load() ([]tasks.Task, error) {
	if _, err := os.Stat(fs.Name); os.IsNotExist(err) {
		var emptyTasks []tasks.Task
		data, _ := json.Marshal(emptyTasks)
		if err := os.WriteFile(fs.Name, data, 0644); err != nil {
			return nil, err
		}
		return emptyTasks, nil
	}

	data, err := os.ReadFile(fs.Name)

	if err != nil {
		return []tasks.Task{}, err
	}

	if len(data) == 0 {
		return []tasks.Task{}, nil
	}

	var tasksList []tasks.Task
	err = json.Unmarshal(data, &tasksList)
	return tasksList, err
}

func (fs *FileStorage) Save(task []tasks.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}
	err = os.WriteFile(fs.Name, data, 0777)
	return err
}
