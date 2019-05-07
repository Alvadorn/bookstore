package handler

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

func Handler(group *routing.RouteGroup) {
	group.Get("", listBooks)
}

func listBooks(c *routing.Context) error {

	return nil
}
