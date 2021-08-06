package repositories

import (
	"GoSOLID/pkg/domain"
)

type ListRepository interface {
	CreateList(list *domain.List) error
	FetchLists() []domain.List
}
