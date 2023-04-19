// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package external

import "github.com/volvo-cars/lingon/pkg/terra"

func NewProvider(args ProviderArgs) *Provider {
	return &Provider{Args: args}
}

var _ terra.Provider = (*Provider)(nil)

type Provider struct {
	Args ProviderArgs
}

// LocalName returns the provider local name for [Provider].
func (p *Provider) LocalName() string {
	return "external"
}

// Source returns the provider source for [Provider].
func (p *Provider) Source() string {
	return "hashicorp/external"
}

// Version returns the provider version for [Provider].
func (p *Provider) Version() string {
	return "2.3.1"
}

// Configuration returns the configuration (args) for [Provider].
func (p *Provider) Configuration() interface{} {
	return p.Args
}

// ProviderArgs contains the configurations for provider.
type ProviderArgs struct{}