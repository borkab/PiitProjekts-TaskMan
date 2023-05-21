package taskMan

import (
	"context"
	"math/rand"
	"strconv"
	"time"
)

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ID          string `json:"id"`
	Info
}

type Info struct {
	MadeDay  time.Time `json:"madeDay"`
	UpdateAt time.Time `json:"updateAt"`
}

type TaskMemoRepo struct {
	//tasks is a map of a *Task, for easier manipulating as search and delete by ID(key)
	tasks map[string]*Task
}

type TaskRepository interface {
	//TaskRepository implements all these methods:
	Create(context.Context, *Task) error
	Update(context.Context, *Task) error
	FindByID(ctx context.Context, ID string) (_ Task, found bool, _ error)
	Delete(ctx context.Context, ID string) error
}

func (tmr *TaskMemoRepo) Create(ctx context.Context, NewTask *Task) error {
	if tmr.tasks == nil { //ha a TaskMemoRepoban levo tasks valtozom erteke nulla, akkor
		tmr.tasks = make(map[string]*Task) //hozd letre ezt a mappet
	}

	NewTask.ID = strconv.Itoa(rand.Int()) //a newTask valtozom ID fieldjenek adok egy random szamot erteknek, amit atalakitok stringge

	tmr.tasks[NewTask.ID] = NewTask //a TaskMemoRepo-ban levo most generalt ID kulcsnak megadom a newTask valtozot ertekparnak
	return nil
}
