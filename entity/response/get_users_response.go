package response

type GetUsers struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Banner   string `json:"banner"`
	Profile  string `json:"profile"`
}
