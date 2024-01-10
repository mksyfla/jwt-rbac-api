package mapping

import (
	"sayembara/entity/model"
	"sayembara/entity/response"
)

func JobsMap(job model.JobUser) response.GetJobs {
	return response.GetJobs{
		Id:          job.Id,
		Title:       job.Title,
		Description: job.Description,
		Image:       job.Image,
		Name:        job.Username,
	}
}
