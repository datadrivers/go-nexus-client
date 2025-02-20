package task

type FrequencyXO struct {
	Schedule       string        `json:"schedule"`
	StartDate      int           `json:"startDate,omitempty"`
	TimeZoneOffset string        `json:"timeZoneOffset,omitempty"`
	RecurringDays  []interface{} `json:"recurringDays,omitempty"`
	CronExpression string        `json:"cronExpression,omitempty"`
}

type Task struct {
	ID           string      `json:"id"`
	Type         string      `json:"type"`
	Name         string      `json:"name"`
	Message      string      `json:"message,omitempty"`
	CurrentState string      `json:"currentState,omitempty"`
	Frequency    FrequencyXO `json:"frequency,omitempty"`
	NextRun      string      `json:"nextRun,omitempty"`
	LastRun      string      `json:"lastRun,omitempty"`
}

type TaskCreateStruct struct {
	Type                  string                 `json:"type,omitempty"`
	Name                  string                 `json:"name"`
	Enabled               bool                   `json:"enabled"`
	AlertEmail            string                 `json:"alertEmail,omitempty"`
	NotificationCondition string                 `json:"notificationCondition"`
	Frequency             FrequencyXO            `json:"frequency,omitempty"`
	Message               string                 `json:"message,omitempty"`
	Properties            map[string]interface{} `json:"properties,omitempty"`
}
