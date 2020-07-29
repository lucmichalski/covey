package task

import (
	"container/list"
	json "github.com/json-iterator/go"
)

// AgentTask contains the information about a task that is send to an agent.
type agentTask struct {
	Command string `json:"command"`
	ID      string `json:"id"`
}

// List implements List type as well as the json.Marshaler interface.
type List struct{ list.List }

// MarshalJSON implements the json.Marshaler interface.
func (l *List) MarshalJSON() ([]byte, error) {
	m := make(map[int]agentTask)
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		m[i] = e.Value.(agentTask)
		i++
	}
	return json.Marshal(m)
}

// Plugin defines the interface for Task module plugins.
type Plugin interface {
	// GetCommand returns the command to run the server.
	GetCommand(taskJSON []byte) (string, error)
}
