package repository

import (
	"database/sql"
	"sayembara/entity/model"
	"sayembara/utils"
)

type JobRepository interface {
	Create(userId string, job model.Job) (string, error)
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
