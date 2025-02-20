package task

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/stretchr/testify/assert"
)

var (
	testClient *client.Client = nil
)

func getTestClient() *client.Client {
	if testClient != nil {
		return testClient
	}
	return client.NewClient(getDefaultConfig())
}

func getTestService() *TaskService {
	return NewTaskService(getTestClient())
}

//func getTestTask(name string) *task.Task {
//	return &task.Task{
//		Name:         name,
//		Type:         "tags.cleanup",
//		CurrentState: "currentState",
//	}
//}

func getDefaultConfig() client.Config {
	timeout := tools.GetEnv("NEXUS_TIMEOUT", 30).(int)
	return client.Config{
		Insecure: tools.GetEnv("NEXUS_INSECURE_SKIP_VERIFY", true).(bool),
		Password: tools.GetEnv("NEXUS_PASSWORD", "admin123").(string),
		URL:      tools.GetEnv("NEXUS_URL", "http://127.0.0.1:8081").(string),
		Username: tools.GetEnv("NEXUS_USERNAME", "admin").(string),
		Timeout:  &timeout,
	}
}

func TestFreezeAndReleaseTaskState(t *testing.T) {
	s := getTestService()

	tasks, _, err := s.ListTasks(nil, nil)
	if err != nil {
		assert.Failf(t, "fail to list task", err.Error())
		return
	}
	for _, task := range tasks {
		assert.NotEmpty(t, task.ID)
		assert.NotEmpty(t, task.Name)
		assert.NotEmpty(t, task.Type)
		assert.NotEmpty(t, task.CurrentState)
	}

	// test get task api
	if len(tasks) > 0 {
		task, err := s.GetTask(tasks[0].ID)
		if err != nil {
			assert.Failf(t, "fail to run task", err.Error())
			return
		}
		assert.NotEmpty(t, task.ID)
		assert.NotEmpty(t, task.Name)
		assert.NotEmpty(t, task.Type)
		assert.NotEmpty(t, task.CurrentState)
	}

	if len(tasks) > 0 {
		err := s.RunTask(tasks[0].ID)
		if err != nil {
			assert.Failf(t, "fail to run task", err.Error())
			return
		}
		task, err := s.GetTask(tasks[0].ID)
		if err != nil {
			assert.Failf(t, "fail to get task", err.Error())
			return
		}
		assert.NotEmpty(t, task.LastRunResult)
		err = s.StopTask(tasks[0].ID)
		if err != nil {
			assert.Failf(t, "fail to stop task", err.Error())
		}
	}
}
