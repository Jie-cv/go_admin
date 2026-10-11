package service

import "github.com/Jie-cv/go_admin/contant"

func NewQueryRequest() *QueryRequest {
	return &QueryRequest{
		Page:    1,
		Size:    contant.DefaultSize,
		Keyword: "",
	}
}
