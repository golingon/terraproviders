// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package tls

import (
	provider "github.com/golingon/terraproviders/tls/4.0.0/provider"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewProvider(args ProviderArgs) *Provider {
	return &Provider{Args: args}
}

var _ terra.Provider = (*Provider)(nil)

type Provider struct {
	Args ProviderArgs
}

func (p *Provider) LocalName() string {
	return "tls"
}

func (p *Provider) Source() string {
	return "hashicorp/tls"
}

func (p *Provider) Version() string {
	return "4.0.0"
}

func (p *Provider) Configuration() interface{} {
	return p.Args
}

type ProviderArgs struct {
	// Proxy: optional
	Proxy *provider.Proxy `hcl:"proxy,block"`
}