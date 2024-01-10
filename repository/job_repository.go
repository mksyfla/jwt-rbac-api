package repository

import (
	"database/sql"
	"sayembara/entity/model"
	"sayembara/utils"
)

type JobRepository interface {
	Create(userId string, job model.Job) (string, error)
	GetJobs() ([]model.JobUser, error)
}

type jobRepository struct {
	idGenerator utils.IdGenerator
	db          *sql.DB
}

func NewJobRepository(idGenerator utils.IdGenerator, db *sql.DB) *jobRepository {
	return &jobRepository{idGenerator, db}
}

func (r *jobRepository) Create(userId string, job model.Job) (string, error) {
	id := r.idGenerator()

	query := "INSERT INTO jobs(id, id_user, title, description, tags, deadline, reward, draft) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, id, userId, job.Title, job.Description, job.Tag, job.Deadline, job.Reward, job.Draft)
	if err != nil {
		return "", err
	}

	for _, img := range job.Image {
		fileId := r.idGenerator()

		query = "INSERT INTO job_files(id, id_job, file) VALUES (?, ?, ?)"
		_, err = r.db.Exec(query, fileId, id, img)
		if err != nil {
			return "", err
		}
	}

	return id, err
}

func (r *jobRepository) GetJobs() ([]model.JobUser, error) {
	query := `
	SELECT
		jobs.id, jobs.title, jobs.description, job_files.file, users.name AS username
	FROM
  	jobs
	INNER JOIN
  	users ON jobs.id_user = users.id
	LEFT JOIN 
  	job_files ON jobs.id = job_files.id_job
	WHERE
  	jobs.draft = 0
	`

	rows, err := r.db.Query(query)

	if err != nil {
		return []model.JobUser{}, err
	}

	jobs := []model.JobUser{}

	for rows.Next() {
		job := model.JobUser{}
		rows.Scan(
			&job.Id, &job.Title, &job.Description, &job.Image, &job.Username,
		)
		jobs = append(jobs, job)
	}

	return jobs, err
}
