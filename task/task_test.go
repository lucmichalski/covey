package task

import (
	"testing"
	"time"

	"github.com/chabad360/covey/task/types"
)

var task1 = &types.Task{
	ID:     "2778ffc302b6920c2589795ed6a7cad067eb8f8cb31b079725d0a20bfe6c3b6c",
	State:  types.StateRunning,
	Plugin: "test",
	Node:   "test",
	Time:   time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
	Details: struct {
		Test string `json:"Test"`
	}{"test"},
}

func TestGetTask(t *testing.T) {
	type args struct {
		identifier string
	}
	tests := []struct {
		name  string
		args  args
		want  types.ITask
		want1 bool
	}{
		{"db", args{"3778ffc302b6920c2589795ed6a7cad067eb8f8cb31b079725d0a20bfe6c3b6e"}, task, true},
		{"noDB", args{"31b079725d0a20bfe6c3b6e"}, nil, false},
		{"map", args{"2778ffc302b6920c2589795ed6a7cad067eb8f8cb31b079725d0a20bfe6c3b6c"}, task1, true},
		{"mapShort", args{"2778ffc302b6920c"}, task1, true},
	}

	tasks[task1.GetID()] = task1
	tasksShort[task1.GetIDShort()] = task1.GetID()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetTask(tt.args.identifier)
			if got1 != tt.want1 && got.GetID() != tt.want.GetID() {
				t.Errorf("GetTask() got = %v, want %v, got1 = %v, want %v", got, tt.want, got1, tt.want1)
			}
		})
	}
}
