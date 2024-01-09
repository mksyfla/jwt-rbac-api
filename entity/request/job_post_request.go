package request

type JobPostRequest struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Deadline    int64    `json:"deadline" binding:"required"`
	Reward      string   `json:"reward"`
	Tag         string   `json:"tag" binding:"required"`
	Image       []string `json:"image" binding:"required"`
}
