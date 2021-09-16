//go:generate go run ../../cmd/providergenerate/gen.go .

package provider

import (
	"GoSOLID/pkg/drivers"
	"github.com/sarulabs/dingo/v4"
)

var ProviderDrivers = []dingo.Def{
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
