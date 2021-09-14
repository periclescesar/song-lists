//+build wireinject

package pkg

import (
	"GoSOLID/pkg/drivers"
	"GoSOLID/pkg/services/repositories"
	"github.com/google/wire"
)

var sqlRepo = wire.NewSet(repositories.NewSqliteListRepository)
var jsonRepo = wire.NewSet(wire.Value("../data/list.json"), drivers.NewJsonDriver, repositories.NewJsonListRepository)

func ProviderListRepository(draft bool) repositories.ListRepository {
	wire.Build(jsonRepo, wire.InterfaceValue(new(repositories.ListRepository), new(*repositories.JsonListRepository)))
	return nil
}

func InitializeListRepository(draft bool) repositories.ListRepository {
	wire.Build(ProviderListRepository)
	return nil
}
