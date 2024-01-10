package service

import (
	"sayembara/entity/model"
	"sayembara/entity/request"
	"sayembara/repository"
)

type JobService interface {
	Create(userId string, job request.JobPostRequest) (string, error)
	Draft(userId string, job request.DraftPostRequest) (string, error)
	GetJobs() ([]model.JobUser, error)
}

type jobService struct {
	jobRepository repository.JobRepository
}

func NewJobService(jobRepository repository.JobRepository) *jobService {
	return &jobService{jobRepository}
}

func (s *jobService) Create(userId string, job request.JobPostRequest) (string, error) {
	jobMap := model.Job{
		Title:       job.Title,
		Description: job.Description,
		Deadline:    job.Deadline,
		Reward:      job.Reward,
		Tag:         job.Tag,
		Image:       job.Image,
		Draft:       false,
	}

	id, err := s.jobRepository.Create(userId, jobMap)

	return id, err
}

func (s *jobService) Draft(userId string, job request.DraftPostRequest) (string, error) {
	jobMap := model.Job{
		Title:       job.Title,
		Description: job.Description,
		Deadline:    job.Deadline,
		Reward:      job.Reward,
		Tag:         job.Tag,
		Image:       job.Image,
		Draft:       true,
	}

	id, err := s.jobRepository.Create(userId, jobMap)

	return id, err
}

func (s *jobService) GetJobs() ([]model.JobUser, error) {
	jobs, err := s.jobRepository.GetJobs()

	return jobs, err
}
