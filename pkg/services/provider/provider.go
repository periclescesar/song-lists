package provider

import (
	"GoSOLID/pkg/drivers"
	"GoSOLID/pkg/services/repositories"
	"github.com/sarulabs/dingo/v4"
)

type Provider struct {
	dingo.BaseProvider
}

const JsonListStore = "../data/list.json"

func (p *Provider) Load() error {
	services := []dingo.Def{
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
		{
			Name: "json-driver",
			Build: func(filepath string) (*drivers.JsonDriver, error) {
				return drivers.NewJsonDriver(filepath), nil
			},
			Params: dingo.Params{
				"0": JsonListStore,
			},
		},
	}

	if err := p.AddDefSlice(services); err != nil {
		return err
	}

	return nil
}
