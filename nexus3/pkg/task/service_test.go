package task

import (
	"testing"
	"time"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/task"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/stretchr/testify/assert"
)

const dummyTask = "dummy"

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
}

func TestTaskService_CreateTask(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}
	s := getTestService()
	newTask := getTestTask()
	createdTask, err := s.CreateTask(newTask)
	if err != nil {
		assert.Failf(t, "fail to create task", err.Error())
		return
	}
	assert.NotNil(t, createdTask)
}

func getTestTask() *task.TaskCreateStruct {
	newTask := &task.TaskCreateStruct{
		Name:                  "test-task",
		Type:                  "tags.cleanup",
		Enabled:               true,
		AlertEmail:            "abc@acb.com",
		NotificationCondition: "FAILURE",
		Frequency: &task.FrequencyXO{
			Schedule:       "manual",
			StartDate:      int(time.Now().Unix()),
			TimeZoneOffset: "-08:00",
			CronExpression: "string",
		},
		Properties: map[string]interface{}{},
	}
	return newTask
}

func TestTaskService_UpdateTask(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}
	s := getTestService()
	newTask := getTestTask()
	createdTask, err := s.CreateTask(newTask)
	if err != nil {
		assert.Failf(t, "fail to create task", err.Error())
		return
	}
	assert.NotNil(t, createdTask)
	newTask.Type = ""
	newTask.Name = "test-task-updated"
	err = s.UpdateTask(createdTask.ID, newTask)
	if err != nil {
		assert.Failf(t, "fail to update task", err.Error())
		return
	}
	updatedTask, err := s.GetTask(createdTask.ID)
	if err != nil {
		assert.Failf(t, "fail to get task", err.Error())
		return
	}
	assert.NotNil(t, updatedTask)
	assert.Equal(t, newTask.Name, updatedTask.Name, "newTask.Name should be equal to updatedTask.Name")

	err = s.UpdateTask(dummyTask, newTask)
	assert.NotNil(t, err)
}

func TestTaskService_DeleteTaskTask(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}
	s := getTestService()
	newTask := getTestTask()
	createdTask, err := s.CreateTask(newTask)
	if err != nil {
		assert.Failf(t, "fail to create task", err.Error())
		return
	}
	assert.NotNil(t, createdTask)
	err = s.DeleteTask(createdTask.ID)
	removedTask, err := s.GetTask(createdTask.ID)
	if err != nil {
		assert.Failf(t, "fail to update task", err.Error())
		return
	}
	assert.Nil(t, removedTask)
	err = s.DeleteTask(dummyTask)
	assert.NotNil(t, err)
}

func TestTaskService_RunTask(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}
	s := getTestService()
	newTask := getTestTask()
	createdTask, err := s.CreateTask(newTask)
	if err != nil {
		assert.Failf(t, "fail to create task", err.Error())
		return
	}
	assert.NotNil(t, createdTask)
	err = s.RunTask(createdTask.ID)
	if err != nil {
		assert.Failf(t, "fail to run task", err.Error())
		return
	}
	err = s.RunTask(dummyTask)
	assert.NotNil(t, err)
}

func TestTaskService_StopTask(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}
	s := getTestService()
	newTask := getTestTask()
	createdTask, err := s.CreateTask(newTask)
	if err != nil {
		assert.Failf(t, "fail to create task", err.Error())
		return
	}
	assert.NotNil(t, createdTask)
	err = s.RunTask(createdTask.ID)
	if err != nil {
		assert.Failf(t, "fail to run task", err.Error())
		return
	}
	err = s.StopTask(createdTask.ID)
	if err != nil {
		assert.Failf(t, "fail to stop task", err.Error())
		return
	}
	err = s.StopTask(dummyTask)
	assert.NotNil(t, err)
}
