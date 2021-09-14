//+build wireinject

package pkg

import (
	"GoSOLID/pkg/drivers"
	"GoSOLID/pkg/services/repositories"
	"github.com/google/wire"
)

func InitializeJsonListRepository() *repositories.JsonListRepository {
	wire.Build(wire.Value("../data/list.json"), drivers.NewJsonDriver, repositories.NewJsonListRepository)
	return nil
}

func InitializeSqliteListRepository() *repositories.SqliteListRepository {
	wire.Build(repositories.NewSqliteListRepository)
	return nil
}

func InitializeListRepository() repositories.ListRepository {
	wire.Build(repositories.NewSqliteListRepository, wire.Bind(new(repositories.ListRepository), new(*repositories.SqliteListRepository)))
	return nil
}
