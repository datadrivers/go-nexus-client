package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema"
	"net/http"
	"net/url"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/task"
)

const (
	taskAPIEndpoint = client.BasePath + "v1/tasks"
)

var (
	ErrTaskNotRunning = errors.New("task is not currently running")
)

type TaskService client.Service

func NewTaskService(c *client.Client) *TaskService {
	return &TaskService{Client: c}
}

func (s *TaskService) ListTasks(taskType *string, continuationToken *string) ([]task.Task, *string, error) {
	q := url.Values{}
	if taskType != nil {
		q.Set("type", *taskType)
	}
	if continuationToken != nil {
		q.Set("continuationToken", *continuationToken)
	}

	body, resp, err := s.Client.Get(fmt.Sprintf("%s?%s", taskAPIEndpoint, q.Encode()), nil)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("could not list task: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var result schema.PaginationResult[task.Task]
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, nil, fmt.Errorf("could not unmarshal tasks: %v", err)
	}

	return result.Items, result.ContinuationToken, nil
}

func (s *TaskService) GetTask(id string) (*task.Task, error) {
	body, resp, err := s.Client.Get(fmt.Sprintf("%s/%s", taskAPIEndpoint, id), nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		var t task.Task
		if err := json.Unmarshal(body, &t); err != nil {
			return nil, fmt.Errorf("could not unmarshal task: %v", err)
		}
		return &t, nil
	case http.StatusNotFound:
		return nil, nil
	default:
		return nil, fmt.Errorf("could not get task '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
}

func (s *TaskService) RunTask(id string) error {
	body, resp, err := s.Client.Post(fmt.Sprintf("%s/%s/run", taskAPIEndpoint, id), nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not run task '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *TaskService) StopTask(id string) error {
	body, resp, err := s.Client.Post(fmt.Sprintf("%s/%s/stop", taskAPIEndpoint, id), nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusConflict:
		return ErrTaskNotRunning
	default:
		return fmt.Errorf("could not stop task '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
}

func (s *TaskService) CreateTask(newTask *task.TaskCreateStruct) (*task.Task, error) {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(newTask)
	if err != nil {
		return nil, err
	}
	body, resp, err := s.Client.Post(taskAPIEndpoint, ioReader)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("could not create task: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var createdTask task.Task
	if err := json.Unmarshal(body, &createdTask); err != nil {
		return nil, fmt.Errorf("could not unmarshal created task: %v", err)
	}

	return &createdTask, nil
}

func (s *TaskService) UpdateTask(id string, updatedTask *task.TaskCreateStruct) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(updatedTask)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(fmt.Sprintf("%s/%s", taskAPIEndpoint, id), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update task '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *TaskService) DeleteTask(id string) error {
	body, resp, err := s.Client.Delete(fmt.Sprintf("%s/%s", taskAPIEndpoint, id))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete task '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}

	return nil
}
