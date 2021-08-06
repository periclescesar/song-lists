package repositories

import "GoSOLID/pkg/domain"

type SqliteListRepository struct{}

func NewSqliteListRepository() *SqliteListRepository {
	return &SqliteListRepository{}
}

func (j *SqliteListRepository) CreateList(list *domain.List) error {
	//sql := "INSERT INTO `lists` set name= :name"

	return nil
}

func (j *SqliteListRepository) FetchLists() []domain.List {
	return []domain.List{
		{Name: "mock 1", Songs: []int{1, 2, 3}, Type: "custom"},
		{Name: "mock 2", Songs: []int{1, 2, 3}, Type: "custom"},
	}
}
