//go:generate go run ../../cmd/providergen/gen.go .

package provider

import (
	"GoSOLID/pkg/drivers"
	"GoSOLID/pkg/services/repositories"
	"github.com/sarulabs/dingo/v4"
)

var RepositoriesDef = []dingo.Def{
	{
		Name: "list-repository",
		Build: func(repo repositories.ListRepository) (repositories.ListRepository, error) {
			return repo, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("sqlite-list-repository"),
		},
	},
	{
		Name: "sqlite-list-repository",
		Build: func() (*repositories.SqliteListRepository, error) {
			return repositories.NewSqliteListRepository(), nil
		},
	},
	{
		Name: "json-list-repository",
		Build: func(jd *drivers.JsonDriver) (*repositories.JsonListRepositories, error) {
			return repositories.NewJsonListRepository(*jd), nil
		},
		Params: dingo.Params{
			"0": dingo.Service("json-driver"),
		},
	},
}
