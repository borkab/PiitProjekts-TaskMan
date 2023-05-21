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

	if got != nil {
		t.Fatal("couldn't create new task")
	}
}
