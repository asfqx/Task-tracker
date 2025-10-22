package tasks

import "time"

type Tracker struct {
	store Storage
}

func NewTracker(store Storage) *Tracker {
	return &Tracker{store: store}
}

func (tracker *Tracker) Add(description string) error {
	tasks, err := tracker.store.Load()
	if err != nil {
		return err
	}
	task := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format("02-01-06 15:04"),
		UpdatedAt:   time.Now().Format("02-01-06 15:04"),
	}
	tasks = append(tasks, task)
	err = tracker.store.Save(tasks)
	if err != nil {
		return err
	}
	return nil
}

func (tracker *Tracker) Delete(id int) error {
	tasks, err := tracker.store.Load()
	if err != nil {
		return err
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	err = tracker.store.Save(tasks)
	if err != nil {
		return err
	}
	return nil
}

func (tracker *Tracker) List(category string) ([]Task, error) {
	tasks, err := tracker.store.Load()
	if err != nil {
		return []Task{}, err
	}
	if category == "" {
		return tasks, err
	}
	var result []Task
	for _, task := range tasks {
		if task.Status == category {
			result = append(result, task)
		}
	}
	return result, nil
}

func (tracker *Tracker) Update(id int, description string) error {
	tasks, err := tracker.store.Load()
	if err != nil {
		return err
	}

	if description == "in-progress" || description == "done" {
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Status = description
				tasks[i].UpdatedAt = time.Now().Format("06-01-02 15:04:05")
				break
			}
		}
	} else {
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Description = description
				tasks[i].UpdatedAt = time.Now().Format("06-01-02 15:04:05")
				break
			}
		}
	}

	err = tracker.store.Save(tasks)
	if err != nil {
		return err
	}

	return nil
}
