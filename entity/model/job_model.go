package model

type Job struct {
	Id          string
	Title       string
	Description string
	Deadline    int64
	Reward      string
	Tag         string
	Image       []string
	Draft       bool
}

type JobUser struct {
	Job
	Username string
}
