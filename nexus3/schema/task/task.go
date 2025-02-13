package task

type Task struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	Message       *string `json:"message"`
	CurrentState  string  `json:"currentState`
	LastRunResult *string `json:"lastRunResult`
	NextRun       *string `json:"nextRun"`
	LastRun       *string `json:"lastRun"`
}
