//go:generate go run ../../cmd/providergen/gen.go .

package provider

import (
	"github.com/sarulabs/dingo/v4"
)

type Provider struct {
	dingo.BaseProvider
}

const JsonListStore = "../data/list.json"

func (p *Provider) Load() error {
	if err := p.AddDefSlice(DriversDef); err != nil {
		return err
	}

	if err := p.AddDefSlice(RepositoriesDef); err != nil {
		return err
	}

	return nil
}
