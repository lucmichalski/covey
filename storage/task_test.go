package storage

import (
	"github.com/chabad360/covey/models"
	"reflect"
	"testing"
)

var task = &models.Task{
	ID:       "3778ffc302b6920c2589795ed6a7cad067eb8f8cb31b079725d0a20bfe6c3b6e",
	State:    models.StateQueued,
	Plugin:   "test",
	Details:  map[string]string{"test": "test"},
	ExitCode: 258,
}

func TestAddTask(t *testing.T) {
	//revive:disable:line-length-limit
	var tests = []struct {
		name string
		id   string
		want *models.Task
	}{
		{"success", task.ID, task},
		{"fail", "3", &models.Task{}},
	}
	//revive:enable:line-length-limit

	testError := AddTask(task)

	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			var got models.Task
			if DB.Where("id = ?", tt.id).First(&got); reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTask() = %v, want %v, error: %v", got, tt.want, testError)
			}
		})
	}
}

func TestSaveTask(t *testing.T) {
	//revive:disable:line-length-limit
	var tests = []struct {
		name   string
		id     string
		update TaskInfo
		want   []string
		want2  bool
		want3  models.TaskState
	}{
		{"start", task.ID, TaskInfo{nil, 257, task.ID}, nil, false, models.StateRunning},
		{"log", task.ID, TaskInfo{[]string{"hello"}, 257, task.ID}, []string{"hello"}, false, models.StateRunning},
		{"logError", task.ID, TaskInfo{[]string{"world"}, 1, task.ID}, []string{"hello", "world"}, false, models.StateError},
		{"success", task.ID, TaskInfo{nil, 0, task.ID}, []string{"hello", "world"}, false, models.StateDone},
		{"fail", "3", TaskInfo{}, nil, true, 0},
	}
	//revive:enable:line-length-limit

	for _, tt := range tests {
		testName := tt.name
		t.Run(testName, func(t *testing.T) {
			err := SaveTask(&tt.update)
			if (err != nil) != tt.want2 {
				t.Errorf("SaveTask(): error %v", err)
			}

			var got models.Task
			DB.Where("id = ?", tt.id).First(&got)

			if got.ExitCode != tt.update.ExitCode {
				t.Errorf("SaveTask(): ExitCode = %d, want %d", got.ExitCode, tt.update.ExitCode)
			}
			if reflect.DeepEqual(got.Log, tt.want) {
				t.Errorf("SaveTask(): Log = %v, want %v", got.Log, tt.want)
			}
			if got.State != tt.want3 {
				t.Errorf("SaveTask(): State = %v, want %v", got.State, tt.want3)
			}
		})
	}
}

func TestGetTask(t *testing.T) {
	//revive:disable:line-length-limit
	var tests = []struct {
		name string
		id   string
		want models.Task
	}{
		{"success", task.ID, *task},
		{"fail", "3", models.Task{}},
	}
	//revive:enable:line-length-limit

	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			if got, err := GetTask(tt.id); reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTask() = %v, want %v, error: %v", got, tt.want, err)
			}
		})
	}
}

//func TestMain(m *testing.M) {
//	pool, resource, pdb, err := test.Boilerplate()
//	db = pdb
//	DB = pdb
//	if err != nil {
//		log.Fatalf("Could not setup DB connection: %s", err)
//	}
//
//	err = db.AutoMigrate(&models.Task{})
//	if err != nil {
//		log.Fatalf("error preping the database: %s", err)
//	}
//
//	db.Create(task)
//
//	code := m.Run()
//
//	// You can't defer this because os.Exit doesn't care for defer
//	if err := pool.Purge(resource); err != nil {
//		log.Fatalf("Could not purge resource: %s", err)
//	}
//
//	os.Exit(code)
//}
