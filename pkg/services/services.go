package services

import (
	"GoSOLID/pkg/domain"
	"GoSOLID/pkg/services/repositories"
	"errors"
)

const MaxUserList = 5

type ListService struct {
	repo repositories.ListRepository
}

func NewListService(repo repositories.ListRepository) *ListService {
	return &ListService{repo: repo}
}

func (l *ListService) ValidateList(list domain.List) bool {
	return list.IsValid()
}

func (l *ListService) CreateList(list *domain.List) (int, error) {
	if !list.IsValid() {
		return domain.BADREQUEST, errors.New("bad arguments")
	}

	if len(l.repo.FetchLists()) > MaxUserList {
		return domain.BADREQUEST, errors.New("Maximum number of lists exceeded")
	}
	err := l.repo.CreateList(list)
	if err != nil {
		return domain.ERROR, err
	}

	return domain.SUCCESS, nil
}

func (l *ListService) FetchLists() []domain.List {
	return l.repo.FetchLists()
}
