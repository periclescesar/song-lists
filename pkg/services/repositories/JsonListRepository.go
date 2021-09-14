package repositories

import (
	"GoSOLID/pkg/domain"
	"GoSOLID/pkg/drivers"
)

type JsonListRepository struct {
	jsonDriver drivers.JsonDriver
}

func NewJsonListRepository(jsonDriver drivers.JsonDriver) *JsonListRepository {
	return &JsonListRepository{jsonDriver: jsonDriver}
}

func (j *JsonListRepository) CreateList(list *domain.List) error {
	return j.jsonDriver.Write(list)
}

func (j *JsonListRepository) FetchLists() []domain.List {
	return []domain.List{
		{Name: "mock 1", Songs: []int{1, 2, 3}, Type: "custom"},
		{Name: "mock 2", Songs: []int{1, 2, 3}, Type: "custom"},
	}
}
