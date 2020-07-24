package job

import (
	"github.com/chabad360/covey/models"
	"github.com/chabad360/covey/test"
	"log"
	"os"
	"testing"

	"github.com/chabad360/covey/storage"
)

var j = &models.Job{
	Name:  "update",
	ID:    "3778ffc302b6920c2589795ed6a7cad067eb8f8cb31b079725d0a20bfe6c3b6e",
	Nodes: []string{"node1"},
	Tasks: map[string]models.JobTask{
		"update": {
			Plugin:  "shell",
			Details: map[string]string{"command": "sudo apt update && sudo apt upgrade -y"},
		},
	},
}

func TestGetJob(t *testing.T) {
	//revive:disable:line-length-limit
	var tests = []struct {
		id   string
		want bool
	}{
		{"3778ffc302b6920c2589795ed6a7cad067eb8f8cb31b079725d0a20bfe6c3b6e", true},
		{"3", false},
	}
	//revive:enable:line-length-limit

	for _, tt := range tests {
		testname := tt.id
		t.Run(testname, func(t *testing.T) {
			if _, ok := storage.GetJob(tt.id); ok != tt.want {
				t.Errorf("GetJob() = %v, want %v", ok, tt.want)
			}
		})
	}
}

func TestGetJobWithTasks(t *testing.T) {
	//revive:disable:line-length-limit
	var tests = []struct {
		id   string
		want bool
	}{
		{"3778ffc302b6920c2589795ed6a7cad067eb8f8cb31b079725d0a20bfe6c3b6e", true},
		{"3", false},
	}
	//revive:enable:line-length-limit

	for _, tt := range tests {
		testname := tt.id
		t.Run(testname, func(t *testing.T) {
			if _, ok := storage.GetJob(tt.id); ok != tt.want {
				t.Errorf("GetJobWithTasks() = %v, want %v", ok, tt.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	pool, resource, pdb, err := test.Boilerplate()
	storage.DB = pdb
	if err != nil {
		log.Fatalf("Could not setup DB connection: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
