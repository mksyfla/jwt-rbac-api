package response

type GetJobs struct {
	Id          string   `json:"job_id"`
	Title       string   `json:"job_title"`
	Description string   `json:"job_description"`
	Image       []string `json:"job_image"`
	Name        string   `json:"user_name"`
}
