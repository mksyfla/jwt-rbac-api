package mapping

import (
	"sayembara/entity/model"
	"sayembara/entity/response"
)

func UsersMap(user model.User) response.GetUsers {
	return response.GetUsers{
		Id:       user.Id,
		Name:     user.Name,
		Category: user.Category,
		Banner:   user.Banner,
		Profile:  user.Profile,
	}
}
