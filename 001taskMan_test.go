package taskMan

import (
	"context"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	newTask := &Task{
		Name:        "Kitchen",
		Description: "CleanUp and Cooking",
		Status:      "Open",
		Info: Info{
			MadeDay:  time.Now(),
			UpdateAt: time.Now(),
		},
	}

	repo := TaskMemoRepo{}
	ctx := context.Background()

	got := repo.Create(ctx, newTask)
	t.Run("naked test", func(t *testing.T) {

		if got != nil {
			t.Fatal("couldn't create new task")
		}
	})
	t.Run("empty or not empty", func(t *testing.T) {
		if repo.tasks == nil {
			t.Fatal("couldn't create any tasks")
		}
	})
}
