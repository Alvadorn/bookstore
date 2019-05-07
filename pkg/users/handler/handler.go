package handler

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

// Handler users
func Handler(group *routing.RouteGroup) {
	group.Post("", saveUser)
}

func saveUser(c *routing.Context) error {

	return nil
}
