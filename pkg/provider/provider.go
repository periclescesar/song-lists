//go:generate go run ../../cmd/providergenerate/gen.go .

package provider

import (
	"github.com/sarulabs/dingo/v4"
)

type Provider struct {
	dingo.BaseProvider
}

const JsonListStore = "../data/list.json"

func (p *Provider) Load() error {
	if err := p.AddDefSlice(ProviderDrivers); err != nil {
		return err
	}
	if err := p.AddDefSlice(ProviderDefRepositories); err != nil {
		return err
	}

	return nil
}
