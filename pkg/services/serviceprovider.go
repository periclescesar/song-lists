package services

import (
	"GoSOLID/pkg/drivers"
	"GoSOLID/pkg/services/repositories"
	"github.com/sarulabs/di"
)

var Container di.Container

func init() {
	Container = CreateServiceProvider()
}

func CreateServiceProvider() di.Container {
	jsonListStore := "../data/list.json"
	builder, _ := di.NewBuilder()

	_ = builder.Add([]di.Def{
		{
			Name:  "json-driver",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return drivers.NewJsonDriver(jsonListStore), nil
			},
		},
		{
			Name:  "json-list-repository",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				jd := ctn.Get("json-driver").(*drivers.JsonDriver)
				return repositories.NewJsonListRepository(*jd), nil
			},
		},
		{
			Name:  "sqlite-list-repository",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return repositories.NewSqliteListRepository(), nil
			},
		},
	}...)

	return builder.Build()
}
