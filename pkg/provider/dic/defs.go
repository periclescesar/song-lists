package dic

import (
	"errors"

	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"

	drivers "GoSOLID/pkg/drivers"
	repositories "GoSOLID/pkg/services/repositories"
)

func getDiDefs(provider dingo.Provider) []di.Def {
	return []di.Def{
		{
			Name:  "json-driver",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("json-driver")
				if err != nil {
					var eo *drivers.JsonDriver
					return eo, err
				}
				pi0, ok := d.Params["0"]
				if !ok {
					var eo *drivers.JsonDriver
					return eo, errors.New("could not find parameter 0")
				}
				p0, ok := pi0.(string)
				if !ok {
					var eo *drivers.JsonDriver
					return eo, errors.New("could not cast parameter 0 to string")
				}
				b, ok := d.Build.(func(string) (*drivers.JsonDriver, error))
				if !ok {
					var eo *drivers.JsonDriver
					return eo, errors.New("could not cast build function to func(string) (*drivers.JsonDriver, error)")
				}
				return b(p0)
			},
			Unshared: false,
		},
		{
			Name:  "json-list-repository",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("json-list-repository")
				if err != nil {
					var eo *repositories.JsonListRepositories
					return eo, err
				}
				pi0, err := ctn.SafeGet("json-driver")
				if err != nil {
					var eo *repositories.JsonListRepositories
					return eo, err
				}
				p0, ok := pi0.(*drivers.JsonDriver)
				if !ok {
					var eo *repositories.JsonListRepositories
					return eo, errors.New("could not cast parameter 0 to *drivers.JsonDriver")
				}
				b, ok := d.Build.(func(*drivers.JsonDriver) (*repositories.JsonListRepositories, error))
				if !ok {
					var eo *repositories.JsonListRepositories
					return eo, errors.New("could not cast build function to func(*drivers.JsonDriver) (*repositories.JsonListRepositories, error)")
				}
				return b(p0)
			},
			Unshared: false,
		},
		{
			Name:  "list-repository",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("list-repository")
				if err != nil {
					var eo repositories.ListRepository
					return eo, err
				}
				pi0, err := ctn.SafeGet("sqlite-list-repository")
				if err != nil {
					var eo repositories.ListRepository
					return eo, err
				}
				p0, ok := pi0.(repositories.ListRepository)
				if !ok {
					var eo repositories.ListRepository
					return eo, errors.New("could not cast parameter 0 to repositories.ListRepository")
				}
				b, ok := d.Build.(func(repositories.ListRepository) (repositories.ListRepository, error))
				if !ok {
					var eo repositories.ListRepository
					return eo, errors.New("could not cast build function to func(repositories.ListRepository) (repositories.ListRepository, error)")
				}
				return b(p0)
			},
			Unshared: false,
		},
		{
			Name:  "sqlite-list-repository",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("sqlite-list-repository")
				if err != nil {
					var eo *repositories.SqliteListRepository
					return eo, err
				}
				b, ok := d.Build.(func() (*repositories.SqliteListRepository, error))
				if !ok {
					var eo *repositories.SqliteListRepository
					return eo, errors.New("could not cast build function to func() (*repositories.SqliteListRepository, error)")
				}
				return b()
			},
			Unshared: false,
		},
	}
}
