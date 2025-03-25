package test

import (
	"github.com/VadimGossip/extRoutingClientTester/internal/repository"
	"github.com/VadimGossip/extRoutingClientTester/internal/service/test/model"
)

type service struct {
	repoTest repository.TestRepository
	tasks    map[int64]*model.TestTask
}

func NewService(repoTest repository.TestRepository) *service {
	return &service{repoTest: repoTest, tasks: map[int64]*model.TestTask{}}
}
