package data

import (
	"time"
)

// Deploy status constants
const (
	StatusPending = "pending"
	StatusSuccess = "success"
	StatusError   = "error"
	StatusFailure = "failure"
)

// Deploy represents a deployment
type Deployment struct {
	Url           string      `json:"url"`
	Id            int         `json:"id"`
	Sha           string      `json:"sha"`
	Ref           string      `json:"ref"` // Required always
	Task          string      `json:"task"`
	Payload       interface{} `json:"payload"`
	Environment   string      `json:"environment"`
	Description   string      `json:"description"`
	Creator       User        `json:"creator"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	StatusesUrl   string      `json:"statuses_url"`
	RepositoryUrl string      `json:"repository_url"`
	StatusId      int         `json:"status_id"`
}

// Status represents the current status of a deployment
type Status struct {
	Url           string    `json:"url"`
	Id            int       `json:"id"`
	State         string    `json:"state"` // Required always
	Creator       User      `json:"creator"`
	Description   string    `json:"description"`
	TargetUrl     string    `json:"target_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeploymentUrl string    `json:"deployment_url"`
	RepositoryUrl string    `json:"repository_url"`
	DeploymentId  int       `json:"deploy_id"`
}
